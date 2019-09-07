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
		return fmt.Errorf("invalid platform")
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	// create tables
	for _, table := range schema.Tables {
		if _, err := tx.Exec(platform.GetTableCreateSQL(&table)); err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()
	return err
}
