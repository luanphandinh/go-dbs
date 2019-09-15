package dbs

import "strconv"

const SQLITE3 string = "sqlite3"

type dbSqlitePlatform struct{}

func (platform *dbSqlitePlatform) getDriverName() string {
	return SQLITE3
}

func (platform *dbSqlitePlatform) getDBConnectionString(server string, port int, user string, password string, dbName string) string {
	return dbName
}

func (platform *dbSqlitePlatform) chainCommands(commands ...string) string {
	return concatStrings(commands, ";\n")
}

func (platform *dbSqlitePlatform) getTypeDeclaration(col *Column) string {
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

func (platform *dbSqlitePlatform) getUniqueDeclaration() string {
	return _getUniqueDeclaration()
}

func (platform *dbSqlitePlatform) getNotNullDeclaration() string {
	return _getNotNullDeclaration()
}

func (platform *dbSqlitePlatform) getPrimaryDeclaration(key []string) string {
	return _getPrimaryDeclaration(key)
}

func (platform *dbSqlitePlatform) getAutoIncrementDeclaration() string {
	return ""
}

func (platform *dbSqlitePlatform) getUnsignedDeclaration() string {
	return ""
}

func (platform *dbSqlitePlatform) getDefaultDeclaration(expression string) string {
	return _getDefaultDeclaration(expression)
}

func (platform *dbSqlitePlatform) getColumnCommentDeclaration(expression string) string {
	return ""
}

func (platform *dbSqlitePlatform) getColumnsCommentDeclaration(schema string, table *Table) []string {
	return make([]string, 0)
}

func (platform *dbSqlitePlatform) buildColumnsDeclarationSQL(cols []*Column) []string {
	return _buildColumnsDeclarationSQL(platform, cols)
}

func (platform *dbSqlitePlatform) getColumnCheckDeclaration(expression string) string {
	return _getColumnCheckDeclaration(expression)
}

func (platform *dbSqlitePlatform) buildColumnDeclarationSQL(col *Column) string {
	return _buildColumnDeclarationSQL(platform, col)
}

func (platform *dbSqlitePlatform) buildSchemaCreateSQL(schema *Schema) string {
	return ""
}

func (platform *dbSqlitePlatform) getSchemaCreateDeclarationSQL(schema string) string {
	return ""
}

func (platform *dbSqlitePlatform) getSchemaDropDeclarationSQL(schema string) string {
	return ""
}

func (platform *dbSqlitePlatform) getSchemaAccessName(schema string, name string) string {
	return name
}

func (platform *dbSqlitePlatform) getSchemaCommentDeclaration(schema string, expression string) string {
	return ""
}

func (platform *dbSqlitePlatform) getTableChecksDeclaration(expressions []string) []string {
	return _getTableChecksDeclaration(expressions)
}

func (platform *dbSqlitePlatform) getTableReferencesDeclarationSQL(schema string, foreignKeys []ForeignKey) []string {
	return _getTableReferencesDeclarationSQL(platform, schema, foreignKeys)
}

func (platform *dbSqlitePlatform) getTableCommentDeclarationSQL(name string, expression string) string {
	return ""
}

func (platform *dbSqlitePlatform) buildTableCreateSQL(schema string, table *Table) (tableString string) {
	return _buildTableCreateSQL(platform, schema, table)
}

func (platform *dbSqlitePlatform) getTableDropSQL(schema string, table string) (tableString string) {
	return _getTableDropSQL(platform, schema, table)
}

func (platform *dbSqlitePlatform) getSequenceCreateSQL(sequence string) string {
	return ""
}

func (platform *dbSqlitePlatform) getSequenceDropSQL(sequence string) string {
	return ""
}