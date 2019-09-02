package dbs

import "testing"

func TestColumnTypeValidate(t *testing.T)  {
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
}

func TestColumnValidate(t *testing.T) {
	id := Column{}
	err := id.Validate()
	if err == nil {
		t.Fail()
	} else {
		if err.Error() != "column name should not empty" {
			t.Fail()
		}
	}

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

	id.Type = TEXT
	id.AutoIncrement = true
	if err := id.Validate(); err == nil {
		t.Fail()
	} else {
		if err.Error() != "TEXT can not auto_increment" {
			t.Fail()
		}
	}
}
