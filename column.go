package dbs

import (
	"fmt"
)

type Column struct {
	Name          string `json:"name"`
	Type          string `json:"type"`
	NotNull       bool   `json:"not_null"`
	Primary       bool   `json:"primary"`
	AutoIncrement bool   `json:"auto_increment"`
}

func (col *Column) ValidateName() error  {
	if col.Name == "" {
		return fmt.Errorf("column name should not empty")
	}

	return nil
}

func (col *Column) ValidateType() error  {
	if col.Type == "" {
		return fmt.Errorf("column type should not empty")
	}

	return nil
}

func (col *Column) ValidateAutoIncrement() error  {
	if col.Type != "INT" && col.AutoIncrement {
		return fmt.Errorf("%s can not auto_increment", col.Type)
	}

	return nil
}

func (col *Column) Validate() error {
	if err := col.ValidateName(); err != nil {
		return err
	}

	if err := col.ValidateType(); err != nil {
		return err
	}

	if err := col.ValidateAutoIncrement(); err != nil {
		return err
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
