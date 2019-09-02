package dbs

import "testing"

func TestToColumnString(t *testing.T) {
	id := Column{Name: "id", Type: INT, Primary: true, NotNull: true, AutoIncrement: true}
	if id.ToString() != "id INT AUTO_INCREMENT PRIMARY KEY NOT NULL" {
		t.Fail()
	}

	name := Column{Name: "name", Type: TEXT, NotNull: true}
	if name.ToString() != "name TEXT NOT NULL" {
		t.Fail()
	}

	age := Column{}
	age.Name = "age"
	age.Type = "INT"
	if age.ToString() != "age INT" {
		t.Fail()
	}
}
