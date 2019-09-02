package dbs

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"testing"
)

// Use SQlite for testing schema install process
func TestSchemaInstall(t *testing.T) {
	dbSchema := &Schema{
		Name: "workspace",
		Tables: []Table{
			{
				"user",
				[]Column{
					{Name: "id", Type: INT, Primary: true, NotNull: true, AutoIncrement: true},
					{Name: "name", Type: TEXT, NotNull: true},
				},
			},
		},
	}

	db, err := sql.Open("sqlite3", "test.sqlite")
	if err != nil {
		fmt.Println(err.Error())
		t.Fail()
	}

	if err := dbSchema.Install(db); err != nil {
		fmt.Println(err.Error())
		t.Fail()
	}

	_, err = db.Exec("INSERT INTO user (id, name) VALUES(1, \"Luan Phan\")")
	if err != nil {
		fmt.Println(err.Error())
		t.Fail()
	}

	var id int
	var name string
	err = db.QueryRow("select id, name from user").Scan(&id, &name)
	if name != "Luan Phan" {
		t.Fail()
	}

	os.Remove("test.sqlite")
}
