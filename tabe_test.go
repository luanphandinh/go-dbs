package dbs

import "testing"


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
