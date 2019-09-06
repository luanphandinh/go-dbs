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

func (col *Column) GetSQLDeclaration(platform Platform) string {
	columnString := fmt.Sprintf("%s %s", col.Name, col.Type)

	if col.Unsigned {
		columnString += " " + platform.GetUnsignedDeclaration()
	}

	if col.NotNull {
		columnString += " " + platform.GetNotNullDeclaration()
	}

	if col.Primary {
		columnString += " " + platform.GetPrimaryDeclaration()
	}

	if col.AutoIncrement {
		columnString += " " + platform.GetAutoIncrementDeclaration()
	}

	return columnString
}
