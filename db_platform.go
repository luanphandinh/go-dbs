package dbs

import "fmt"

type Platform interface {
	GetDBConnectionString(server string, port int, user string, password string, dbName string) string

	// Column attributes declarations
	GetTypeDeclaration(col *Column) string
	GetUniqueDeclaration() string
	GetNotNullDeclaration() string
	GetPrimaryDeclaration(key []string) string
	GetAutoIncrementDeclaration() string
	GetUnsignedDeclaration() string
	GetDefaultDeclaration(expression string) string
	// Check constraint is parsed but will be ignored in mysql5.7
	GetColumnCheckDeclaration(expression string) string

	GetColumnDeclarationSQL(col *Column) string
	GetColumnsDeclarationSQL(cols []Column) string

	// schema SQL declarations
	GetSchemaCreateDeclarationSQL(schema string) string
	GetSchemaDropDeclarationSQL(schema string) string

	// table SQL declarations
	GetTableName(schema string, table string) string
	// Check constraint is parsed but will be ignored in mysql5.7
	GetTableCheckDeclaration(expressions []string) string
	GetTableCreateSQL(schema string, table *Table) string
	GetTableDropSQL(schema string, table string) string
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

func _getPrimaryDeclaration(key []string) string {
	return fmt.Sprintf("PRIMARY KEY (%s)", concatString(key, ","))
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

func _getColumnCheckDeclaration(expression string) string {
	return fmt.Sprintf("CHECK (%s)", expression)
}

func _getTableCheckDeclaration(expressions []string) (evaluatedExpression string) {
	for index, expression := range expressions {
		if index == 0 {
			evaluatedExpression += fmt.Sprintf("CHECK (%s)", expression)
		} else {
			evaluatedExpression += fmt.Sprintf(", CHECK (%s)", expression)
		}
	}

	return evaluatedExpression
}

func _getSchemaCreateDeclarationSQL(schema string) string {
	return fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS %s", schema)
}

func _getSchemaDropDeclarationSQL(schema string) string {
	return fmt.Sprintf("DROP SCHEMA IF EXISTS %s", schema)
}

func _getColumnsDeclarationSQL(platform Platform, cols []Column) (colString string) {
	for index, col := range cols {
		if index == 0 {
			colString += fmt.Sprintf("%s", platform.GetColumnDeclarationSQL(&col))
		} else {
			colString += fmt.Sprintf(", %s", platform.GetColumnDeclarationSQL(&col))
		}
	}

	return colString
}

func _getTableCreateSQL(platform Platform, schema string, table *Table) string {
	check := ""
	if len(table.Check) > 0 {
		check = fmt.Sprintf(", %s", platform.GetTableCheckDeclaration(table.Check))
	}

	return fmt.Sprintf(
		"CREATE TABLE IF NOT EXISTS %s (%s, %s%s)",
		platform.GetTableName(schema, table.Name),
		platform.GetColumnsDeclarationSQL(table.Columns),
		platform.GetPrimaryDeclaration(table.PrimaryKey),
		check,
	)
}

func _getTableDropSQL(platform Platform, schema string, table string) string {
	return fmt.Sprintf("DROP TABLE IF EXISTS %s", platform.GetTableName(schema, table))
}
