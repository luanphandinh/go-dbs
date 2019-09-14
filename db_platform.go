package dbs

import "fmt"

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
	// Checks constraint is parsed but will be ignored in mysql5.7
	GetColumnCheckDeclaration(expression string) string

	BuildColumnDeclarationSQL(col *Column) string
	BuildColumnsDeclarationSQL(cols []*Column) []string

	// schema SQL declarations
	BuildSchemaCreateSQL(schema *Schema) string
	GetSchemaCreateDeclarationSQL(schema string) string
	GetSchemaDropDeclarationSQL(schema string) string

	// table SQL declarations
	GetSchemaAccessName(schema string, name string) string
	GetSchemaCommentDeclaration(schema string, expression string) string
	// Checks constraint is parsed but will be ignored in mysql5.7
	GetTableChecksDeclaration(expressions []string) []string
	BuildTableCreateSQL(schema string, table *Table) string
	GetTableDropSQL(schema string, table string) string
	GetTableCommentDeclarationSQL(name string, expression string) string
	GetTableReferencesDeclarationSQL(schema string, foreignKeys []ForeignKey) []string

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

	if platform == MSSQL {
		return &MsSqlPlatform{}
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
	return fmt.Sprintf("PRIMARY KEY (%s)", concatStrings(key, ", "))
}

func _getDefaultDeclaration(expression string) string {
	return fmt.Sprintf("DEFAULT %s", expression)
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

func _getSchemaDropDeclarationSQL(schema string) string {
	return fmt.Sprintf("DROP SCHEMA IF EXISTS %s CASCADE", schema)
}

func _buildColumnDeclarationSQL(platform Platform, col *Column) (colString string) {
	declaration := make([]string, 0)
	declaration = append(declaration, fmt.Sprintf("%s %s", col.Name, platform.GetTypeDeclaration(col)))

	if col.Unsigned {
		declaration = append(declaration, platform.GetUnsignedDeclaration())
	}

	if col.NotNull {
		declaration = append(declaration, platform.GetNotNullDeclaration())
	}

	if col.Default != "" {
		declaration = append(declaration, platform.GetDefaultDeclaration(col.Default))
	}

	if col.AutoIncrement {
		declaration = append(declaration, platform.GetAutoIncrementDeclaration())
	}

	if col.Unique {
		declaration = append(declaration, platform.GetUniqueDeclaration())
	}

	if col.Check != "" {
		declaration = append(declaration, platform.GetColumnCheckDeclaration(col.Check))
	}

	if col.Comment != "" {
		declaration = append(declaration, platform.GetColumnCommentDeclaration(col.Comment))
	}

	return concatStrings(declaration, " ")
}

func _buildColumnsDeclarationSQL(platform Platform, cols []*Column) []string {
	declarations := make([]string, len(cols))
	for index, col := range cols {
		declarations[index] = platform.BuildColumnDeclarationSQL(col)
	}

	return declarations
}

func _getTableReferencesDeclarationSQL(platform Platform, schema string, foreignKeys []ForeignKey) []string {
	keys := make([]string, 0)
	for _, key := range foreignKeys {
		keys = append(
			keys,
			fmt.Sprintf(
				"FOREIGN KEY (%s) REFERENCES %s",
				key.Referer,
				platform.GetSchemaAccessName(schema, key.Reference),
			),
		)
	}

	return keys
}

func _buildTableCreateSQL(platform Platform, schema string, table *Table) string {
	tableName := platform.GetSchemaAccessName(schema, table.Name)
	tableCreation := make([]string, 0)
	tableCreation = append(tableCreation, platform.BuildColumnsDeclarationSQL(table.Columns)...)
	tableCreation = append(tableCreation, platform.GetPrimaryDeclaration(table.PrimaryKey))
	tableCreation = append(tableCreation, platform.GetTableReferencesDeclarationSQL(schema, table.ForeignKeys)...)
	tableCreation = append(tableCreation, platform.GetTableChecksDeclaration(table.Checks)...)

	tableDeclaration :=  fmt.Sprintf(
		"CREATE TABLE %s (\n\t%s\n)",
		tableName,
		concatStrings(tableCreation, ",\n\t"),
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
