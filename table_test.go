package dbs

import "testing"

func TestToTableString(t *testing.T) {
	id := Column{
		Name:          "id",
		Type:          INT,
		Primary:       true,
		NotNull:       true,
		AutoIncrement: true,
	}

	name := Column{
		Name:    "name",
		Type:    TEXT,
		NotNull: true,
	}

	age := Column{
		Name: "age",
		Type: INT,
	}

	table := Table{
		Name: "user",
		Columns: []Column{
			id,
			name,
			age,
		},
	}
	assertStringEquals(t, "CREATE TABLE IF NOT EXISTS user (id INT NOT NULL PRIMARY KEY AUTO_INCREMENT, name TEXT NOT NULL, age INT)", table.ToString())
}
