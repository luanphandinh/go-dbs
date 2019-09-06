package dbs

import "database/sql"

type Schema struct {
	Name     string  `json:"name"`
	Platform string  `json:"platform"`
	Tables   []Table `json:"tables"`
}

func (schema *Schema) Install(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	// create table
	for _, table := range schema.Tables {
		_, err := tx.Exec(table.GetSQLCreateTable(GetPlatform(schema.Platform)))
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()
	return err
}
