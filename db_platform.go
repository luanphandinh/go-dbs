package dbs

import "fmt"

type Platform interface {
	// Column attributes declarations
	GetTypeDeclaration(col *Column) string
	GetUniqueDeclaration() string
	GetNotNullDeclaration() string
	GetPrimaryDeclaration(table *Table) string
	GetAutoIncrementDeclaration() string
	GetUnsignedDeclaration() string
	GetDefaultDeclaration(expression string) string

	GetColumnDeclarationSQL(col *Column) string

	// schema SQL declarations
	GetSchemaCreateDeclarationSQL(schema *Schema) string
	GetSchemaDropDeclarationSQL(schema *Schema) string

	// table SQL declarations
	GetTableName(schema string, table *Table) string
	GetTableCreateSQL(schema string, table *Table) string
	GetTableDropSQL(schema string, table *Table) string
}

func GetPlatform(platform string) Platform {
	if platform == MYSQL {
		return &MySqlPlatform{}
	}

	if platform == SQLITE3 {
		return &SqlitePlatform{}
	}

	if platform == POSTGRES {
		return &PostgresPlatform{}
	}

	return nil
}

func _getUniqueDeclaration() string {
	return "UNIQUE"
}

func _getNotNullDeclaration() string {
	return "NOT NULL"
}

func _getPrimaryDeclaration(table *Table) string {
	return fmt.Sprintf("PRIMARY KEY (%s)", concatString(table.PrimaryKey, ","))
}

func _getAutoIncrementDeclaration() string {
	return "AUTO_INCREMENT"
}

func _getUnsignedDeclaration() string {
	return "UNSIGNED"
}

func _getDefaultDeclaration(expression string) string {
	return fmt.Sprintf("DEFAULT %s", expression)
}

func _getSchemaCreateDeclarationSQL(schema *Schema) string {
	return fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS %s", schema.Name)
}

func _getSchemaDropDeclarationSQL(schema *Schema) string {
	return fmt.Sprintf("DROP SCHEMA IF EXISTS %s", schema.Name)
}

func _getTableCreateSQL(platform Platform, schema string, table *Table) string {
	cols := ""
	for index, col := range table.Columns {
		if index == 0 {
			cols += fmt.Sprintf("%s", platform.GetColumnDeclarationSQL(&col))
		} else {
			cols += fmt.Sprintf(", %s", platform.GetColumnDeclarationSQL(&col))
		}
	}

	return fmt.Sprintf(
		"CREATE TABLE IF NOT EXISTS %s (%s, %s)",
		platform.GetTableName(schema, table),
		cols,
		platform.GetPrimaryDeclaration(table),
	)
}

func _getTableDropSQL(platform Platform, schema string, table *Table) string {
	return fmt.Sprintf("DROP TABLE IF EXISTS %s", platform.GetTableName(schema, table))
}
