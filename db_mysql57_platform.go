package dbs

import (
	"database/sql"
	"strconv"
)

const (
	mysql   string = "mysql"
	mysql57 string = "mysql:5.7"
)

type dbMySQL57Platform struct{}

func (platform *dbMySQL57Platform) getDriverName() string {
	return mysql
}

func (platform *dbMySQL57Platform) getDBConnectionString(server string, port int, user string, password string, dbName string) string {
	return user + ":" + password + "@tcp(" + server + ")/" + dbName
}

func (platform *dbMySQL57Platform) chainCommands(commands ...string) string {
	return concatStrings(commands, "\n")
}

func (platform *dbMySQL57Platform) getTypeDeclaration(col *Column) string {
	if col.Length > 0 {
		return col.Type + "(" + strconv.Itoa(col.Length) + ")"
	}

	return col.Type
}

func (platform *dbMySQL57Platform) getUniqueDeclaration() string {
	return _getUniqueDeclaration()
}

func (platform *dbMySQL57Platform) getNotNullDeclaration() string {
	return _getNotNullDeclaration()
}

func (platform *dbMySQL57Platform) getPrimaryDeclaration(key []string) string {
	return _getPrimaryDeclaration(key)
}

func (platform *dbMySQL57Platform) getAutoIncrementDeclaration() string {
	return "AUTO_INCREMENT"
}

func (platform *dbMySQL57Platform) getUnsignedDeclaration() string {
	return "UNSIGNED"
}

func (platform *dbMySQL57Platform) getDefaultDeclaration(expression string) string {
	return _getDefaultDeclaration(expression)
}

func (platform *dbMySQL57Platform) getColumnCommentDeclaration(expression string) string {
	return "COMMENT '" + expression + "'"
}

func (platform *dbMySQL57Platform) getColumnsCommentDeclaration(schema string, table *Table) []string {
	return make([]string, 0)
}

func (platform *dbMySQL57Platform) getColumnCheckDeclaration(expression string) string {
	return _getColumnCheckDeclaration(expression)
}

func (platform *dbMySQL57Platform) buildColumnDeclarationSQL(col *Column) string {
	return _buildColumnDeclarationSQL(platform, col)
}

func (platform *dbMySQL57Platform) buildColumnsDeclarationSQL(cols []*Column) []string {
	return _buildColumnsDeclarationSQL(platform, cols)
}

func (platform *dbMySQL57Platform) buildSchemaCreateSQL(schema *Schema) string {
	return ""
}

func (platform *dbMySQL57Platform) getSchemaCreateDeclarationSQL(schema string) string {
	return ""
}

func (platform *dbMySQL57Platform) getSchemaDropDeclarationSQL(schema string) string {
	return ""
}

func (platform *dbMySQL57Platform) getSchemaAccessName(schema string, name string) string {
	return name
}

func (platform *dbMySQL57Platform) getSchemaCommentDeclaration(schema string, expression string) string {
	return ""
}

func (platform *dbMySQL57Platform) getTableChecksDeclaration(expressions []string) []string {
	return _getTableChecksDeclaration(expressions)
}

func (platform *dbMySQL57Platform) getTableReferencesDeclarationSQL(schema string, foreignKeys []ForeignKey) []string {
	return _getTableReferencesDeclarationSQL(platform, schema, foreignKeys)
}

func (platform *dbMySQL57Platform) getTableCommentDeclarationSQL(name string, expression string) string {
	return "COMMENT '" + expression + "'"
}

func (platform *dbMySQL57Platform) buildTableCreateSQL(schema string, table *Table) (tableString string) {
	return _buildTableCreateSQL(platform, schema, table)
}

func (platform *dbMySQL57Platform) getTableDropSQL(schema string, table string) (tableString string) {
	return _getTableDropSQL(platform, schema, table)
}

func (platform *dbMySQL57Platform) getSequenceCreateSQL(sequence string) string {
	return ""
}

func (platform *dbMySQL57Platform) getSequenceDropSQL(sequence string) string {
	return ""
}

func (platform *dbMySQL57Platform) checkSchemaExistSQL(schema string) string {
	return ""
}

func (platform *dbMySQL57Platform) checkSchemaHasTableSQL(schema string, table string) string {
	return "SHOW TABLES LIKE '" + platform.getSchemaAccessName(schema, table) + "'"
}

func (platform *dbMySQL57Platform) getSchemaTablesSQL(schema string) string {
	return "SHOW TABLES"
}

func (platform *dbMySQL57Platform) getTableColumnsSQL(schema string , table string) string {
	return "SHOW COLUMNS FROM " + platform.getSchemaAccessName(schema, table)
}

func (platform *dbMySQL57Platform) parseTableColumns(rows *sql.Rows) []*Column {
	return make([]*Column, 0)
}
