package dbs

import "strconv"

const (
	MYSQL   string = "mysql"
	MYSQL57 string = "mysql:5.7"
)

type MySql57Platform struct{}

func (platform *MySql57Platform) GetDriverName() string {
	return MYSQL
}

func (platform *MySql57Platform) GetDBConnectionString(server string, port int, user string, password string, dbName string) string {
	return user + ":" + password + "@tcp(" + server + ")/" + dbName
}

func (platform *MySql57Platform) ChainCommands(commands ...string) string {
	return concatStrings(commands, "\n")
}

func (platform *MySql57Platform) GetTypeDeclaration(col *Column) string {
	if col.Length > 0 {
		return col.Type + "(" + strconv.Itoa(col.Length) + ")"
	}

	return col.Type
}

func (platform *MySql57Platform) GetUniqueDeclaration() string {
	return _getUniqueDeclaration()
}

func (platform *MySql57Platform) GetNotNullDeclaration() string {
	return _getNotNullDeclaration()
}

func (platform *MySql57Platform) GetPrimaryDeclaration(key []string) string {
	return _getPrimaryDeclaration(key)
}

func (platform *MySql57Platform) GetAutoIncrementDeclaration() string {
	return "AUTO_INCREMENT"
}

func (platform *MySql57Platform) GetUnsignedDeclaration() string {
	return "UNSIGNED"
}

func (platform *MySql57Platform) GetDefaultDeclaration(expression string) string {
	return _getDefaultDeclaration(expression)
}

func (platform *MySql57Platform) GetColumnCommentDeclaration(expression string) string {
	return "COMMENT '" + expression + "'"
}

func (platform *MySql57Platform) GetColumnsCommentDeclaration(schema string, table *Table) []string {
	return make([]string, 0)
}

func (platform *MySql57Platform) GetColumnCheckDeclaration(expression string) string {
	return _getColumnCheckDeclaration(expression)
}

func (platform *MySql57Platform) BuildColumnDeclarationSQL(col *Column) string {
	return _buildColumnDeclarationSQL(platform, col)
}

func (platform *MySql57Platform) BuildColumnsDeclarationSQL(cols []*Column) []string {
	return _buildColumnsDeclarationSQL(platform, cols)
}

func (platform *MySql57Platform) BuildSchemaCreateSQL(schema *Schema) string {
	return ""
}

func (platform *MySql57Platform) GetSchemaCreateDeclarationSQL(schema string) string {
	return ""
}

func (platform *MySql57Platform) GetSchemaDropDeclarationSQL(schema string) string {
	return ""
}

func (platform *MySql57Platform) GetSchemaAccessName(schema string, name string) string {
	return name
}

func (platform *MySql57Platform) GetSchemaCommentDeclaration(schema string, expression string) string {
	return ""
}

func (platform *MySql57Platform) GetTableChecksDeclaration(expressions []string) []string {
	return _getTableChecksDeclaration(expressions)
}

func (platform *MySql57Platform) GetTableReferencesDeclarationSQL(schema string, foreignKeys []ForeignKey) []string {
	return _getTableReferencesDeclarationSQL(platform, schema, foreignKeys)
}

func (platform *MySql57Platform) GetTableCommentDeclarationSQL(name string, expression string) string {
	return "COMMENT '" + expression + "'"
}

func (platform *MySql57Platform) BuildTableCreateSQL(schema string, table *Table) (tableString string) {
	return _buildTableCreateSQL(platform, schema, table)
}

func (platform *MySql57Platform) GetTableDropSQL(schema string, table string) (tableString string) {
	return _getTableDropSQL(platform, schema, table)
}

func (platform *MySql57Platform) GetSequenceCreateSQL(sequence string) string {
	return ""
}

func (platform *MySql57Platform) GetSequenceDropSQL(sequence string) string {
	return ""
}
