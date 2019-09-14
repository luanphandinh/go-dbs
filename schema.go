package dbs

import (
	"database/sql"
	"fmt"
)

const SCHEMA = "SCHEMA"

type Schema struct {
	Name     string  `json:"name"`
	Platform string  `json:"platform"`
	Tables   []*Table `json:"tables"`
	Comment  string  `json:"comment"`
}

func (schema *Schema) WithName(name string) *Schema {
	schema.Name = name

	return schema
}

func (schema *Schema) OnPlatform(platform string) *Schema {
	schema.Platform = platform

	return schema
}

func (schema *Schema) WithComment(comment string) *Schema {
	schema.Comment = comment

	return schema
}

func (schema *Schema) AddTable(table *Table) *Schema {
	schema.Tables = append(schema.Tables, table)

	return schema
}

func (schema *Schema) AddTables(tables []*Table) *Schema {
	schema.Tables = append(schema.Tables, tables...)

	return schema
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
	if schemaCreation := platform.BuildSchemaCreateSQL(schema); schemaCreation != "" {
		if _, err := tx.Exec(schemaCreation); err != nil {
			tx.Rollback()
			return err
		}
	}

	// create tables
	for _, table := range schema.Tables {
		if _, err := tx.Exec(platform.BuildTableCreateSQL(schema.Name, table)); err != nil {
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

	for i := len(schema.Tables) - 1; i >= 0; i-- {
		if _, err := tx.Exec(platform.GetTableDropSQL(schema.Name, schema.Tables[i].Name)); err != nil {
			tx.Rollback()
			return err
		}
	}

	if schemaDrop := platform.GetSchemaDropDeclarationSQL(schema.Name); schemaDrop != "" {
		if _, err := tx.Exec(schemaDrop); err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()
	return err
}
