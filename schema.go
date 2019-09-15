package dbs

import (
	"database/sql"
	"fmt"
)

type Schema struct {
	Name     string   `json:"name"`
	Platform string   `json:"platform"`
	Tables   []*Table `json:"tables"`
	Comment  string   `json:"comment"`
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

	if schemaCreation := platform.buildSchemaCreateSQL(schema); schemaCreation != "" {
		if _, err := tx.Exec(schemaCreation); err != nil {
			tx.Rollback()
			return err
		}
	}

	for _, table := range schema.Tables {
		if _, err := tx.Exec(platform.buildTableCreateSQL(schema.Name, table)); err != nil {
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
		if _, err := tx.Exec(platform.getTableDropSQL(schema.Name, schema.Tables[i].Name)); err != nil {
			tx.Rollback()
			return err
		}
	}

	if schemaDrop := platform.getSchemaDropDeclarationSQL(schema.Name); schemaDrop != "" {
		if _, err := tx.Exec(schemaDrop); err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()
	return err
}
