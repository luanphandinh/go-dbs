package dbs

import "fmt"

type Table struct {
	Name    string   `json:"name"`
	Columns []Column `json:"columns"`
}

func (table *Table) GetSQLCreateTable(platform Platform) (tableString string) {
	tableString = fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (", table.Name)
	for index, col := range table.Columns {
		if index == 0 {
			tableString += fmt.Sprintf("%s", col.GetSQLDeclaration(platform))
		} else {
			tableString += fmt.Sprintf(", %s", col.GetSQLDeclaration(platform))
		}
	}

	return tableString + ")"
}
