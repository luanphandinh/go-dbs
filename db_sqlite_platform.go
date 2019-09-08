package dbs

import "fmt"

type SqlitePlatform struct {

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

func (platform *SqlitePlatform) GetColumnsDeclarationSQL(cols []Column) string {
	return _getColumnsDeclarationSQL(platform, cols)
}

func (platform *SqlitePlatform) GetColumnCheckDeclaration(expression string) string {
	return _getColumnCheckDeclaration(expression)
}

func (platform *SqlitePlatform) GetColumnDeclarationSQL(col *Column) string {
	columnString := fmt.Sprintf("%s %s", col.Name, platform.GetTypeDeclaration(col))

	if col.Unique {
		columnString += " " + platform.GetUniqueDeclaration()
	}

	if col.Default != "" {
		columnString += " " + platform.GetDefaultDeclaration(col.Default)
	}

	if col.Check != "" {
		columnString += " " + platform.GetColumnCheckDeclaration(col.Check)
	}

	return columnString
}

func (platform *SqlitePlatform) GetSchemaCreateDeclarationSQL(schema *Schema) string {
	return ""
}

func (platform *SqlitePlatform) GetSchemaDropDeclarationSQL(schema *Schema) string {
	return ""
}

func (platform *SqlitePlatform) GetTableName(schema string, table* Table) string {
	return table.Name
}

func (platform *SqlitePlatform) GetTableCreateSQL(schema string, table *Table) (tableString string) {
	return _getTableCreateSQL(platform, schema, table)
}

func (platform *SqlitePlatform) GetTableDropSQL(schema string, table *Table) (tableString string) {
	return _getTableDropSQL(platform, schema, table)
}