package schema

import (
	"fmt"
	"log"
)

type Schema struct {
	Name   string  `json:"name"`
	Tables []Table `json:"tables"`
}

type Table struct {
	Name    string   `json:"name"`
	Columns []Column `json:"columns"`
}

type Column struct {
	Name          string `json:"name"`
	Type          string `json:"type"`
	NotNull       bool   `json:"not_null"`
	Primary       bool   `json:"primary"`
	AutoIncrement bool   `json:"auto_increment"`
}


func (col *Column) Validate() error {
	if col.Name == "" {
		return fmt.Errorf("column name should not empty")
	}

	if col.Type == "" {
		return fmt.Errorf("column type should not empty")
	}

	if col.Type != "INT" && col.AutoIncrement {
		log.Fatal(fmt.Sprintf("%s can not auto_increment", col.Type))
	}

	return nil
}

func (col *Column) ToString() string {
	columnString := fmt.Sprintf("%s %s", col.Name, col.Type)

	if col.AutoIncrement {
		columnString += " AUTO_INCREMENT"
	}

	if col.Primary {
		columnString += " PRIMARY KEY"
	}

	if col.NotNull {
		columnString += " NOT NULL"
	}

	return columnString
}

func (table *Table) Validate() error  {
	if table.Name == "" {
		return fmt.Errorf("table name should not empty")
	}

	for _, col := range table.Columns {
		if err := col.Validate(); err != nil {
			return err
		}
	}

	return nil
}

func (table *Table) ToString() (tableString string) {
	tableString = fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (", table.Name)
	cols := table.Columns
	for i := 0; i < len(cols); i++ {
		if i == 0 {
			tableString += fmt.Sprintf("%s", cols[i].ToString())
		} else {
			tableString += fmt.Sprintf(", %s", cols[i].ToString())
		}
	}

	return tableString + ")"
}
