package dbs

import "fmt"

const (
	MYSQL   string = "mysql"
	MYSQL57 string = "mysql:5.7"
)

type MySql57Platform struct {
}

func (platform *MySql57Platform) GetDriverName() string {
	return MYSQL
}

func (platform *MySql57Platform) GetDBConnectionString(server string, port int, user string, password string, dbName string) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s)/%s",
		user,
		password,
		server,
		dbName,
	)
}

func (platform *MySql57Platform) GetTypeDeclaration(col *Column) string {
	if col.Length > 0 {
		return fmt.Sprintf("%s(%d)", col.Type, col.Length)
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
	return _getUnsignedDeclaration()
}

func (platform *MySql57Platform) GetDefaultDeclaration(expression string) string {
	return _getDefaultDeclaration(expression)
}

func (platform *MySql57Platform) GetColumnCheckDeclaration(expression string) string {
	return _getColumnCheckDeclaration(expression)
}

func (platform *MySql57Platform) GetColumnDeclarationSQL(col *Column) string {
	return _getColumnDeclarationSQL(platform, col)
}

func (platform *MySql57Platform) GetColumnsDeclarationSQL(cols []Column) string {
	return _getColumnsDeclarationSQL(platform, cols)
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

func (platform *MySql57Platform) GetTableCheckDeclaration(expressions []string) string {
	return _getTableCheckDeclaration(expressions)
}

func (platform *MySql57Platform) GetTableCreateSQL(schema string, table *Table) (tableString string) {
	return _getTableCreateSQL(platform, schema, table)
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
