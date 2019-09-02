package dbs

import "testing"

func TestTableValidate(t *testing.T) {
	user := Table{}
	assertHasErrorMessage(t, user.Validate(), "table name should not empty")

	user.Name = "user"
	user.Columns = []Column{{}}
	assertHasErrorMessage(t, user.Validate(), "column name should not empty")
}
