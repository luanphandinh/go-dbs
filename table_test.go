package dbs

import "testing"

func TestToTableDeclaration(t *testing.T) {
	mysqlPlatform := GetPlatform(MYSQL57)
	sqlitePlatform := GetPlatform(SQLITE3)
	postgresPlatform := GetPlatform(POSTGRES)

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
		Default: "10",
		Check: "age < 1000",
		Comment: "age should less than 1000",
	}

	table := Table{
		Name: "user",
		PrimaryKey: []string{"id"},
		Columns: []Column{
			id,
			name,
			age,
		},
		Check: []string{"age > 50"},
	}
	assertStringEquals(t, "CREATE TABLE IF NOT EXISTS user (id INT NOT NULL AUTO_INCREMENT, name TEXT NOT NULL, age INT(4) DEFAULT 10 CHECK (age < 1000) COMMENT 'age should less than 1000', PRIMARY KEY (id), CHECK (age > 50))",mysqlPlatform.GetTableCreateSQL("", &table))
	assertStringEquals(t, "CREATE TABLE IF NOT EXISTS user (id INTEGER NOT NULL, name TEXT NOT NULL, age INTEGER(4) CHECK (age < 1000) DEFAULT 10, PRIMARY KEY (id), CHECK (age > 50))", sqlitePlatform.GetTableCreateSQL("", &table))
	assertStringEquals(t, `CREATE TABLE IF NOT EXISTS public.user (id INT NOT NULL, name TEXT NOT NULL, age INT DEFAULT 10 CHECK (age < 1000), PRIMARY KEY (id), CHECK (age > 50));
 CREATE SEQUENCE IF NOT EXISTS public.user_id_seq; ALTER TABLE public.user ALTER id SET DEFAULT NEXTVAL('public.user_id_seq');
 COMMENT ON COLUMN public.user.age IS 'age should less than 1000'`, postgresPlatform.GetTableCreateSQL("public", &table))

	table.PrimaryKey = []string{"id", "name"}
	assertStringEquals(t, "PRIMARY KEY (id, name)", mysqlPlatform.GetPrimaryDeclaration(table.PrimaryKey))
	assertStringEquals(t, "PRIMARY KEY (id, name)", sqlitePlatform.GetPrimaryDeclaration(table.PrimaryKey))
	assertStringEquals(t, "PRIMARY KEY (id, name)", postgresPlatform.GetPrimaryDeclaration(table.PrimaryKey))

	table.Check = []string{"age > 50", "length(name) < 100"}
	assertStringEquals(t, "CHECK (age > 50), CHECK (length(name) < 100)", mysqlPlatform.GetTableCheckDeclaration(table.Check))
	assertStringEquals(t, "CHECK (age > 50), CHECK (length(name) < 100)", sqlitePlatform.GetTableCheckDeclaration(table.Check))
	assertStringEquals(t, "CHECK (age > 50), CHECK (length(name) < 100)", postgresPlatform.GetTableCheckDeclaration(table.Check))
}
