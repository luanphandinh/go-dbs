package dbs

import "strconv"

const mysql80 string = "mysql:8.0"

type dbMySQL80Platform struct{}

func (platform *dbMySQL80Platform) getDriverName() string {
	return mysql
}

func (platform *dbMySQL80Platform) getDBConnectionString(server string, port int, user string, password string, dbName string) string {
	return user + ":" + password + "@tcp(" + server + ")/" + dbName
}

func (platform *dbMySQL80Platform) getTypeDeclaration(col *Column) string {
	if col.Length > 0 {
		return col.Type + "(" + strconv.Itoa(col.Length) + ")"
	}

	return col.Type
}

func (platform *dbMySQL80Platform) chainCommands(commands ...string) string {
	return concatStrings(commands, "\n")
}

func (platform *dbMySQL80Platform) getUniqueDeclaration() string {
	return _getUniqueDeclaration()
}

func (platform *dbMySQL80Platform) getNotNullDeclaration() string {
	return _getNotNullDeclaration()
}

func (platform *dbMySQL80Platform) getPrimaryDeclaration(key []string) string {
	return _getPrimaryDeclaration(key)
}

func (platform *dbMySQL80Platform) getAutoIncrementDeclaration() string {
	return "AUTO_INCREMENT"
}

func (platform *dbMySQL80Platform) getUnsignedDeclaration() string {
	return "UNSIGNED"
}

func (platform *dbMySQL80Platform) getDefaultDeclaration(expression string) string {
	return _getDefaultDeclaration(expression)
}

func (platform *dbMySQL80Platform) getColumnCommentDeclaration(expression string) string {
	return "COMMENT '" + expression + "'"
}

func (platform *dbMySQL80Platform) getColumnsCommentDeclaration(schema string, table *Table) []string {
	return make([]string, 0)
}

func (platform *dbMySQL80Platform) getColumnCheckDeclaration(expression string) string {
	return _getColumnCheckDeclaration(expression)
}

func (platform *dbMySQL80Platform) buildColumnDeclarationSQL(col *Column) string {
	return _buildColumnDeclarationSQL(platform, col)
}

func (platform *dbMySQL80Platform) buildColumnsDeclarationSQL(cols []*Column) []string {
	return _buildColumnsDeclarationSQL(platform, cols)
}

func (platform *dbMySQL80Platform) buildSchemaCreateSQL(schema *Schema) string {
	return ""
}

func (platform *dbMySQL80Platform) getSchemaCreateDeclarationSQL(schema string) string {
	return ""
}

func (platform *dbMySQL80Platform) getSchemaDropDeclarationSQL(schema string) string {
	return ""
}

func (platform *dbMySQL80Platform) getSchemaAccessName(schema string, name string) string {
	return name
}

func (platform *dbMySQL80Platform) getSchemaCommentDeclaration(schema string, expression string) string {
	return ""
}

func (platform *dbMySQL80Platform) getTableChecksDeclaration(expressions []string) []string {
	return _getTableChecksDeclaration(expressions)
}

func (platform *dbMySQL80Platform) getTableReferencesDeclarationSQL(schema string, foreignKeys []ForeignKey) []string {
	return _getTableReferencesDeclarationSQL(platform, schema, foreignKeys)
}

func (platform *dbMySQL80Platform) getTableCommentDeclarationSQL(name string, expression string) string {
	return "COMMENT '" + expression + "'"
}

func (platform *dbMySQL80Platform) buildTableCreateSQL(schema string, table *Table) (tableString string) {
	return _buildTableCreateSQL(platform, schema, table)
}

func (platform *dbMySQL80Platform) getTableDropSQL(schema string, table string) (tableString string) {
	return _getTableDropSQL(platform, schema, table)
}

func (platform *dbMySQL80Platform) getSequenceCreateSQL(sequence string) string {
	return ""
}

func (platform *dbMySQL80Platform) getSequenceDropSQL(sequence string) string {
	return ""
}
