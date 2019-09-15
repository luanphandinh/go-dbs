package dbs

import "strconv"

const MYSQL80 string = "mysql:8.0"

type MySQL80Platform struct{}

func (platform *MySQL80Platform) getDriverName() string {
	return MYSQL
}

func (platform *MySQL80Platform) getDBConnectionString(server string, port int, user string, password string, dbName string) string {
	return user + ":" + password + "@tcp(" + server + ")/" + dbName
}

func (platform *MySQL80Platform) getTypeDeclaration(col *Column) string {
	if col.Length > 0 {
		return col.Type + "(" + strconv.Itoa(col.Length) + ")"
	}

	return col.Type
}

func (platform *MySQL80Platform) chainCommands(commands ...string) string {
	return concatStrings(commands, "\n")
}

func (platform *MySQL80Platform) getUniqueDeclaration() string {
	return _getUniqueDeclaration()
}

func (platform *MySQL80Platform) getNotNullDeclaration() string {
	return _getNotNullDeclaration()
}

func (platform *MySQL80Platform) getPrimaryDeclaration(key []string) string {
	return _getPrimaryDeclaration(key)
}

func (platform *MySQL80Platform) getAutoIncrementDeclaration() string {
	return "AUTO_INCREMENT"
}

func (platform *MySQL80Platform) getUnsignedDeclaration() string {
	return "UNSIGNED"
}

func (platform *MySQL80Platform) getDefaultDeclaration(expression string) string {
	return _getDefaultDeclaration(expression)
}

func (platform *MySQL80Platform) getColumnCommentDeclaration(expression string) string {
	return "COMMENT '" + expression + "'"
}

func (platform *MySQL80Platform) getColumnsCommentDeclaration(schema string, table *Table) []string {
	return make([]string, 0)
}

func (platform *MySQL80Platform) getColumnCheckDeclaration(expression string) string {
	return _getColumnCheckDeclaration(expression)
}

func (platform *MySQL80Platform) buildColumnDeclarationSQL(col *Column) string {
	return _buildColumnDeclarationSQL(platform, col)
}

func (platform *MySQL80Platform) buildColumnsDeclarationSQL(cols []*Column) []string {
	return _buildColumnsDeclarationSQL(platform, cols)
}

func (platform *MySQL80Platform) buildSchemaCreateSQL(schema *Schema) string {
	return ""
}

func (platform *MySQL80Platform) getSchemaCreateDeclarationSQL(schema string) string {
	return ""
}

func (platform *MySQL80Platform) getSchemaDropDeclarationSQL(schema string) string {
	return ""
}

func (platform *MySQL80Platform) getSchemaAccessName(schema string, name string) string {
	return name
}

func (platform *MySQL80Platform) getSchemaCommentDeclaration(schema string, expression string) string {
	return ""
}

func (platform *MySQL80Platform) getTableChecksDeclaration(expressions []string) []string {
	return _getTableChecksDeclaration(expressions)
}

func (platform *MySQL80Platform) getTableReferencesDeclarationSQL(schema string, foreignKeys []ForeignKey) []string {
	return _getTableReferencesDeclarationSQL(platform, schema, foreignKeys)
}

func (platform *MySQL80Platform) getTableCommentDeclarationSQL(name string, expression string) string {
	return "COMMENT '" + expression + "'"
}

func (platform *MySQL80Platform) buildTableCreateSQL(schema string, table *Table) (tableString string) {
	return _buildTableCreateSQL(platform, schema, table)
}

func (platform *MySQL80Platform) getTableDropSQL(schema string, table string) (tableString string) {
	return _getTableDropSQL(platform, schema, table)
}

func (platform *MySQL80Platform) getSequenceCreateSQL(sequence string) string {
	return ""
}

func (platform *MySQL80Platform) getSequenceDropSQL(sequence string) string {
	return ""
}
