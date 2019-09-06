package dbs

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"testing"
)

func prepareDBSource() *DBSource {
	if os.Getenv("ENV") == "CI" {
		return prepareCIDBSource()
	}

	return prepareLocalDBSource()
}

func prepareCIDBSource() *DBSource {
	if os.Getenv("DRIVER") == MYSQL {
		return &DBSource{ServerName: "127.0.0.1", Name: "workspace", Driver: MYSQL, User: "root"}
	}

	return &DBSource{Name: "test.sqlite", Driver: SQLITE3}
}

func prepareLocalDBSource() *DBSource {
	return &DBSource{Name: "test.sqlite", Driver: SQLITE3}
}

// Use SQlite for testing schema install process
func TestSchemaInstall(t *testing.T) {
	dbSchema := &Schema{
		Name: "workspace",
		Platform: MYSQL,
		Tables: []Table{
			{
				Name: "user",
				PrimaryKey: []string{"id"},
				Columns: []Column{
					{Name: "id", Type: INT, NotNull: true, Unsigned: true, AutoIncrement:true},
					{Name: "name", Type: TEXT, NotNull: true},
					{Name: "age", Type: SMALLINT, NotNull: true, Unsigned: true},
				},
			},
		},
	}

	dbSource := prepareDBSource()

	db, err := dbSource.Connection()
	assertNotHasError(t, err)

	assertNotHasError(t, dbSchema.Install(db))

	_, err = db.Exec("INSERT INTO user (name, age) VALUES(\"Luan Phan\", 22)")
	assertNotHasError(t, err)

	var id, age int
	var name string
	err = db.QueryRow("select id, name, age from user").Scan(&id, &name, &age)
	assertNotHasError(t, err)
	assertStringEquals(t, "Luan Phan", name)
	assertIntEquals(t, 22, age)

	os.Remove("test.sqlite")
}
