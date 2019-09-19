package dbs

import "strconv"

const mssql string = "sqlserver"

type dbMsSQLPlatform struct{}

func (platform *dbMsSQLPlatform) getDriverName() string {
	return mssql
}

func (platform *dbMsSQLPlatform) getDBConnectionString(server string, port int, user string, password string, dbName string) string {
	info := make([]string, 0)
	info = append(info, "server=" + server)
	info = append(info, "user id=" + user)
	info = append(info, "password=" + password)
	info = append(info, "database=" + dbName)

	return concatStrings(info, ";")
}

func (platform *dbMsSQLPlatform) chainCommands(commands ...string) string {
	return concatStrings(commands, ";\nGO\n")
}

func (platform *dbMsSQLPlatform) getTypeDeclaration(col *Column) string {
	if col.Length > 0 {
		return col.Type + "(" + strconv.Itoa(col.Length) + ")"
	}

	return col.Type
}

func (platform *dbMsSQLPlatform) getUniqueDeclaration() string {
	return _getUniqueDeclaration()
}

func (platform *dbMsSQLPlatform) getNotNullDeclaration() string {
	return _getNotNullDeclaration()
}

func (platform *dbMsSQLPlatform) getPrimaryDeclaration(key []string) string {
	return _getPrimaryDeclaration(key)
}

func (platform *dbMsSQLPlatform) getAutoIncrementDeclaration() string {
	return "IDENTITY(1,1)"
}

func (platform *dbMsSQLPlatform) getUnsignedDeclaration() string {
	return ""
}

func (platform *dbMsSQLPlatform) buildColumnDeclarationSQL(col *Column) string {
	return _buildColumnDeclarationSQL(platform, col)
}

func (platform *dbMsSQLPlatform) buildColumnsDeclarationSQL(cols []*Column) []string {
	return _buildColumnsDeclarationSQL(platform, cols)
}

func (platform *dbMsSQLPlatform) getColumnCommentDeclaration(expression string) string {
	return ""
}

func (platform *dbMsSQLPlatform) getColumnsCommentDeclaration(schema string, table *Table) []string {
	return make([]string, 0)
}

func (platform *dbMsSQLPlatform) getColumnCheckDeclaration(expression string) string {
	return _getColumnCheckDeclaration(expression)
}

func (platform *dbMsSQLPlatform) buildSchemaCreateSQL(schema *Schema) string {
	return platform.getSchemaCreateDeclarationSQL(schema.Name)
}

func (platform *dbMsSQLPlatform) getSchemaCreateDeclarationSQL(schema string) string {
	return "CREATE SCHEMA " + schema
}

func (platform *dbMsSQLPlatform) getSchemaDropDeclarationSQL(schema string) string {
	return "DROP SCHEMA IF EXISTS " + schema
}

func (platform *dbMsSQLPlatform) getDefaultDeclaration(expression string) string {
	return _getDefaultDeclaration(expression)
}

func (platform *dbMsSQLPlatform) getSchemaAccessName(schema string, name string) string {
	return schema + "." + name
}

func (platform *dbMsSQLPlatform) getSchemaCommentDeclaration(schema string, expression string) string {
	return ""
}

func (platform *dbMsSQLPlatform) getTableChecksDeclaration(expressions []string) []string {
	return _getTableChecksDeclaration(expressions)
}

func (platform *dbMsSQLPlatform) getTableReferencesDeclarationSQL(schema string, foreignKeys []ForeignKey) []string {
	return _getTableReferencesDeclarationSQL(platform, schema, foreignKeys)
}

func (platform *dbMsSQLPlatform) getTableCommentDeclarationSQL(name string, expression string) string {
	return ""
}

func (platform *dbMsSQLPlatform) buildTableCreateSQL(schema string, table *Table) (tableString string) {
	return _buildTableCreateSQL(platform, schema, table)
}

func (platform *dbMsSQLPlatform) getTableDropSQL(schema string, table string) (tableString string) {
	return _getTableDropSQL(platform, schema, table)
}

func (platform *dbMsSQLPlatform) getSequenceCreateSQL(sequence string) string {
	return "CREATE SEQUENCE " + sequence
}

func (platform *dbMsSQLPlatform) getSequenceDropSQL(sequence string) string {
	return "DROP SEQUENCE " + sequence
}

func (platform *dbMsSQLPlatform) checkSchemaExistSQL(schema string) string {
	return "SELECT name FROM sys.schemas WHERE name = '" + schema + "'"
}

func (platform *dbMsSQLPlatform) checkSchemaHasTableSQL(schema string, table string) string {
	return "SELECT TABLE_NAME FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_NAME = '" + table + "' AND TABLE_SCHEMA = '" + schema + "'"
}

func (platform *dbMsSQLPlatform) getSchemaTablesSQL(schema string) string {
	return "SELECT TABLE_NAME FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_SCHEMA = '" + schema + "'"
}

func (platform *dbMsSQLPlatform) getTableColumnsSQL(schema string , table string) string {
	return ""
}
