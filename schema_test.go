package dbs

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"testing"
)

func getEnvPlatform() string {
	if driver := os.Getenv("PLATFORM"); driver != "" {
		return driver
	}

	return SQLITE3
}

func prepareDBSource(platform string) *DBSource {
	serverName := os.Getenv("SERVER_NAME")
	dbName := os.Getenv("DB_NAME")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")

	return &DBSource{ServerName: serverName, Name: dbName, Driver: platform, User: user, Password: password}
}

func TestSchemaInstall(t *testing.T) {
	platform := getEnvPlatform()
	dbSchema := &Schema{
		Name: "company",
		Platform: platform,
		Tables: []Table{
			{
				Name: "employee",
				PrimaryKey: []string{"id"},
				Columns: []Column{
					{Name: "id", Type: INT, NotNull: true, Unsigned: true, AutoIncrement:true},
					{Name: "name", Type: TEXT, NotNull: true},
					{Name: "department_id", Type: INT},
					{Name: "valid", Type: SMALLINT, Default: "1"},
					{Name: "age", Type: SMALLINT, NotNull: true, Unsigned: true, Length: 2},
				},
			},
			{
				Name: "department",
				PrimaryKey: []string{"id"},
				Columns: []Column{
					{Name: "id", Type: INT, NotNull: true, Unsigned: true, AutoIncrement:true},
					{Name: "name", Type: TEXT, NotNull: true, Length: 2},
					{Name: "revenue", Type: FLOAT, NotNull: true, Default: "1.01"},
					{Name: "rank", Type: SMALLINT, NotNull: true, Unsigned: true, Unique: true, Length: 1},
				},
			},
		},
	}

	dbSource := prepareDBSource(platform)

	db, err := dbSource.Connection()
	assertNotHasError(t, err)
	assertNotHasError(t, dbSchema.Drop(db))
	assertNotHasError(t, dbSchema.Install(db))

	// Temporary disabled for postgres
	// @TODO query builder will help to create query across platforms
	if platform != POSTGRES {
		_, err = db.Exec("INSERT INTO department (name, rank) VALUES ('Luan Phan Corps', 1)")
		_, err = db.Exec("INSERT INTO employee (name, age) VALUES ('Luan Phan', 22)")
		assertNotHasError(t, err)

		var valid, age, rank int
		var name string
		var revenue float32
		err = db.QueryRow("select valid, name, age from employee").Scan(&valid, &name, &age)
		assertNotHasError(t, err)
		assertStringEquals(t, "Luan Phan", name)
		assertIntEquals(t, 22, age)
		assertIntEquals(t, 1, valid)

		err = db.QueryRow("select name, rank, revenue from department").Scan(&name, &rank, &revenue)
		assertNotHasError(t, err)
		assertStringEquals(t, "Luan Phan Corps", name)
		assertIntEquals(t, 1, rank)
		assertFloatEquals(t, 1.01, revenue)
	}
}
