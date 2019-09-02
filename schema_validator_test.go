package dbs

import "testing"

func TestSchemaValidate(t *testing.T) {
	tables := []Table{
		{
			Columns: []Column{
				{Name: "id", Type: INT, Primary: true, NotNull: true, AutoIncrement: true},
				{Name: "name", Type: TEXT, NotNull: true},
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
