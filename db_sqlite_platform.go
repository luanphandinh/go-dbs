package dbs

import "fmt"

type SqlitePlatform struct {

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

func (platform *SqlitePlatform) GetPrimaryDeclaration(table *Table) string {
	return _getPrimaryDeclaration(table)
}

func (platform *SqlitePlatform) GetAutoIncrementDeclaration() string {
	return "AUTOINCREMENT"
}

func (platform *SqlitePlatform) GetUnsignedDeclaration() string {
	return _getUnsignedDeclaration()
}

func (platform *SqlitePlatform) GetColumnDeclarationSQL(col *Column) string {
	columnString := fmt.Sprintf("%s %s", col.Name, platform.GetTypeDeclaration(col))

	if col.Unique {
		columnString += " " + platform.GetUniqueDeclaration()
	}

	return columnString
}

func (platform *SqlitePlatform) GetTableCreateSQL(table *Table) (tableString string) {
	return _getTableCreateSQL(platform, table)
}

func (platform *SqlitePlatform) GetTableDropSQL(table *Table) (tableString string) {
	return _getTableDropSQL(platform, table)
}