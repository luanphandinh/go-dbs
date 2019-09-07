package dbs

import "fmt"

type MySqlPlatform struct {
}

func (platform *MySqlPlatform) GetTypeDeclaration(col *Column) string {
	if col.Length > 0 {
		return fmt.Sprintf("%s(%d)", col.Type, col.Length)
	}

	return col.Type
}

func (platform *MySqlPlatform) GetUniqueDeclaration() string {
	return "UNIQUE"
}

func (platform *MySqlPlatform) GetNotNullDeclaration() string {
	return "NOT NULL"
}

func (platform *MySqlPlatform) GetPrimaryDeclaration() string {
	return "PRIMARY KEY"
}

func (platform *MySqlPlatform) GetAutoIncrementDeclaration() string {
	return "AUTO_INCREMENT"
}

func (platform *MySqlPlatform) GetUnsignedDeclaration() string {
	return "UNSIGNED"
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

	return columnString
}

func (platform *MySqlPlatform) GetTableCreateSQL(table *Table) (tableString string) {
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

func (platform *MySqlPlatform) GetPrimaryKeyCreateSQL(table *Table) string {
	return fmt.Sprintf("PRIMARY KEY (%s)", concatString(table.PrimaryKey, ","))
}
