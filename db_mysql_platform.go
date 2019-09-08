package dbs

import "fmt"

type MySqlPlatform struct {
}

// Column declarations

func (platform *MySqlPlatform) GetTypeDeclaration(col *Column) string {
	if col.Length > 0 {
		return fmt.Sprintf("%s(%d)", col.Type, col.Length)
	}

	return col.Type
}

func (platform *MySqlPlatform) GetUniqueDeclaration() string {
	return _getUniqueDeclaration()
}

func (platform *MySqlPlatform) GetNotNullDeclaration() string {
	return _getNotNullDeclaration()
}

func (platform *MySqlPlatform) GetPrimaryDeclaration(table *Table) string {
	return _getPrimaryDeclaration(table)
}

func (platform *MySqlPlatform) GetAutoIncrementDeclaration() string {
	return _getAutoIncrementDeclaration()
}

func (platform *MySqlPlatform) GetUnsignedDeclaration() string {
	return _getUnsignedDeclaration()
}

func (platform *MySqlPlatform) GetDefaultDeclaration(expression string) string {
	return _getDefaultDeclaration(expression)
}

func (platform *MySqlPlatform) GetColumnCheckDeclaration(expression string) string {
	return _getColumnCheckDeclaration(expression)
}

func (platform *MySqlPlatform) GetColumnDeclarationSQL(col *Column) string {
	columnString := fmt.Sprintf("%s %s", col.Name, platform.GetTypeDeclaration(col))

	if col.Unsigned {
		columnString += " " + platform.GetUnsignedDeclaration()
	}

	if col.NotNull {
		columnString += " " + platform.GetNotNullDeclaration()
	}

	if col.AutoIncrement {
		columnString += " " + platform.GetAutoIncrementDeclaration()
	}

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

func (platform *MySqlPlatform) GetColumnsDeclarationSQL(cols []Column) string {
	return _getColumnsDeclarationSQL(platform, cols)
}

func (platform *MySqlPlatform) GetSchemaCreateDeclarationSQL(schema *Schema) string {
	return ""
}

func (platform *MySqlPlatform) GetSchemaDropDeclarationSQL(schema *Schema) string {
	return ""
}

func (platform *MySqlPlatform) GetTableName(schema string, table* Table) string {
	return table.Name
}

func (platform *MySqlPlatform) GetTableCreateSQL(schema string, table *Table) (tableString string) {
	return _getTableCreateSQL(platform, schema, table)
}

func (platform *MySqlPlatform) GetTableDropSQL(schema string, table *Table) (tableString string) {
	return _getTableDropSQL(platform, schema, table)
}

