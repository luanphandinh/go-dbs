package dbs

import "testing"

func TestColumnDeclaration(t *testing.T) {
	mysqlPlatform := GetPlatform(MYSQL57)
	sqlitePlatform := GetPlatform(SQLITE3)
	postgresPlatfrom := GetPlatform(POSTGRES)

	id := Column{
		Name:          "id",
		Type:          INT,
		NotNull:       true,
		AutoIncrement: true,
	}
	assertStringEquals(t, "id INT NOT NULL AUTO_INCREMENT", mysqlPlatform.GetColumnDeclarationSQL(&id))
	assertStringEquals(t, "id INTEGER NOT NULL", sqlitePlatform.GetColumnDeclarationSQL(&id))
	assertStringEquals(t, "id INT NOT NULL", postgresPlatfrom.GetColumnDeclarationSQL(&id))

	name := Column{Name: "name", Type: TEXT, NotNull: true}
	assertStringEquals(t, "name TEXT NOT NULL", mysqlPlatform.GetColumnDeclarationSQL(&name))
	assertStringEquals(t, "name TEXT NOT NULL", sqlitePlatform.GetColumnDeclarationSQL(&name))
	assertStringEquals(t, "name TEXT NOT NULL", postgresPlatfrom.GetColumnDeclarationSQL(&name))

	age := Column{}
	age.Name = "age"
	age.Type = "INT"
	age.Unsigned = true
	assertStringEquals(t, "age INT UNSIGNED", mysqlPlatform.GetColumnDeclarationSQL(&age))
	assertStringEquals(t, "age INTEGER", sqlitePlatform.GetColumnDeclarationSQL(&age))
	assertStringEquals(t, "age INT", postgresPlatfrom.GetColumnDeclarationSQL(&age))

	age.Length = 2
	assertStringEquals(t, "age INT(2) UNSIGNED", mysqlPlatform.GetColumnDeclarationSQL(&age))
	assertStringEquals(t, "age INTEGER(2)", sqlitePlatform.GetColumnDeclarationSQL(&age))
	assertStringEquals(t, "age INT", postgresPlatfrom.GetColumnDeclarationSQL(&age))

	age.Default = "10"
	assertStringEquals(t, "age INT(2) UNSIGNED DEFAULT 10", mysqlPlatform.GetColumnDeclarationSQL(&age))
	assertStringEquals(t, "age INTEGER(2) DEFAULT 10", sqlitePlatform.GetColumnDeclarationSQL(&age))
	assertStringEquals(t, "age INT DEFAULT 10", postgresPlatfrom.GetColumnDeclarationSQL(&age))

	age.Check = "age < 150"
	assertStringEquals(t, "age INT(2) UNSIGNED DEFAULT 10 CHECK (age < 150)", mysqlPlatform.GetColumnDeclarationSQL(&age))
	assertStringEquals(t, "age INTEGER(2) CHECK (age < 150) DEFAULT 10", sqlitePlatform.GetColumnDeclarationSQL(&age))
	assertStringEquals(t, "age INT DEFAULT 10 CHECK (age < 150)", postgresPlatfrom.GetColumnDeclarationSQL(&age))
}
