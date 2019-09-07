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

func (platform *PostgresPlatform) GetPrimaryDeclaration(table *Table) string {
	return _getPrimaryDeclaration(table)
}

func (platform *PostgresPlatform) GetAutoIncrementDeclaration() string {
	return _getAutoIncrementDeclaration()
}

func (platform *PostgresPlatform) GetUnsignedDeclaration() string {
	return _getUnsignedDeclaration()
}

func (platform *PostgresPlatform) GetTableName(schema string, table *Table) string {
	return fmt.Sprintf("%s.%s", schema, table.Name)
}

func (platform *PostgresPlatform) GetColumnDeclarationSQL(col *Column) string {
	columnString := fmt.Sprintf("%s %s", col.Name, platform.GetTypeDeclaration(col))

	if col.NotNull {
		columnString += " " + platform.GetNotNullDeclaration()
	}

	if col.Unique {
		columnString += " " + platform.GetUniqueDeclaration()
	}

	return columnString
}

func (platform *PostgresPlatform) GetTableCreateSQL(schema string, table *Table) (tableString string) {
	return _getTableCreateSQL(platform, schema, table)
}

func (platform *PostgresPlatform) GetTableDropSQL(schema string, table *Table) (tableString string) {
	return fmt.Sprintf("DROP TABLE IF EXISTS public.%s", table.Name)
}


