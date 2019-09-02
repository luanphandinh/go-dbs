package dbs

import "testing"

func TestColumnNameValidate(t *testing.T) {
	id := Column{}
	assertHasErrorMessage(t, id.Validate(), "column name should not empty")
}

func TestColumnTypeValidate(t *testing.T) {
	for _, colType := range allTypes {
		test := Column{Type: colType}
		if err := test.ValidateType(); err != nil {
			t.Fail()
		}
	}

	test := Column{Type: "WRONG"}
	assertHasError(t, test.ValidateType())

	id := Column{}
	id.Name = "id"
	assertHasErrorMessage(t, id.ValidateType(), "column type should not empty")

	id.Type = "SOMETHING"
	assertHasErrorMessage(t, id.ValidateType(), "incorrect type name")
}

func TestColumnValidateAutoIncrement(t *testing.T) {
	id := Column{Name: "id", Type: TEXT, AutoIncrement: true}
	assertHasErrorMessage(t, id.Validate(), "TEXT can not auto_increment")

	id.Type = INT
	assertHasErrorMessage(t, id.Validate(), "auto_increment must not null")

	id.NotNull = true
	for _, integerType := range integerTypes {
		id.Type = integerType
		assertNotHasError(t, id.Validate())
	}

	for _, floatingType := range floatingTypes {
		id.Type = floatingType
		assertNotHasError(t, id.Validate())
	}
}
