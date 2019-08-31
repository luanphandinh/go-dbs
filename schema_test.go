package go_db

import "testing"

func TestColumnValidate(t *testing.T) {
	id := Column{}
	if err := id.validate(); err != nil {
		if err.Error() != "column name should not empty" {
			t.Fail()
		}
	}

	id.Name = "id"
	if err := id.validate(); err != nil {
		if err.Error() != "column type should not empty" {
			t.Fail()
		}
	}

	id.Type = "NVARCHAR"
	if err := id.validate(); err != nil {
		if err.Error() != "NVARCHAR can not auto_increment" {
			t.Fail()
		}
	}
}

func TestToColumnString(t *testing.T) {
	id := Column{"id", "INT", true, true, true}
	if id.toString() != "id INT AUTO_INCREMENT PRIMARY KEY NOT NULL" {
		t.Fail()
	}

	name := Column{"name", "NVARCHAR(50)", true, false, false}
	if name.toString() != "name NVARCHAR(50) NOT NULL" {
		t.Fail()
	}

	age := Column{}
	age.Name = "age"
	age.Type = "INT"
	if age.toString() != "age INT" {
		t.Fail()
	}
}
