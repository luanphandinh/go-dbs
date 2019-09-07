package dbs

import "fmt"

type SqlitePlatform struct {

}

func (platform *SqlitePlatform) GetUniqueDeclaration() string {
	return "UNIQUE"
}

func (platform *SqlitePlatform) GetNotNullDeclaration() string {
	return "NOT NULL"
}

func (platform *SqlitePlatform) GetPrimaryDeclaration() string {
	return "PRIMARY KEY"
}

func (platform *SqlitePlatform) GetAutoIncrementDeclaration() string {
	return "AUTOINCREMENT"
}

func (platform *SqlitePlatform) GetUnsignedDeclaration() string {
	return "UNSIGNED"
}

func (platform *SqlitePlatform) GetColumnDeclarationSQL(col *Column) string {
	name := col.Name
	dbType := col.Type
	if inStringArray(col.Type, integerTypes) {
		dbType = "INTEGER"
	}

	columnString := fmt.Sprintf("%s %s", name, dbType)

	return columnString
}

func (platform *SqlitePlatform) GetTableCreateSQL(table *Table) (tableString string) {
	// tableString = fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (", table.Name)
	cols := ""
	for index, col := range table.Columns {
		if index == 0 {
			cols += fmt.Sprintf("%s", platform.GetColumnDeclarationSQL(&col))
		} else {
			cols += fmt.Sprintf(", %s", platform.GetColumnDeclarationSQL(&col))
		}
	}

	return fmt.Sprintf(
		"CREATE TABLE IF NOT EXISTS %s (%s, %s)",
		table.Name,
		cols,
		platform.GetPrimaryKeyCreateSQL(table),
	)
}

func (platform *SqlitePlatform) GetPrimaryKeyCreateSQL(table *Table) string {
	return fmt.Sprintf("PRIMARY KEY (%s)", concatString(table.PrimaryKey, ","))
}
