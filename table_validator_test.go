package dbs

import "testing"

func TestTableValidate(t *testing.T) {
	user := Table{}
	if err := user.Validate(); err != nil {
		if err.Error() != "table name should not empty" {
			t.Fail()
		}
	}
	user.Name = "user"

	user.Columns = []Column{{}}
	if err := user.Validate(); err != nil {
		if err.Error() != "column name should not empty" {
			t.Fail()
		}
	}
}
