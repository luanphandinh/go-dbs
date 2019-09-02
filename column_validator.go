package dbs

import (
	"fmt"
)

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

	if !col.isOneOf(allTypes) {
		return fmt.Errorf("incorrect type name")
	}

	return nil
}

func (col *Column) ValidateAutoIncrement() error {
	if !col.isOneOf(integerTypes) && !col.isOneOf(floatingTypes) && col.AutoIncrement {
		return fmt.Errorf("%s can not auto_increment", col.Type)
	}

	if !col.NotNull {
		return fmt.Errorf("auto_increment must not null")
	}

	return nil
}

func (col *Column) ValidateUnsigned() error {
	if !col.isOneOf(integerTypes) && col.Unsigned {
		return fmt.Errorf("only integer types can be unsigned")
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

	if err := col.ValidateUnsigned(); err != nil {
		return err
	}

	return nil
}

// Temporary use linear search
func (col *Column) isOneOf(types []string) bool  {
	for _, dbType := range types {
		if col.Type == dbType {
			return true
		}
	}

	return false
}
