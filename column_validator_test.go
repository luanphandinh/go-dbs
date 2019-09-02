package dbs

import "testing"

func TestColumnNameValidate(t *testing.T) {
	id := Column{}
	err := id.Validate()
	if err == nil {
		t.Fail()
	} else {
		if err.Error() != "column name should not empty" {
			t.Fail()
		}
	}
}

func TestColumnTypeValidate(t *testing.T) {
	for _, colType := range allTypes {
		test := Column{Type: colType}
		if err := test.ValidateType(); err != nil {
			t.Fail()
		}
	}

	test := Column{Type: "WRONG"}
	if err := test.ValidateType(); err == nil {
		t.Fail()
	}

	id := Column{}
	id.Name = "id"
	if err := id.Validate(); err == nil {
		t.Fail()
	} else {
		if err.Error() != "column type should not empty" {
			t.Fail()
		}
	}

	id.Type = "SOMETHING"
	if err := id.Validate(); err == nil {
		t.Fail()
	} else {
		if err.Error() != "incorrect type name" {
			t.Fail()
		}
	}
}

func TestColumnValidateAutoIncrement(t *testing.T) {
	id := Column{}
	id.Name = "id"
	id.Type = TEXT
	id.AutoIncrement = true
	if err := id.Validate(); err == nil {
		t.Fail()
	} else {
		if err.Error() != "TEXT can not auto_increment" {
			t.Fail()
		}
	}

	id.Type = INT
	if err := id.Validate(); err == nil {
		t.Fail()
	} else {
		if err.Error() != "auto_increment must not null" {
			t.Fail()
		}
	}

	id.NotNull = true
	for _, integerType := range integerTypes {
		id.Type = integerType
		if err := id.Validate(); err != nil {
			t.Fail()
		}
	}

	for _, floatingType := range floatingTypes {
		id.Type = floatingType
		if err := id.Validate(); err != nil {
			t.Fail()
		}
	}
}
