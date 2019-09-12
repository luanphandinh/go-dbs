package dbs

import (
	"fmt"
)

type Platform interface {
	GetDriverName() string
	GetDBConnectionString(server string, port int, user string, password string, dbName string) string
	ChainCommands(commands ...string) string

	// Column attributes declarations
	GetTypeDeclaration(col *Column) string
	GetUniqueDeclaration() string
	GetNotNullDeclaration() string
	GetPrimaryDeclaration(key []string) string
	GetAutoIncrementDeclaration() string
	GetUnsignedDeclaration() string
	GetDefaultDeclaration(expression string) string
	GetColumnCommentDeclaration(expression string) string // For inline comment
	GetColumnsCommentDeclaration(schema string, table *Table) []string // For external SQL COMMENT on postgresql
	// Check constraint is parsed but will be ignored in mysql5.7
	GetColumnCheckDeclaration(expression string) string

	GetColumnDeclarationSQL(col *Column) string
	GetColumnsDeclarationSQL(cols []Column) []string

	// schema SQL declarations
	GetSchemaCreateDeclarationSQL(schema string) string
	GetSchemaDropDeclarationSQL(schema string) string

	// table SQL declarations
	GetSchemaAccessName(schema string, name string) string
	// Check constraint is parsed but will be ignored in mysql5.7
	GetTableChecksDeclaration(expressions []string) []string
	GetTableCreateSQL(schema string, table *Table) string
	GetTableDropSQL(schema string, table string) string
	GetTableCommentDeclarationSQL(name string, expression string) string

	GetSequenceCreateSQL(sequence string) string
	GetSequenceDropSQL(sequence string) string
}

func GetPlatform(platform string) Platform {
	if platform == MYSQL57 {
		return &MySql57Platform{}
	}

	if platform == MYSQL80 {
		return &MySql80Platform{}
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
	return fmt.Sprintf("PRIMARY KEY (%s)", concatString(key, ", "))
}

func _getUnsignedDeclaration() string {
	return "UNSIGNED"
}

func _getDefaultDeclaration(expression string) string {
	return fmt.Sprintf("DEFAULT %s", expression)
}

func _getColumnCommentDeclaration(expression string) string {
	return fmt.Sprintf("COMMENT '%s'", expression)
}

func _getColumnCheckDeclaration(expression string) string {
	return fmt.Sprintf("CHECK (%s)", expression)
}

func _getTableChecksDeclaration(expressions []string) []string {
	evaluated := make([]string, 0)

	for _, expression := range expressions {
		evaluated = append(evaluated, fmt.Sprintf("CHECK (%s)", expression))
	}

	return evaluated
}

func _getSchemaCreateDeclarationSQL(schema string) string {
	return fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS %s", schema)
}

func _getSchemaDropDeclarationSQL(schema string) string {
	return fmt.Sprintf("DROP SCHEMA IF EXISTS %s CASCADE", schema)
}

func _getColumnDeclarationSQL(platform Platform, col *Column) (colString string) {
	columnString := fmt.Sprintf("%s %s", col.Name, platform.GetTypeDeclaration(col))

	if col.Unsigned {
		columnString += " " + platform.GetUnsignedDeclaration()
	}

	if col.NotNull {
		columnString += " " + platform.GetNotNullDeclaration()
	}

	if col.Default != "" {
		columnString += " " + platform.GetDefaultDeclaration(col.Default)
	}

	if col.AutoIncrement {
		columnString += " " + platform.GetAutoIncrementDeclaration()
	}

	if col.Unique {
		columnString += " " + platform.GetUniqueDeclaration()
	}

	if col.Check != "" {
		columnString += " " + platform.GetColumnCheckDeclaration(col.Check)
	}

	if col.Comment != "" {
		columnString += " " + platform.GetColumnCommentDeclaration(col.Comment)
	}

	return columnString
}

func _getColumnsDeclarationSQL(platform Platform, cols []Column) []string {
	declarations := make([]string, len(cols))
	for index, col := range cols {
		declarations[index] = platform.GetColumnDeclarationSQL(&col)
	}

	return declarations
}

func _getTableCreateSQL(platform Platform, schema string, table *Table) string {
	tableName := platform.GetSchemaAccessName(schema, table.Name)
	tableCreation := make([]string, 0)
	tableCreation = append(tableCreation, platform.GetColumnsDeclarationSQL(table.Columns)...)
	tableCreation = append(tableCreation, platform.GetPrimaryDeclaration(table.PrimaryKey))
	tableCreation = append(tableCreation, platform.GetTableChecksDeclaration(table.Check)...)

	tableDeclaration :=  fmt.Sprintf(
		"CREATE TABLE IF NOT EXISTS %s (\n\t%s\n)",
		tableName,
		concatString(tableCreation, ",\n\t"),
	)

	commands := make([]string, 0)
	commands = append(commands, tableDeclaration)
	if table.Comment != "" {
		commands = append(commands, platform.GetTableCommentDeclarationSQL(tableName, table.Comment))
	}
	commands = append(commands, platform.GetColumnsCommentDeclaration(schema, table)...)

	return platform.ChainCommands(commands...)
}

func _getTableDropSQL(platform Platform, schema string, table string) string {
	return fmt.Sprintf("DROP TABLE IF EXISTS %s", platform.GetSchemaAccessName(schema, table))
}

func _getSequenceCreateSQL(sequence string) string {
	return fmt.Sprintf("CREATE SEQUENCE IF NOT EXISTS %s", sequence)
}

func _getSequenceDropSQL(sequence string) string {
	return fmt.Sprintf("DROP SEQUENCE IF EXISTS %s", sequence)
}

func _chainCommands(commands ...string) string {
	return concatString(commands, ";\n")
}
