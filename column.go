package dbs

import "fmt"

type Column struct {
	Name          string `json:"name"`
	Type          string `json:"type"`
	NotNull       bool   `json:"not_null"`
	Primary       bool   `json:"primary"`
	AutoIncrement bool   `json:"auto_increment"`
	Unsigned	  bool   `json:"unsigned"`
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

	if col.Unsigned {
		columnString += " UNSIGNED"
	}

	return columnString
}
