package dbs

import "fmt"

type Platform interface {
	GetTypeDeclaration(col *Column) string
	GetUniqueDeclaration() string
	GetNotNullDeclaration() string
	GetPrimaryDeclaration(table *Table) string
	GetAutoIncrementDeclaration() string
	GetUnsignedDeclaration() string
	GetColumnDeclarationSQL(col *Column) string
	GetTableCreateSQL(table *Table) string
	GetTableDropSQL(table *Table) string
}

func GetPlatform(platform string) Platform {
	if platform == MYSQL {
		return &MySqlPlatform{}
	}

	if platform == SQLITE3 {
		return &SqlitePlatform{}
	}

	if platform == POSTGRES {
		return &PostgresPlatform{}
	}

	return nil
}

func _getUniqueDeclaration() string {
	return "UNIQUE"
}

func _getNotNullDeclaration() string {
	return "NOT NULL"
}

func _getPrimaryDeclaration(table *Table) string {
	return fmt.Sprintf("PRIMARY KEY (%s)", concatString(table.PrimaryKey, ","))
}

func _getAutoIncrementDeclaration() string {
	return "AUTO_INCREMENT"
}

func _getUnsignedDeclaration() string {
	return "UNSIGNED"
}

func _getTableCreateSQL(platform Platform, table *Table) string {
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
		platform.GetPrimaryDeclaration(table),
	)
}

func _getTableDropSQL(platform Platform, table *Table) string {
	return fmt.Sprintf("DROP TABLE IF EXISTS %s", table.Name)
}
