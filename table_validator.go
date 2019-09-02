package dbs

import "fmt"

func (table *Table) ValidateName() error {
	if table.Name == "" {
		return fmt.Errorf("table name should not empty")
	}

	return nil
}

func (table *Table) ValidateColumns() error  {
	for _, col := range table.Columns {
		if err := col.Validate(); err != nil {
			return err
		}
	}

	return nil
}

func (table *Table) Validate() error  {
	if err := table.ValidateName(); err != nil {
		return err
	}

	if err := table.ValidateColumns(); err != nil {
		return err
	}

	return nil
}
