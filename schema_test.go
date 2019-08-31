package dbs

import "testing"

func TestSchemaValidate(t *testing.T) {
	tables := []Table{
		{
			Columns: []Column{
				{"id", "INT", true, true, true},
				{"name", "NVARCHAR(50)", true, false, false},
			},
		},
	}

	dbSchema := &Schema{
		Name: "workspace",
		Tables: tables,
	}

	if err := dbSchema.Validate(); err != nil {
		if err.Error() != "table name should not empty" {
			t.Fail()
		}
	}

	tables[0].Name = "user"
	if err := dbSchema.Validate(); err != nil {
		t.Fail()
	}
}
