package dbs

import "fmt"

type PostgresPlatform struct {
}

func (platform *PostgresPlatform) GetSchemaDeclarationSQL(schema string) string {
	return schema
}

func (platform *PostgresPlatform) GetTypeDeclaration(col *Column) string {
	return col.Type
}

func (platform *PostgresPlatform) GetUniqueDeclaration() string {
	return _getUniqueDeclaration()
}

func (platform *PostgresPlatform) GetNotNullDeclaration() string {
	return _getNotNullDeclaration()
}

func (platform *PostgresPlatform) GetPrimaryDeclaration(key []string) string {
	return _getPrimaryDeclaration(key)
}

func (platform *PostgresPlatform) GetAutoIncrementDeclaration() string {
	return ""
}

func (platform *PostgresPlatform) GetUnsignedDeclaration() string {
	return _getUnsignedDeclaration()
}

func (platform *PostgresPlatform) GetColumnDeclarationSQL(col *Column) string {
	columnString := fmt.Sprintf("%s %s", col.Name, platform.GetTypeDeclaration(col))

	if col.NotNull {
		columnString += " " + platform.GetNotNullDeclaration()
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

func (platform *PostgresPlatform) GetColumnsDeclarationSQL(cols []Column) string {
	return _getColumnsDeclarationSQL(platform, cols)
}

func (platform *PostgresPlatform) GetColumnCheckDeclaration(expression string) string {
	return _getColumnCheckDeclaration(expression)
}

func (platform *PostgresPlatform) GetSchemaCreateDeclarationSQL(schema *Schema) string {
	return _getSchemaCreateDeclarationSQL(schema)
}

func (platform *PostgresPlatform) GetSchemaDropDeclarationSQL(schema *Schema) string {
	return _getSchemaDropDeclarationSQL(schema)
}

func (platform *PostgresPlatform) GetDefaultDeclaration(expression string) string {
	return _getDefaultDeclaration(expression)
}

func (platform *PostgresPlatform) GetTableName(schema string, table string) string {
	return fmt.Sprintf("%s.%s", schema, table)
}

func (platform *PostgresPlatform) GetTableCreateSQL(schema string, table *Table) (tableString string) {
	return _getTableCreateSQL(platform, schema, table)
}

func (platform *PostgresPlatform) GetTableDropSQL(schema string, table string) (tableString string) {
	return _getTableDropSQL(platform, schema, table)
}


