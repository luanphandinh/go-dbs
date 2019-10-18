package dbs

import (
	"database/sql"
)

type dbPlatform interface {
	getDriverName() string
	getDBConnectionString(server string, port int, user string, password string, dbName string) string
	chainCommands(commands ...string) string

	// Column attributes declarations
	getTypeDeclaration(col *Column) string
	getUniqueDeclaration() string
	getNotNullDeclaration() string
	getPrimaryDeclaration(key []string) string
	getAutoIncrementDeclaration() string
	getUnsignedDeclaration() string
	getDefaultDeclaration(expression string) string
	getColumnCommentDeclaration(expression string) string // For inline comment
	getColumnsCommentDeclaration(schema string, table *Table) []string // For external SQL COMMENT on postgresql
	getColumnCheckDeclaration(expression string) string // checks constraint is parsed but will be ignored in mysql5.7
	buildColumnDefinitionSQL(col *Column) string
	buildColumnsDefinitionSQL(cols []*Column) []string

	// schema SQL declarations
	buildSchemaCreateSQL(schema *Schema) string
	getSchemaCreateDeclarationSQL(schema string) string
	getSchemaDropDeclarationSQL(schema string) string
	getSchemaCommentDeclaration(schema string, expression string) string

	// table SQL declarations
	getSchemaAccessName(schema string, name string) string
	getTableChecksDeclaration(expressions []string) []string // checks constraint is parsed but will be ignored in mysql5.7
	buildTableCreateSQL(schema string, table *Table) string
	buildTableAddColumnSQL(schema string, table string, col *Column) string
	getTableDropSQL(schema string, table string) string
	getTableCommentDeclarationSQL(name string, expression string) string
	getTableReferencesDeclarationSQL(schema string, foreignKeys []*ForeignKey) []string
	getTableIndexesDeclarationSQL(schema string, table string, indexes []*TableIndex) []string

	getSequenceCreateSQL(sequence string) string
	getSequenceDropSQL(sequence string) string

	// Actions get, set, check
	checkSchemaExistSQL(schema string) string
	checkSchemaHasTableSQL(schema string, table string) string
	getSchemaTablesSQL(schema string) string

	getTableColumnNamesSQL(schema string, table string) string

	// Query
	getQueryLimitDeclaration(limit string) string
	getQueryOffsetDeclaration(offset string) string
	// Combine of getQueryLimitDeclaration() & getQueryOffsetDeclaration()
	getPagingDeclaration(limit string, offset string) string

	// @TODO: these are experiment methods and have no actual value for now.
	getTableColumnsSQL(schema string , table string) string
	parseTableColumns(rows *sql.Rows) []*Column // parse rows returned from getTableColumnsSQL()
	columnDiff(col1 *Column, col2 *Column) bool
}

func _getUniqueDeclaration() string {
	return "UNIQUE"
}

func _getNotNullDeclaration() string {
	return "NOT NULL"
}

func _getPrimaryDeclaration(key []string) string {
	return "PRIMARY KEY (" + concatStrings(key, ", ") + ")"
}

func _getDefaultDeclaration(expression string) string {
	return "DEFAULT " + expression
}

func _getColumnCheckDeclaration(expression string) string {
	return "CHECK (" + expression + ")"
}

func _getTableChecksDeclaration(expressions []string) []string {
	evaluated := make([]string, 0)

	for _, expression := range expressions {
		evaluated = append(evaluated, "CHECK (" + expression + ")")
	}

	return evaluated
}

func _getSchemaDropDeclarationSQL(schema string) string {
	return "DROP SCHEMA IF EXISTS " + schema + " CASCADE"
}

func _buildColumnDefinitionSQL(platform dbPlatform, col *Column) (colString string) {
	declaration := make([]string, 0)
	declaration = append(declaration, col.name)
	declaration = append(declaration, platform.getTypeDeclaration(col))

	if col.unsigned {
		declaration = append(declaration, platform.getUnsignedDeclaration())
	}

	if col.notNull {
		declaration = append(declaration, platform.getNotNullDeclaration())
	}

	if col.defaultValue != "" {
		declaration = append(declaration, platform.getDefaultDeclaration(col.defaultValue))
	}

	if col.autoIncrement {
		declaration = append(declaration, platform.getAutoIncrementDeclaration())
	}

	if col.unique {
		declaration = append(declaration, platform.getUniqueDeclaration())
	}

	if col.check != "" {
		declaration = append(declaration, platform.getColumnCheckDeclaration(col.check))
	}

	if col.comment != "" {
		declaration = append(declaration, platform.getColumnCommentDeclaration(col.comment))
	}

	return concatStrings(declaration, " ")
}

func _buildColumnsDefinitionSQL(platform dbPlatform, cols []*Column) []string {
	declarations := make([]string, len(cols))
	for index, col := range cols {
		declarations[index] = platform.buildColumnDefinitionSQL(col)
	}

	return declarations
}

func _getTableReferencesDeclarationSQL(platform dbPlatform, schema string, foreignKeys []*ForeignKey) []string {
	keys := make([]string, 0)
	for _, key := range foreignKeys {
		keys = append(
			keys,
			"FOREIGN KEY (" + key.referer + ") REFERENCES " + platform.getSchemaAccessName(schema, key.reference),
		)
	}

	return keys
}

func _getTableIndexesDeclarationSQL(platform dbPlatform, schema string, table string, indexes []*TableIndex) []string {
	statements := make([]string, len(indexes))
	for i, index := range indexes {
		statements[i] = "CREATE INDEX " + index.name +
			" ON " + platform.getSchemaAccessName(schema, table) +
			" (" + concatStrings(index.cols, ", ") + ")"
	}

	return statements
}

func _buildTableCreateSQL(platform dbPlatform, schema string, table *Table) string {
	tableName := platform.getSchemaAccessName(schema, table.name)
	tableCreation := make([]string, 0)
	tableCreation = append(tableCreation, platform.buildColumnsDefinitionSQL(table.columns)...)
	if len(table.primaryKey) > 0 {
		tableCreation = append(tableCreation, platform.getPrimaryDeclaration(table.primaryKey))
	}
	tableCreation = append(tableCreation, platform.getTableReferencesDeclarationSQL(schema, table.foreignKeys)...)
	tableCreation = append(tableCreation, platform.getTableChecksDeclaration(table.checks)...)

	tableDeclaration := "CREATE TABLE " + tableName + " (\n\t" + concatStrings(tableCreation, ",\n\t") + "\n)"

	commands := make([]string, 0)
	commands = append(commands, tableDeclaration)
	if table.comment != "" {
		commands = append(commands, platform.getTableCommentDeclarationSQL(tableName, table.comment))
	}
	commands = append(commands, platform.getColumnsCommentDeclaration(schema, table)...)

	return platform.chainCommands(commands...)
}

func _getTableDropSQL(platform dbPlatform, schema string, table string) string {
	return "DROP TABLE IF EXISTS " + platform.getSchemaAccessName(schema, table)
}

func _getPagingDeclaration(platform dbPlatform, limit string, offset string) string {
	query := make([]string, 0)
	if limit != "" {
		query = append(query, platform.getQueryLimitDeclaration(limit))
	}

	if offset != "" {
		query = append(query, platform.getQueryOffsetDeclaration(offset))
	}

	return concatStrings(query, " ")
}
