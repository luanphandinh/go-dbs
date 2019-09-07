package dbs

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"testing"
)

func getTestEnv() string {
	if env := os.Getenv("ENV"); env != "" {
		return env
	}

	return ""
}

func getTestPlatform() string {
	if driver := os.Getenv("DRIVER"); driver != "" {
		return driver
	}

	return SQLITE3
}

func prepareDBSource(env string, platform string) *DBSource {
	if env == "CI" {
		return prepareCIDBSource(platform)
	}

	return prepareLocalDBSource(platform)
}

func prepareCIDBSource(platform string) *DBSource {
	if platform == MYSQL {
		return &DBSource{ServerName: "127.0.0.1", Name: "workspace", Driver: MYSQL, User: "root"}
	}

	return &DBSource{Name: "test.sqlite", Driver: SQLITE3}
}

func prepareLocalDBSource(platform string) *DBSource {
	if platform == MYSQL {
		return nil
	}

	return &DBSource{Name: "test.sqlite", Driver: SQLITE3}
}

// Use SQlite for testing schema install process
func TestSchemaInstall(t *testing.T) {
	env := getTestEnv()
	platform := getTestPlatform()
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
					{Name: "age", Type: SMALLINT, NotNull: true, Unsigned: true},
				},
			},
			{
				Name: "company",
				PrimaryKey: []string{"id"},
				Columns: []Column{
					{Name: "id", Type: INT, NotNull: true, Unsigned: true, AutoIncrement:true},
					{Name: "name", Type: TEXT, NotNull: true},
					{Name: "rank", Type: SMALLINT, NotNull: true, Unsigned: true, Unique: true},
				},
			},
		},
	}

	dbSource := prepareDBSource(env, platform)

	db, err := dbSource.Connection()
	assertNotHasError(t, err)

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
	os.Remove("test.sqlite")
}
