package dbs

import "testing"

func TestColumnNameValidate(t *testing.T) {
	id := Column{}
	assertHasErrorMessage(t, "column name should not empty", id.Validate())
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
	assertHasErrorMessage(t, "column type should not empty", id.ValidateType())

	id.Type = "SOMETHING"
	assertHasErrorMessage(t, "incorrect type name", id.ValidateType())
}

func TestColumnValidateAutoIncrement(t *testing.T) {
	id := Column{Name: "id", Type: TEXT, AutoIncrement: true}
	assertHasErrorMessage(t, "TEXT can not auto_increment", id.Validate())

	id.Type = INT
	assertHasErrorMessage(t, "auto_increment must not null", id.Validate())

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

func TestColumnValidateUnsigned(t *testing.T) {
	id := Column{Name: "id", Type: TEXT, NotNull: true, AutoIncrement: true, Unsigned: true}
	assertHasErrorMessage(t, "only integer types can be unsigned", id.ValidateUnsigned())

	for _, integerType := range integerTypes {
		id.Type = integerType
		assertNotHasError(t, id.Validate())
	}

	for _, floatingType := range floatingTypes {
		id.Type = floatingType
		assertHasErrorMessage(t, "only integer types can be unsigned", id.Validate())
	}
}
