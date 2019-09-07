package dbs

import "testing"

func TestToTableDeclaration(t *testing.T) {
	mysqlPlatform := GetPlatform(MYSQL)
	sqlitePlatform := GetPlatform(SQLITE3)

	id := Column{
		Name:          "id",
		Type:          INT,
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
		Length: 4,
	}

	table := Table{
		Name: "user",
		PrimaryKey: []string{"id"},
		Columns: []Column{
			id,
			name,
			age,
		},
	}
	assertStringEquals(t, "CREATE TABLE IF NOT EXISTS user (id INT NOT NULL AUTO_INCREMENT, name TEXT NOT NULL, age INT(4), PRIMARY KEY (id))", mysqlPlatform.GetTableCreateSQL(&table))
	assertStringEquals(t, "PRIMARY KEY (id)", mysqlPlatform.GetPrimaryKeyCreateSQL(&table))

	assertStringEquals(t, "CREATE TABLE IF NOT EXISTS user (id INTEGER, name TEXT, age INTEGER(4), PRIMARY KEY (id))", sqlitePlatform.GetTableCreateSQL(&table))
}
