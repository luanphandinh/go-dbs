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
	return _getUnsignedDeclaration()
}

func (platform *MySql80Platform) GetDefaultDeclaration(expression string) string {
	return _getDefaultDeclaration(expression)
}

func (platform *MySql80Platform) GetColumnCheckDeclaration(expression string) string {
	return _getColumnCheckDeclaration(expression)
}

func (platform *MySql80Platform) GetColumnDeclarationSQL(col *Column) string {
	return _getColumnDeclarationSQL(platform, col)
}

func (platform *MySql80Platform) GetColumnsDeclarationSQL(cols []Column) string {
	return _getColumnsDeclarationSQL(platform, cols)
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

func (platform *MySql80Platform) GetTableCheckDeclaration(expressions []string) string {
	return _getTableCheckDeclaration(expressions)
}

func (platform *MySql80Platform) GetTableCreateSQL(schema string, table *Table) (tableString string) {
	return _getTableCreateSQL(platform, schema, table)
}

func (platform *MySql80Platform) GetTableDropSQL(schema string, table string) (tableString string) {
	return _getTableDropSQL(platform, schema, table)
}

func (platform *MySql80Platform) GetSequenceCreateSQL(schema string, sequence string) string {
	return ""
}

func (platform *MySql80Platform) GetSequenceDropSQL(schema string, sequence string) string {
	return ""
}
