package dbs

import (
	"fmt"
	"log"
)

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
