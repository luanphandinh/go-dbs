package dbs

import (
	"database/sql"
	"log"
)

// Schema defined the db schema structure
type Schema struct {
	Name    string   `json:"name"`
	Tables  []*Table `json:"tables"`
	Comment string   `json:"comment"`

	db       *sql.DB
}

// WithName set the schema name
func (schema *Schema) WithName(name string) *Schema {
	schema.Name = name

	return schema
}

// SetDB set a db connection to schema
func (schema *Schema) SetDB(db *sql.DB) *Schema {
	schema.db = db

	return schema
}

// WithComment Set comment for schema
// This only works on postgresql
func (schema *Schema) WithComment(comment string) *Schema {
	schema.Comment = comment

	return schema
}

// AddTable add defined table to schema
func (schema *Schema) AddTable(table *Table) *Schema {
	schema.Tables = append(schema.Tables, table)

	return schema
}

// AddTables add a list of defined tables to schema
func (schema *Schema) AddTables(tables []*Table) *Schema {
	schema.Tables = append(schema.Tables, tables...)

	return schema
}

// HasTable return true if table exists
func (schema *Schema) HasTable(table string) bool {
	db := schema.db

	var name string
	if err := db.QueryRow(_platform().checkSchemaHasTableSQL(schema.Name, table)).Scan(&name); err != nil {
		return false
	} else {
		return name == table || name == _platform().getSchemaAccessName(schema.Name, table)
	}
}

// GetTables return all tables in schema
func (schema *Schema) GetTables() []string {
	db := schema.db

	tables := make([]string, 0)
	rows, err := db.Query(_platform().getSchemaTablesSQL(schema.Name))
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

// IsExists return true if schema exists
func (schema *Schema) IsExists() bool {
	db := schema.db

	command := _platform().checkSchemaExistSQL(schema.Name)
	if command == "" {
		return true
	}

	var name string
	if err := db.QueryRow(command).Scan(&name); err != nil {
		return false
	} else {
		return name == schema.Name
	}
}

// Install the schema
func (schema *Schema) Install() error {
	createSchemaSQL := _platform().buildSchemaCreateSQL(schema)
	if schema.IsExists() {
		createSchemaSQL = ""
	}

	createTableSQLs := make([]string, 0)
	tables := schema.GetTables()
	for _, table := range schema.Tables {
		if inStringArray(table.Name, tables) {
			continue
		}
		createTableSQLs = append(createTableSQLs, _platform().buildTableCreateSQL(schema.Name, table))
	}

	tx, err := schema.db.Begin()
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

	err = tx.Commit()
	return err
}

// Drop the schema
func (schema *Schema) Drop() error {
	tx, err := schema.db.Begin()
	if err != nil {
		return err
	}

	for i := len(schema.Tables) - 1; i >= 0; i-- {
		if _, err := tx.Exec(_platform().getTableDropSQL(schema.Name, schema.Tables[i].Name)); err != nil {
			tx.Rollback()
			return err
		}
	}

	if schemaDrop := _platform().getSchemaDropDeclarationSQL(schema.Name); schemaDrop != "" {
		if _, err := tx.Exec(schemaDrop); err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()
	return err
}
