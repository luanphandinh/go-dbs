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
	}

	table := Table{
		Name: "user",
		Columns: []Column{
			id,
			name,
			age,
		},
	}
	assertStringEquals(t, "CREATE TABLE IF NOT EXISTS user (id INT NOT NULL AUTO_INCREMENT, name TEXT NOT NULL, age INT)", mysqlPlatform.GetTableSQLCreate(&table))
	assertStringEquals(t, "CREATE TABLE IF NOT EXISTS user (id INT NOT NULL AUTOINCREMENT, name TEXT NOT NULL, age INT)", sqlitePlatform.GetTableSQLCreate(&table))
}
