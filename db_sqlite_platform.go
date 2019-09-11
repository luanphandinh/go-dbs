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
	return "AUTOINCREMENT"
}

func (platform *SqlitePlatform) GetUnsignedDeclaration() string {
	return _getUnsignedDeclaration()
}

func (platform *SqlitePlatform) GetDefaultDeclaration(expression string) string {
	return _getDefaultDeclaration(expression)
}

func (platform *SqlitePlatform) GetColumnCommentDeclaration(expression string) string {
	return fmt.Sprintf("-- %s", expression)
}

func (platform *SqlitePlatform) GetColumnsDeclarationSQL(cols []Column) []string {
	return _getColumnsDeclarationSQL(platform, cols)
}

func (platform *SqlitePlatform) GetColumnCheckDeclaration(expression string) string {
	return _getColumnCheckDeclaration(expression)
}

func (platform *SqlitePlatform) GetColumnDeclarationSQL(col *Column) string {
	columnString := fmt.Sprintf("%s %s", col.Name, platform.GetTypeDeclaration(col))

	if col.NotNull {
		columnString += " " + platform.GetNotNullDeclaration()
	}

	if col.Unique {
		columnString += " " + platform.GetUniqueDeclaration()
	}

	if col.Check != "" {
		columnString += " " + platform.GetColumnCheckDeclaration(col.Check)
	}

	if col.Default != "" {
		columnString += " " + platform.GetDefaultDeclaration(col.Default)
	}

	if col.Comment != "" {
		columnString += " " + platform.GetColumnCommentDeclaration(col.Comment)
	}

	return columnString
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

func (platform *SqlitePlatform) GetTableCheckDeclaration(expressions []string) string {
	return _getTableCheckDeclaration(expressions)
}

func (platform *SqlitePlatform) GetTableCreateSQL(schema string, table *Table) (tableString string) {
	return _getTableCreateSQL(platform, schema, table)
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