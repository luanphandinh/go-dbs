package dbs

import (
	_ "github.com/go-sql-driver/mysql"
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
		Name: "workspace",
		Platform: platform,
		Tables: []Table{
			{
				Name: "user",
				PrimaryKey: []string{"id"},
				Columns: []Column{
					{Name: "id", Type: INT, NotNull: true, Unsigned: true, AutoIncrement:true},
					{Name: "name", Type: TEXT, NotNull: true},
					{Name: "age", Type: SMALLINT, NotNull: true, Unsigned: true, Length: 2},
				},
			},
			{
				Name: "company",
				PrimaryKey: []string{"id"},
				Columns: []Column{
					{Name: "id", Type: INT, NotNull: true, Unsigned: true, AutoIncrement:true},
					{Name: "name", Type: TEXT, NotNull: true, Length: 2},
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

	_, err = db.Exec("INSERT INTO user (name, age) VALUES(\"Luan Phan\", 22)")
	_, err = db.Exec("INSERT INTO company (name, rank) VALUES(\"Luan Phan Corps\", 1)")
	assertNotHasError(t, err)

	var id, age, rank int
	var name string
	err = db.QueryRow("select id, name, age from user").Scan(&id, &name, &age)
	assertNotHasError(t, err)
	assertStringEquals(t, "Luan Phan", name)
	assertIntEquals(t, 22, age)

	err = db.QueryRow("select name, rank from company").Scan(&name, &rank)
	assertNotHasError(t, err)
	assertStringEquals(t, "Luan Phan Corps", name)
	assertIntEquals(t, 1, rank)
}
