package dbs

import "testing"


func TestToTableString(t *testing.T) {
	id := Column{Name: "id", Type: INT, Primary: true, NotNull: true, AutoIncrement: true}
	name := Column{Name: "name", Type: TEXT, NotNull:true}
	age := Column{}
	age.Name = "age"
	age.Type = INT

	cols := []Column{id, name, age}
	table := Table{"user", cols}
	if table.ToString() != "CREATE TABLE IF NOT EXISTS user (id INT AUTO_INCREMENT PRIMARY KEY NOT NULL, name TEXT NOT NULL, age INT)" {
		t.Fail()
	}
}
