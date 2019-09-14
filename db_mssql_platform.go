package dbs

import "fmt"

const MSSQL string = "sqlserver"

type MsSqlPlatform struct {
}

func (platform *MsSqlPlatform) GetDriverName() string {
	return MSSQL
}

func (platform *MsSqlPlatform) GetDBConnectionString(server string, port int, user string, password string, dbName string) string {
	return fmt.Sprintf(
		"server=%s;user id=%s;password=%s;database=%s;",
		server,
		user,
		password,
		dbName,
	)
}

func (platform *MsSqlPlatform) ChainCommands(commands ...string) string {
	return concatStrings(commands, ";\nGO\n")
}

func (platform *MsSqlPlatform) GetTypeDeclaration(col *Column) string {
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

func (platform *MsSqlPlatform) BuildColumnsDeclarationSQL(cols []Column) []string {
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
	return fmt.Sprintf("CREATE SCHEMA %s", schema)
}

func (platform *MsSqlPlatform) GetSchemaDropDeclarationSQL(schema string) string {
	return fmt.Sprintf("DROP SCHEMA IF EXISTS %s", schema)
}

func (platform *MsSqlPlatform) GetDefaultDeclaration(expression string) string {
	return _getDefaultDeclaration(expression)
}

func (platform *MsSqlPlatform) GetSchemaAccessName(schema string, name string) string {
	return fmt.Sprintf("%s.%s", schema, name)
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
	return fmt.Sprintf("CREATE SEQUENCE %s", sequence)
}

func (platform *MsSqlPlatform) GetSequenceDropSQL(sequence string) string {
	return fmt.Sprintf("DROP SEQUENCE %s", sequence)
}
