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
	assertStringEquals(t, "id INT NOT NULL AUTO_INCREMENT", mysqlPlatform.GetColumnSQLDeclaration(&id))
	assertStringEquals(t, "id INT NOT NULL AUTOINCREMENT", sqlitePlatform.GetColumnSQLDeclaration(&id))

	name := Column{Name: "name", Type: TEXT, NotNull: true}
	assertStringEquals(t, "name TEXT NOT NULL", mysqlPlatform.GetColumnSQLDeclaration(&name))
	assertStringEquals(t, "name TEXT NOT NULL", sqlitePlatform.GetColumnSQLDeclaration(&name))

	age := Column{}
	age.Name = "age"
	age.Type = "INT"
	age.Unsigned = true
	assertStringEquals(t, "age INT UNSIGNED", mysqlPlatform.GetColumnSQLDeclaration(&age))
	assertStringEquals(t, "age INT UNSIGNED", sqlitePlatform.GetColumnSQLDeclaration(&age))
}
