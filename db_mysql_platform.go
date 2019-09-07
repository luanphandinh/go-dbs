package dbs

import "fmt"

type MySqlPlatform struct {
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
	columnString := fmt.Sprintf("%s %s", col.Name, col.Type)

	if col.Unsigned {
		columnString += " " + platform.GetUnsignedDeclaration()
	}

	if col.NotNull {
		columnString += " " + platform.GetNotNullDeclaration()
	}

	if col.AutoIncrement {
		columnString += " " + platform.GetAutoIncrementDeclaration()
	}

	return columnString
}

func (platform *MySqlPlatform) GetTableCreateSQL(table *Table) (tableString string) {
	tableString = fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (", table.Name)
	for index, col := range table.Columns {
		if index == 0 {
			tableString += fmt.Sprintf("%s", platform.GetColumnDeclarationSQL(&col))
		} else {
			tableString += fmt.Sprintf(", %s", platform.GetColumnDeclarationSQL(&col))
		}
	}

	return tableString + ")"
}

func (platform *MySqlPlatform) GetPrimaryKeyCreateSQL(table *Table) string {
	return fmt.Sprintf("ALTER TABLE %s ADD PRIMARY KEY (%s)", table.Name, concatString(table.PrimaryKey, ","))
}
