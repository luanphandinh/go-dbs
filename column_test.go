package dbs

import "testing"

func TestColumnDeclaration(t *testing.T) {
	mysqlPlatform := GetPlatform(MYSQL)
	sqlitePlatform := GetPlatform(SQLITE3)

	id := Column{
		Name:          "id",
		Type:          INT,
		NotNull:       true,
		AutoIncrement: true,
	}
	assertStringEquals(t, "id INT NOT NULL AUTO_INCREMENT", mysqlPlatform.GetColumnDeclarationSQL(&id))
	assertStringEquals(t, "id INTEGER", sqlitePlatform.GetColumnDeclarationSQL(&id))

	name := Column{Name: "name", Type: TEXT, NotNull: true}
	assertStringEquals(t, "name TEXT NOT NULL", mysqlPlatform.GetColumnDeclarationSQL(&name))
	assertStringEquals(t, "name TEXT", sqlitePlatform.GetColumnDeclarationSQL(&name))

	age := Column{}
	age.Name = "age"
	age.Type = "INT"
	age.Unsigned = true
	assertStringEquals(t, "age INT UNSIGNED", mysqlPlatform.GetColumnDeclarationSQL(&age))
	assertStringEquals(t, "age INTEGER", sqlitePlatform.GetColumnDeclarationSQL(&age))
}
