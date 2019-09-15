package dbs

import "strconv"

const SQLITE3 string = "sqlite3"

type SqlitePlatform struct{}

func (platform *SqlitePlatform) getDriverName() string {
	return SQLITE3
}

func (platform *SqlitePlatform) getDBConnectionString(server string, port int, user string, password string, dbName string) string {
	return dbName
}

func (platform *SqlitePlatform) chainCommands(commands ...string) string {
	return concatStrings(commands, ";\n")
}

func (platform *SqlitePlatform) getTypeDeclaration(col *Column) string {
	dbType := col.Type

	// @TODO: make some type reference that centralized all types together across platforms
	if inStringArray(col.Type, integerTypes) {
		dbType = "INTEGER"
	}

	if col.Length > 0 {
		return dbType + "(" + strconv.Itoa(col.Length) + ")"
	}

	return dbType
}

func (platform *SqlitePlatform) getUniqueDeclaration() string {
	return _getUniqueDeclaration()
}

func (platform *SqlitePlatform) getNotNullDeclaration() string {
	return _getNotNullDeclaration()
}

func (platform *SqlitePlatform) getPrimaryDeclaration(key []string) string {
	return _getPrimaryDeclaration(key)
}

func (platform *SqlitePlatform) getAutoIncrementDeclaration() string {
	return ""
}

func (platform *SqlitePlatform) getUnsignedDeclaration() string {
	return ""
}

func (platform *SqlitePlatform) getDefaultDeclaration(expression string) string {
	return _getDefaultDeclaration(expression)
}

func (platform *SqlitePlatform) getColumnCommentDeclaration(expression string) string {
	return ""
}

func (platform *SqlitePlatform) getColumnsCommentDeclaration(schema string, table *Table) []string {
	return make([]string, 0)
}

func (platform *SqlitePlatform) buildColumnsDeclarationSQL(cols []*Column) []string {
	return _buildColumnsDeclarationSQL(platform, cols)
}

func (platform *SqlitePlatform) getColumnCheckDeclaration(expression string) string {
	return _getColumnCheckDeclaration(expression)
}

func (platform *SqlitePlatform) buildColumnDeclarationSQL(col *Column) string {
	return _buildColumnDeclarationSQL(platform, col)
}

func (platform *SqlitePlatform) buildSchemaCreateSQL(schema *Schema) string {
	return ""
}

func (platform *SqlitePlatform) getSchemaCreateDeclarationSQL(schema string) string {
	return ""
}

func (platform *SqlitePlatform) getSchemaDropDeclarationSQL(schema string) string {
	return ""
}

func (platform *SqlitePlatform) getSchemaAccessName(schema string, name string) string {
	return name
}

func (platform *SqlitePlatform) getSchemaCommentDeclaration(schema string, expression string) string {
	return ""
}

func (platform *SqlitePlatform) getTableChecksDeclaration(expressions []string) []string {
	return _getTableChecksDeclaration(expressions)
}

func (platform *SqlitePlatform) getTableReferencesDeclarationSQL(schema string, foreignKeys []ForeignKey) []string {
	return _getTableReferencesDeclarationSQL(platform, schema, foreignKeys)
}

func (platform *SqlitePlatform) getTableCommentDeclarationSQL(name string, expression string) string {
	return ""
}

func (platform *SqlitePlatform) buildTableCreateSQL(schema string, table *Table) (tableString string) {
	return _buildTableCreateSQL(platform, schema, table)
}

func (platform *SqlitePlatform) getTableDropSQL(schema string, table string) (tableString string) {
	return _getTableDropSQL(platform, schema, table)
}

func (platform *SqlitePlatform) getSequenceCreateSQL(sequence string) string {
	return ""
}

func (platform *SqlitePlatform) getSequenceDropSQL(sequence string) string {
	return ""
}