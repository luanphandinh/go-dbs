package dbs

import "strconv"

const MSSQL string = "sqlserver"

type MsSQLPlatform struct{}

func (platform *MsSQLPlatform) getDriverName() string {
	return MSSQL
}

func (platform *MsSQLPlatform) getDBConnectionString(server string, port int, user string, password string, dbName string) string {
	info := make([]string, 0)
	info = append(info, "server=" + server)
	info = append(info, "user id=" + user)
	info = append(info, "password=" + password)
	info = append(info, "database=" + dbName)

	return concatStrings(info, ";")
}

func (platform *MsSQLPlatform) chainCommands(commands ...string) string {
	return concatStrings(commands, ";\nGO\n")
}

func (platform *MsSQLPlatform) getTypeDeclaration(col *Column) string {
	if col.Length > 0 {
		return col.Type + "(" + strconv.Itoa(col.Length) + ")"
	}

	return col.Type
}

func (platform *MsSQLPlatform) getUniqueDeclaration() string {
	return _getUniqueDeclaration()
}

func (platform *MsSQLPlatform) getNotNullDeclaration() string {
	return _getNotNullDeclaration()
}

func (platform *MsSQLPlatform) getPrimaryDeclaration(key []string) string {
	return _getPrimaryDeclaration(key)
}

func (platform *MsSQLPlatform) getAutoIncrementDeclaration() string {
	return "IDENTITY(1,1)"
}

func (platform *MsSQLPlatform) getUnsignedDeclaration() string {
	return ""
}

func (platform *MsSQLPlatform) buildColumnDeclarationSQL(col *Column) string {
	return _buildColumnDeclarationSQL(platform, col)
}

func (platform *MsSQLPlatform) buildColumnsDeclarationSQL(cols []*Column) []string {
	return _buildColumnsDeclarationSQL(platform, cols)
}

func (platform *MsSQLPlatform) getColumnCommentDeclaration(expression string) string {
	return ""
}

func (platform *MsSQLPlatform) getColumnsCommentDeclaration(schema string, table *Table) []string {
	return make([]string, 0)
}

func (platform *MsSQLPlatform) getColumnCheckDeclaration(expression string) string {
	return _getColumnCheckDeclaration(expression)
}

func (platform *MsSQLPlatform) buildSchemaCreateSQL(schema *Schema) string {
	return platform.getSchemaCreateDeclarationSQL(schema.Name)
}

func (platform *MsSQLPlatform) getSchemaCreateDeclarationSQL(schema string) string {
	return "CREATE SCHEMA " + schema
}

func (platform *MsSQLPlatform) getSchemaDropDeclarationSQL(schema string) string {
	return "DROP SCHEMA IF EXISTS " + schema
}

func (platform *MsSQLPlatform) getDefaultDeclaration(expression string) string {
	return _getDefaultDeclaration(expression)
}

func (platform *MsSQLPlatform) getSchemaAccessName(schema string, name string) string {
	return schema + "." + name
}

func (platform *MsSQLPlatform) getSchemaCommentDeclaration(schema string, expression string) string {
	return ""
}

func (platform *MsSQLPlatform) getTableChecksDeclaration(expressions []string) []string {
	return _getTableChecksDeclaration(expressions)
}

func (platform *MsSQLPlatform) getTableReferencesDeclarationSQL(schema string, foreignKeys []ForeignKey) []string {
	return _getTableReferencesDeclarationSQL(platform, schema, foreignKeys)
}

func (platform *MsSQLPlatform) getTableCommentDeclarationSQL(name string, expression string) string {
	return ""
}

func (platform *MsSQLPlatform) buildTableCreateSQL(schema string, table *Table) (tableString string) {
	return _buildTableCreateSQL(platform, schema, table)
}

func (platform *MsSQLPlatform) getTableDropSQL(schema string, table string) (tableString string) {
	return _getTableDropSQL(platform, schema, table)
}

func (platform *MsSQLPlatform) getSequenceCreateSQL(sequence string) string {
	return "CREATE SEQUENCE " + sequence
}

func (platform *MsSQLPlatform) getSequenceDropSQL(sequence string) string {
	return "DROP SEQUENCE " + sequence
}
