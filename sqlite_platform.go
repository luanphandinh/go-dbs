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

func (platform *SqlitePlatform) GetColumnSQLDeclaration(col *Column) string {
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

func (platform *SqlitePlatform) GetTableSQLCreate(table *Table) (tableString string) {
	tableString = fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (", table.Name)
	for index, col := range table.Columns {
		if index == 0 {
			tableString += fmt.Sprintf("%s", platform.GetColumnSQLDeclaration(&col))
		} else {
			tableString += fmt.Sprintf(", %s", platform.GetColumnSQLDeclaration(&col))
		}
	}

	return tableString + ")"
}
