package dbs

import (
	"log"
	"sync"
)

func checkSchemaExists(schema string) bool {
	command := _platform().checkSchemaExistSQL(schema)
	if command == "" {
		return true
	}

	var name string
	if err := _db().QueryRow(command).Scan(&name); err != nil {
		return false
	}

	return name == schema
}

func checkSchemaHasTableSQL(schema string, table string) bool {
	var name string
	if err := _db().QueryRow(_platform().checkSchemaHasTableSQL(schema, table)).Scan(&name); err != nil {
		return false
	}

	return name == table || name == _platform().getSchemaAccessName(schema, table)
}

func fetchTables(schema string) []string {
	tables := make([]string, 0)
	rows, err := _db().Query(_platform().getSchemaTablesSQL(schema))
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var name string
	for rows.Next() {
		err := rows.Scan(&name)
		if err != nil {
			log.Fatal(err)
		}

		tables = append(tables, name)
	}

	return tables
}

func fetchTableColumns(schema string, table string) []*Column {
	rows, err := _db().Query(_platform().getTableColumnsSQL(schema, table))
	if err != nil {
		log.Fatal(err)
	}

	return _platform().parseTableColumns(rows)
}

func fetchTableColumnNames(schema string, table string) []string {
	rows, err := _db().Query(_platform().getTableColumnNamesSQL(schema, table))
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	columns := make([]string, 0)
	var name string
	for rows.Next() {
		err := rows.Scan(&name)
		if err != nil {
			log.Fatal(err)
		}

		columns = append(columns, name)
	}

	return columns
}

// @TODO: This func will install schema just like install
// But run concurrently
// Should be benchmark alongside with install to see which one is better
func installConcurrent(schema *Schema) error {
	var wg sync.WaitGroup
	createSchemaSQL := _platform().buildSchemaCreateSQL(schema)

	wg.Add(2)
	go func(wg *sync.WaitGroup) {
		if checkSchemaExists(schema.name) {
			createSchemaSQL = ""
		}
		wg.Done()
	}(&wg)

	var existedTables []string
	go func(wg *sync.WaitGroup) {
		existedTables = fetchTables(schema.name)
		wg.Done()
	}(&wg)

	wg.Wait()

	alterTableSQLs := make([]string, 0)
	createTableSQLs := make([]string, 0)
	createIndexSQLs := make([]string, 0)

	var mutex = &sync.Mutex{}
	for _, table := range schema.tables {
		if inStringArray(table.name, existedTables) {
			wg.Add(1)
			go func(table *Table, wg *sync.WaitGroup) {
				cols := fetchTableColumnNames(schema.name, table.name)
				mutex.Lock()
				for _, col := range table.columns {
					if inStringArray(col.name, cols) {
						continue
					}
					alterTableSQLs = append(alterTableSQLs, _platform().buildTableAddColumnSQL(schema.name, table.name, col))
				}
				mutex.Unlock()

				wg.Done()
			}(table, &wg)
		} else {
			createTableSQLs = append(createTableSQLs, _platform().buildTableCreateSQL(schema.name, table))
			createIndexSQLs = append(createIndexSQLs, _platform().getTableIndexesDeclarationSQL(schema.name, table.name, table.indexes)...)
		}
	}

	wg.Wait()
	tx, err := _db().Begin()
	if err != nil {
		return err
	}

	if createSchemaSQL != "" {
		if _, err := tx.Exec(createSchemaSQL); err != nil {
			tx.Rollback()
			return err
		}
	}

	for _, createTableSQL := range createTableSQLs {
		wg.Add(1)
		go func(createTableSQL string, wg *sync.WaitGroup) {
			if _, err := tx.Exec(createTableSQL); err != nil {
				tx.Rollback()
				log.Fatal(err.Error())
			}
			wg.Done()
		}(createTableSQL, &wg)
	}

	wg.Wait()

	for _, alterTableSQL := range alterTableSQLs {
		wg.Add(1)
		go func(alterTableSQL string, wg *sync.WaitGroup) {
			if _, err := tx.Exec(alterTableSQL); err != nil {
				tx.Rollback()
				log.Fatal(err.Error())
			}
			wg.Done()
		}(alterTableSQL, &wg)
	}

	for _, createIndexSQL := range createIndexSQLs {
		wg.Add(1)
		go func(createIndexSQL string, wg *sync.WaitGroup) {
			if _, err := tx.Exec(createIndexSQL); err != nil {
				tx.Rollback()
				log.Fatal(err.Error())
			}
			wg.Done()
		}(createIndexSQL, &wg)
	}

	wg.Wait()

	err = tx.Commit()

	return err
}

// @TODO: This func is a real mess, need to clean up later.
func install(schema *Schema) error {
	createTableSQLs := make([]string, 0)
	alterTableSQLs := make([]string, 0)
	createIndexSQLs := make([]string, 0)

	createSchemaSQL := _platform().buildSchemaCreateSQL(schema)
	if checkSchemaExists(schema.name) {
		createSchemaSQL = ""
	}

	existedTables := fetchTables(schema.name)
	for _, table := range schema.tables {
		if inStringArray(table.name, existedTables) {
			cols := fetchTableColumnNames(schema.name, table.name)
			for _, col := range table.columns {
				if inStringArray(col.name, cols) {
					continue
				}
				alterTableSQLs = append(alterTableSQLs, _platform().buildTableAddColumnSQL(schema.name, table.name, col))
			}
			continue
		}

		createTableSQLs = append(createTableSQLs, _platform().buildTableCreateSQL(schema.name, table))
		createIndexSQLs = append(createIndexSQLs, _platform().getTableIndexesDeclarationSQL(schema.name, table.name, table.indexes)...)
	}

	tx, err := _db().Begin()
	if err != nil {
		return err
	}

	if createSchemaSQL != "" {
		if _, err := tx.Exec(createSchemaSQL); err != nil {
			tx.Rollback()
			return err
		}
	}

	for _, createTableSQL := range createTableSQLs {
		if _, err := tx.Exec(createTableSQL); err != nil {
			tx.Rollback()
			return err
		}
	}

	for _, alterTableSQL := range alterTableSQLs {
		if _, err := tx.Exec(alterTableSQL); err != nil {
			tx.Rollback()
			return err
		}
	}

	for _, createIndexSQL := range createIndexSQLs {
		if _, err := tx.Exec(createIndexSQL); err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()
	return err
}

func drop(schema *Schema) error {
	tx, err := _db().Begin()
	if err != nil {
		return err
	}

	for i := len(schema.tables) - 1; i >= 0; i-- {
		if _, err := tx.Exec(_platform().getTableDropSQL(schema.name, schema.tables[i].name)); err != nil {
			tx.Rollback()
			return err
		}
	}

	if schemaDrop := _platform().getSchemaDropDeclarationSQL(schema.name); schemaDrop != "" {
		if _, err := tx.Exec(schemaDrop); err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()
	return err
}

func dropConcurrent(schema *Schema) error {
	tx, err := _db().Begin()
	if err != nil {
		return err
	}

	for i := len(schema.tables) - 1; i >= 0; i-- {
		if _, err := tx.Exec(_platform().getTableDropSQL(schema.name, schema.tables[i].name)); err != nil {
			tx.Rollback()
			return err
		}
	}

	if schemaDrop := _platform().getSchemaDropDeclarationSQL(schema.name); schemaDrop != "" {
		if _, err := tx.Exec(schemaDrop); err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()
	return err
}
