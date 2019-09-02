package dbs

import "fmt"

func (col *Column) ValidateName() error {
	if col.Name == "" {
		return fmt.Errorf("column name should not empty")
	}

	return nil
}

func (col *Column) ValidateType() error {
	if col.Type == "" {
		return fmt.Errorf("column type should not empty")
	}

	found := false
	for _, dbType := range allTypes {
		if col.Type == dbType {
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("incorrect type name")
	}

	return nil
}

func (col *Column) ValidateAutoIncrement() error {
	if col.Type != INT && col.AutoIncrement {
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
