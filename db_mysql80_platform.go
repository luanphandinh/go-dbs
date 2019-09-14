package dbs

import "fmt"

const MYSQL80 string = "mysql:8.0"

type MySql80Platform struct {
}

func (platform *MySql80Platform) GetDriverName() string {
	return MYSQL
}

func (platform *MySql80Platform) GetDBConnectionString(server string, port int, user string, password string, dbName string) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s)/%s",
		user,
		password,
		server,
		dbName,
	)
}

func (platform *MySql80Platform) GetTypeDeclaration(col *Column) string {
	if col.Length > 0 {
		return fmt.Sprintf("%s(%d)", col.Type, col.Length)
	}

	return col.Type
}

func (platform *MySql80Platform) ChainCommands(commands ...string) string {
	return concatStrings(commands, "\n")
}

func (platform *MySql80Platform) GetUniqueDeclaration() string {
	return _getUniqueDeclaration()
}

func (platform *MySql80Platform) GetNotNullDeclaration() string {
	return _getNotNullDeclaration()
}

func (platform *MySql80Platform) GetPrimaryDeclaration(key []string) string {
	return _getPrimaryDeclaration(key)
}

func (platform *MySql80Platform) GetAutoIncrementDeclaration() string {
	return "AUTO_INCREMENT"
}

func (platform *MySql80Platform) GetUnsignedDeclaration() string {
	return "UNSIGNED"
}

func (platform *MySql80Platform) GetDefaultDeclaration(expression string) string {
	return _getDefaultDeclaration(expression)
}

func (platform *MySql80Platform) GetColumnCommentDeclaration(expression string) string {
	return fmt.Sprintf("COMMENT '%s'", expression)
}

func (platform *MySql80Platform) GetColumnsCommentDeclaration(schema string, table *Table) []string {
	return make([]string, 0)
}

func (platform *MySql80Platform) GetColumnCheckDeclaration(expression string) string {
	return _getColumnCheckDeclaration(expression)
}

func (platform *MySql80Platform) BuildColumnDeclarationSQL(col *Column) string {
	return _buildColumnDeclarationSQL(platform, col)
}

func (platform *MySql80Platform) BuildColumnsDeclarationSQL(cols []Column) []string {
	return _buildColumnsDeclarationSQL(platform, cols)
}

func (platform *MySql80Platform) BuildSchemaCreateSQL(schema *Schema) string {
	return ""
}

func (platform *MySql80Platform) GetSchemaCreateDeclarationSQL(schema string) string {
	return ""
}

func (platform *MySql80Platform) GetSchemaDropDeclarationSQL(schema string) string {
	return ""
}

func (platform *MySql80Platform) GetSchemaAccessName(schema string, name string) string {
	return name
}

func (platform *MySql80Platform) GetSchemaCommentDeclaration(schema string, expression string) string {
	return ""
}

func (platform *MySql80Platform) GetTableChecksDeclaration(expressions []string) []string {
	return _getTableChecksDeclaration(expressions)
}

func (platform *MySql80Platform) GetTableReferencesDeclarationSQL(schema string, foreignKeys []ForeignKey) []string {
	return _getTableReferencesDeclarationSQL(platform, schema, foreignKeys)
}

func (platform *MySql80Platform) GetTableCommentDeclarationSQL(name string, expression string) string {
	return fmt.Sprintf("COMMENT '%s'", expression)
}

func (platform *MySql80Platform) BuildTableCreateSQL(schema string, table *Table) (tableString string) {
	return _buildTableCreateSQL(platform, schema, table)
}

func (platform *MySql80Platform) GetTableDropSQL(schema string, table string) (tableString string) {
	return _getTableDropSQL(platform, schema, table)
}

func (platform *MySql80Platform) GetSequenceCreateSQL(sequence string) string {
	return ""
}

func (platform *MySql80Platform) GetSequenceDropSQL(sequence string) string {
	return ""
}
