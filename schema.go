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

	// create schema
	if schemaCreation := platform.GetSchemaCreateDeclarationSQL(schema); schemaCreation != "" {
		if _, err := tx.Exec(schemaCreation); err != nil {
			tx.Rollback()
			return err
		}
	}

	// create tables
	for _, table := range schema.Tables {
		if _, err := tx.Exec(platform.GetTableCreateSQL(schema.Name, &table)); err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()
	return err
}

func (schema *Schema) Drop(db *sql.DB) error {
	platform := GetPlatform(schema.Platform)
	if platform == nil {
		return fmt.Errorf("invalid platform")
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	// drop tables
	for _, table := range schema.Tables {
		if _, err := tx.Exec(platform.GetTableDropSQL(schema.Name, table.Name)); err != nil {
			tx.Rollback()
			return err
		}
	}

	// drop schema
	if schemaDrop := platform.GetSchemaDropDeclarationSQL(schema); schemaDrop != "" {
		if _, err := tx.Exec(schemaDrop); err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()
	return err
}
