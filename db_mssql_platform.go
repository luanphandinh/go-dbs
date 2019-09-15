package dbs

import "strconv"

const MSSQL string = "sqlserver"

type MsSqlPlatform struct{}

func (platform *MsSqlPlatform) GetDriverName() string {
	return MSSQL
}

func (platform *MsSqlPlatform) GetDBConnectionString(server string, port int, user string, password string, dbName string) string {
	info := make([]string, 0)
	info = append(info, "server=" + server)
	info = append(info, "user id=" + user)
	info = append(info, "password=" + password)
	info = append(info, "database=" + dbName)

	return concatStrings(info, ";")
}

func (platform *MsSqlPlatform) ChainCommands(commands ...string) string {
	return concatStrings(commands, ";\nGO\n")
}

func (platform *MsSqlPlatform) GetTypeDeclaration(col *Column) string {
	if col.Length > 0 {
		return col.Type + "(" + strconv.Itoa(col.Length) + ")"
	}

	return col.Type
}

func (platform *MsSqlPlatform) GetUniqueDeclaration() string {
	return _getUniqueDeclaration()
}

func (platform *MsSqlPlatform) GetNotNullDeclaration() string {
	return _getNotNullDeclaration()
}

func (platform *MsSqlPlatform) GetPrimaryDeclaration(key []string) string {
	return _getPrimaryDeclaration(key)
}

func (platform *MsSqlPlatform) GetAutoIncrementDeclaration() string {
	return "IDENTITY(1,1)"
}

func (platform *MsSqlPlatform) GetUnsignedDeclaration() string {
	return ""
}

func (platform *MsSqlPlatform) BuildColumnDeclarationSQL(col *Column) string {
	return _buildColumnDeclarationSQL(platform, col)
}

func (platform *MsSqlPlatform) BuildColumnsDeclarationSQL(cols []*Column) []string {
	return _buildColumnsDeclarationSQL(platform, cols)
}

func (platform *MsSqlPlatform) GetColumnCommentDeclaration(expression string) string {
	return ""
}

func (platform *MsSqlPlatform) GetColumnsCommentDeclaration(schema string, table *Table) []string {
	return make([]string, 0)
}

func (platform *MsSqlPlatform) GetColumnCheckDeclaration(expression string) string {
	return _getColumnCheckDeclaration(expression)
}

func (platform *MsSqlPlatform) BuildSchemaCreateSQL(schema *Schema) string {
	return platform.GetSchemaCreateDeclarationSQL(schema.Name)
}

func (platform *MsSqlPlatform) GetSchemaCreateDeclarationSQL(schema string) string {
	return "CREATE SCHEMA " + schema
}

func (platform *MsSqlPlatform) GetSchemaDropDeclarationSQL(schema string) string {
	return "DROP SCHEMA IF EXISTS " + schema
}

func (platform *MsSqlPlatform) GetDefaultDeclaration(expression string) string {
	return _getDefaultDeclaration(expression)
}

func (platform *MsSqlPlatform) GetSchemaAccessName(schema string, name string) string {
	return schema + "." + name
}

func (platform *MsSqlPlatform) GetSchemaCommentDeclaration(schema string, expression string) string {
	return ""
}

func (platform *MsSqlPlatform) GetTableChecksDeclaration(expressions []string) []string {
	return _getTableChecksDeclaration(expressions)
}

func (platform *MsSqlPlatform) GetTableReferencesDeclarationSQL(schema string, foreignKeys []ForeignKey) []string {
	return _getTableReferencesDeclarationSQL(platform, schema, foreignKeys)
}

func (platform *MsSqlPlatform) GetTableCommentDeclarationSQL(name string, expression string) string {
	return ""
}

func (platform *MsSqlPlatform) BuildTableCreateSQL(schema string, table *Table) (tableString string) {
	return _buildTableCreateSQL(platform, schema, table)
}

func (platform *MsSqlPlatform) GetTableDropSQL(schema string, table string) (tableString string) {
	return _getTableDropSQL(platform, schema, table)
}

func (platform *MsSqlPlatform) GetSequenceCreateSQL(sequence string) string {
	return "CREATE SEQUENCE " + sequence
}

func (platform *MsSqlPlatform) GetSequenceDropSQL(sequence string) string {
	return "DROP SEQUENCE " + sequence
}
