package dbs

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"testing"
)

var (
	platform   = os.Getenv("PLATFORM")
	serverName = os.Getenv("SERVER_NAME")
	dbName     = os.Getenv("DB_NAME")
	user       = os.Getenv("USER")
	password   = os.Getenv("PASSWORD")
)

func TestSchemaInstall(t *testing.T) {
	dbSchema := &Schema{
		Name:     "company",
		Platform: platform,
		Tables: []Table{
			{
				Name:       "employee",
				PrimaryKey: []string{"id"},
				Columns: []Column{
					{Name: "id", Type: INT, NotNull: true, Unsigned: true, AutoIncrement: true},
					{Name: "name", Type: TEXT, NotNull: true},
					{Name: "department_id", Type: INT},
					{Name: "valid", Type: SMALLINT, Default: "1"},
					{Name: "age", Type: SMALLINT, NotNull: true, Unsigned: true, Length: 2, Check: "age > 20"},
				},
				Check: []string{"age < 50", "length(name) < 10"},
			},
			{
				Name:       "department",
				PrimaryKey: []string{"id"},
				Columns: []Column{
					{Name: "id", Type: INT, NotNull: true, Unsigned: true, AutoIncrement: true},
					{Name: "name", Type: TEXT, NotNull: true, Length: 2},
					{Name: "revenue", Type: FLOAT, NotNull: true, Default: "1.01"},
					{Name: "rank", Type: SMALLINT, NotNull: true, Unsigned: true, Unique: true, Length: 1},
				},
			},
		},
	}

	dbPlatform := GetPlatform(platform)
	db, err := sql.Open(
		platform,
		dbPlatform.GetDBConnectionString(serverName, 3306, user, password, dbName),
	)
	assertNotHasError(t, err)
	assertNotHasError(t, dbSchema.Drop(db))
	assertNotHasError(t, dbSchema.Install(db))

	employee := dbPlatform.GetTableName(dbSchema.Name, "employee")
	department := dbPlatform.GetTableName(dbSchema.Name, "department")

	// Check constraint is parsed but will be ignored in mysql5.7
	// @TODO query builder will help to create query across platforms
	if platform != MYSQL {
		_, err = db.Exec(fmt.Sprintf("INSERT INTO %s (id, name, age) VALUES (1, 'Luan Phan', 5)", employee))
		assertHasError(t, err)

		_, err = db.Exec(fmt.Sprintf("INSERT INTO %s (id, name, age) VALUES (1, 'Luan Phan', 51)", employee))
		assertHasError(t, err)

		_, err = db.Exec(fmt.Sprintf("INSERT INTO %s (id, name, age) VALUES (1, 'Luan Phan Wrong', 22)", employee))
		assertHasError(t, err)
	}

	_, err = db.Exec(fmt.Sprintf("INSERT INTO %s (id, name, rank) VALUES (1, 'Luan Phan Corps', 1)", department))
	_, err = db.Exec(fmt.Sprintf("INSERT INTO %s (id, name, age) VALUES (1, 'Luan Phan', 22)", employee))

	assertNotHasError(t, err)

	var valid, age, rank int
	var name string
	var revenue float32
	err = db.QueryRow(fmt.Sprintf("select valid, name, age from %s", employee)).Scan(&valid, &name, &age)
	assertNotHasError(t, err)
	assertStringEquals(t, "Luan Phan", name)
	assertIntEquals(t, 22, age)
	assertIntEquals(t, 1, valid)

	err = db.QueryRow(fmt.Sprintf("select name, rank, revenue from %s", department)).Scan(&name, &rank, &revenue)
	assertNotHasError(t, err)
	assertStringEquals(t, "Luan Phan Corps", name)
	assertIntEquals(t, 1, rank)
	assertFloatEquals(t, 1.01, revenue)
}
