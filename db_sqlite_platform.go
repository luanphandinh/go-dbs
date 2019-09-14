package dbs

import "fmt"

const SQLITE3 string = "sqlite3"

type SqlitePlatform struct {
}

func (platform *SqlitePlatform) GetDriverName() string {
	return SQLITE3
}

func (platform *SqlitePlatform) GetDBConnectionString(server string, port int, user string, password string, dbName string) string {
	return dbName
}

func (platform *SqlitePlatform) ChainCommands(commands ...string) string {
	return concatStrings(commands, ";\n")
}

func (platform *SqlitePlatform) GetSchemaDeclarationSQL(schema string) string {
	return ""
}

func (platform *SqlitePlatform) GetTypeDeclaration(col *Column) string {
	dbType := col.Type
	if inStringArray(col.Type, integerTypes) {
		dbType = "INTEGER"
	}

	if col.Length > 0 {
		return fmt.Sprintf("%s(%d)", dbType, col.Length)
	}

	return dbType
}

func (platform *SqlitePlatform) GetUniqueDeclaration() string {
	return _getUniqueDeclaration()
}

func (platform *SqlitePlatform) GetNotNullDeclaration() string {
	return _getNotNullDeclaration()
}

func (platform *SqlitePlatform) GetPrimaryDeclaration(key []string) string {
	return _getPrimaryDeclaration(key)
}

func (platform *SqlitePlatform) GetAutoIncrementDeclaration() string {
	return ""
}

func (platform *SqlitePlatform) GetUnsignedDeclaration() string {
	return ""
}

func (platform *SqlitePlatform) GetDefaultDeclaration(expression string) string {
	return _getDefaultDeclaration(expression)
}

func (platform *SqlitePlatform) GetColumnCommentDeclaration(expression string) string {
	return ""
}

func (platform *SqlitePlatform) GetColumnsCommentDeclaration(schema string, table *Table) []string {
	return make([]string, 0)
}

func (platform *SqlitePlatform) BuildColumnsDeclarationSQL(cols []*Column) []string {
	return _buildColumnsDeclarationSQL(platform, cols)
}

func (platform *SqlitePlatform) GetColumnCheckDeclaration(expression string) string {
	return _getColumnCheckDeclaration(expression)
}

func (platform *SqlitePlatform) BuildColumnDeclarationSQL(col *Column) string {
	return _buildColumnDeclarationSQL(platform, col)
}

func (platform *SqlitePlatform) BuildSchemaCreateSQL(schema *Schema) string {
	return ""
}

func (platform *SqlitePlatform) GetSchemaCreateDeclarationSQL(schema string) string {
	return ""
}

func (platform *SqlitePlatform) GetSchemaDropDeclarationSQL(schema string) string {
	return ""
}

func (platform *SqlitePlatform) GetSchemaAccessName(schema string, name string) string {
	return name
}

func (platform *SqlitePlatform) GetSchemaCommentDeclaration(schema string, expression string) string {
	return ""
}

func (platform *SqlitePlatform) GetTableChecksDeclaration(expressions []string) []string {
	return _getTableChecksDeclaration(expressions)
}

func (platform *SqlitePlatform) GetTableReferencesDeclarationSQL(schema string, foreignKeys []ForeignKey) []string {
	return _getTableReferencesDeclarationSQL(platform, schema, foreignKeys)
}

func (platform *SqlitePlatform) GetTableCommentDeclarationSQL(name string, expression string) string {
	return ""
}

func (platform *SqlitePlatform) BuildTableCreateSQL(schema string, table *Table) (tableString string) {
	return _buildTableCreateSQL(platform, schema, table)
}

func (platform *SqlitePlatform) GetTableDropSQL(schema string, table string) (tableString string) {
	return _getTableDropSQL(platform, schema, table)
}

func (platform *SqlitePlatform) GetSequenceCreateSQL(sequence string) string {
	return ""
}

func (platform *SqlitePlatform) GetSequenceDropSQL(sequence string) string {
	return ""
}