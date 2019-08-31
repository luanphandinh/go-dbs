package dbs

import "testing"

func TestColumnValidate(t *testing.T) {
	id := Column{}
	if err := id.Validate(); err != nil {
		if err.Error() != "column name should not empty" {
			t.Fail()
		}
	}

	id.Name = "id"
	if err := id.Validate(); err != nil {
		if err.Error() != "column type should not empty" {
			t.Fail()
		}
	}

	id.Type = "NVARCHAR"
	if err := id.Validate(); err != nil {
		if err.Error() != "NVARCHAR can not auto_increment" {
			t.Fail()
		}
	}
}

func TestToColumnString(t *testing.T) {
	id := Column{"id", "INT", true, true, true}
	if id.ToString() != "id INT AUTO_INCREMENT PRIMARY KEY NOT NULL" {
		t.Fail()
	}

	name := Column{"name", "NVARCHAR(50)", true, false, false}
	if name.ToString() != "name NVARCHAR(50) NOT NULL" {
		t.Fail()
	}

	age := Column{}
	age.Name = "age"
	age.Type = "INT"
	if age.ToString() != "age INT" {
		t.Fail()
	}
}

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

func TestToTableString(t *testing.T) {
	id := Column{"id", "INT", true, true, true}
	name := Column{Name: "name", Type: "NVARCHAR(50)", NotNull:true}
	age := Column{}
	age.Name = "age"
	age.Type = "INT"

	cols := []Column{id, name, age}
	table := Table{"user", cols}
	if table.ToString() != "CREATE TABLE IF NOT EXISTS user (id INT AUTO_INCREMENT PRIMARY KEY NOT NULL, name NVARCHAR(50) NOT NULL, age INT)" {
		t.Fail()
	}
}
