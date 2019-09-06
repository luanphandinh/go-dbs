package dbs

import "testing"

func TestColumnDeclaration(t *testing.T) {
	id := Column{
		Name:          "id",
		Type:          INT,
		Primary:       true,
		NotNull:       true,
		AutoIncrement: true,
	}
	assertStringEquals(t, "id INT NOT NULL PRIMARY KEY AUTO_INCREMENT", id.GetSQLDeclaration(&MySqlPlatform{}))
	assertStringEquals(t, "id INT NOT NULL PRIMARY KEY AUTOINCREMENT", id.GetSQLDeclaration(&SqlitePlatform{}))

	name := Column{Name: "name", Type: TEXT, NotNull: true}
	assertStringEquals(t, "name TEXT NOT NULL", name.GetSQLDeclaration(&MySqlPlatform{}))
	assertStringEquals(t, "name TEXT NOT NULL", name.GetSQLDeclaration(&SqlitePlatform{}))

	age := Column{}
	age.Name = "age"
	age.Type = "INT"
	age.Unsigned = true
	assertStringEquals(t, "age INT UNSIGNED", age.GetSQLDeclaration(&MySqlPlatform{}))
	assertStringEquals(t, "age INT UNSIGNED", age.GetSQLDeclaration(&SqlitePlatform{}))
}
