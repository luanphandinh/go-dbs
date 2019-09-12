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
	assertStringEquals(t, "id INT NOT NULL AUTO_INCREMENT", mysqlPlatform.BuildColumnDeclarationSQL(&id))
	assertStringEquals(t, "id INTEGER NOT NULL", sqlitePlatform.BuildColumnDeclarationSQL(&id))
	assertStringEquals(t, "id INT NOT NULL", postgresPlatfrom.BuildColumnDeclarationSQL(&id))

	name := Column{Name: "name", Type: TEXT, NotNull: true}
	assertStringEquals(t, "name TEXT NOT NULL", mysqlPlatform.BuildColumnDeclarationSQL(&name))
	assertStringEquals(t, "name TEXT NOT NULL", sqlitePlatform.BuildColumnDeclarationSQL(&name))
	assertStringEquals(t, "name TEXT NOT NULL", postgresPlatfrom.BuildColumnDeclarationSQL(&name))

	age := Column{}
	age.Name = "age"
	age.Type = "INT"
	age.Unsigned = true
	assertStringEquals(t, "age INT UNSIGNED", mysqlPlatform.BuildColumnDeclarationSQL(&age))
	assertStringEquals(t, "age INTEGER", sqlitePlatform.BuildColumnDeclarationSQL(&age))
	assertStringEquals(t, "age INT", postgresPlatfrom.BuildColumnDeclarationSQL(&age))

	age.Length = 2
	assertStringEquals(t, "age INT(2) UNSIGNED", mysqlPlatform.BuildColumnDeclarationSQL(&age))
	assertStringEquals(t, "age INTEGER(2)", sqlitePlatform.BuildColumnDeclarationSQL(&age))
	assertStringEquals(t, "age INT", postgresPlatfrom.BuildColumnDeclarationSQL(&age))

	age.Default = "10"
	assertStringEquals(t, "age INT(2) UNSIGNED DEFAULT 10", mysqlPlatform.BuildColumnDeclarationSQL(&age))
	assertStringEquals(t, "age INTEGER(2) DEFAULT 10", sqlitePlatform.BuildColumnDeclarationSQL(&age))
	assertStringEquals(t, "age INT DEFAULT 10", postgresPlatfrom.BuildColumnDeclarationSQL(&age))

	age.Check = "age < 150"
	assertStringEquals(t, "age INT(2) UNSIGNED DEFAULT 10 CHECK (age < 150)", mysqlPlatform.BuildColumnDeclarationSQL(&age))
	assertStringEquals(t, "age INTEGER(2) CHECK (age < 150) DEFAULT 10", sqlitePlatform.BuildColumnDeclarationSQL(&age))
	assertStringEquals(t, "age INT DEFAULT 10 CHECK (age < 150)", postgresPlatfrom.BuildColumnDeclarationSQL(&age))

	age.Comment = "age should be less than 150"
	assertStringEquals(t, "age INT(2) UNSIGNED DEFAULT 10 CHECK (age < 150) COMMENT 'age should be less than 150'", mysqlPlatform.BuildColumnDeclarationSQL(&age))
	assertStringEquals(t, "age INTEGER(2) CHECK (age < 150) DEFAULT 10", sqlitePlatform.BuildColumnDeclarationSQL(&age))
	assertStringEquals(t, "age INT DEFAULT 10 CHECK (age < 150)", postgresPlatfrom.BuildColumnDeclarationSQL(&age))
}
