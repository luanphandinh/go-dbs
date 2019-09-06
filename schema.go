package dbs

import (
	"database/sql"
	"fmt"
)

type Schema struct {
	Name     string  `json:"name"`
	Platform string  `json:"platform"`
	Tables   []Table `json:"tables"`
}

func (schema *Schema) Install(db *sql.DB) error {
	platform := GetPlatform(schema.Platform)
	if platform == nil {
		return fmt.Errorf("Invalid ")
	}
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	// create table
	for _, table := range schema.Tables {
		_, err := tx.Exec(platform.GetTableSQLCreate(&table))
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()
	return err
}
