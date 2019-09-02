package dbs

import "testing"

func TestTableValidate(t *testing.T) {
	user := Table{}
	assertHasErrorMessage(t, "table name should not empty", user.Validate())

	user.Name = "user"
	user.Columns = []Column{{}}
	assertHasErrorMessage(t, "column name should not empty", user.Validate())
}
