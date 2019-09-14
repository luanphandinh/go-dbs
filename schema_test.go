package dbs

import (
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
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
	user       = os.Getenv("DB_USER")
	password   = os.Getenv("DB_PASSWORD")
)

func getSchema(platform string) *Schema {
	// return &Schema{
	// 	Name:     "company",
	// 	Platform: platform,
	// 	Comment:  "The Company Schema",
	// 	Tables: []Table{
	// 		{
	// 			Name:       "department",
	// 			PrimaryKey: []string{"id"},
	// 			Columns: []Column{
	// 				{Name: "id", Type: INT, NotNull: true, Unsigned: true, AutoIncrement: true},
	// 				{Name: "name", Type: TEXT, NotNull: true, Length: 2},
	// 				{Name: "revenue", Type: FLOAT, NotNull: true, Default: "1.01"},
	// 				{Name: "position", Type: SMALLINT, NotNull: true, Unsigned: true, Unique: true, Length: 1},
	// 			},
	// 			Comment: "Departments of company",
	// 		},
	// 		{
	// 			Name:       "employee",
	// 			PrimaryKey: []string{"id"},
	// 			Columns: []Column{
	// 				{Name: "id", Type: INT, NotNull: true, Unsigned: true, AutoIncrement: true},
	// 				{Name: "name", Type: TEXT, NotNull: true},
	// 				{Name: "department_id", Type: INT, Unsigned: true},
	// 				{Name: "valid", Type: SMALLINT, Default: "1", Comment: "Indicate employee status"},
	// 				{Name: "age", Type: SMALLINT, NotNull: true, Unsigned: true, Length: 2, Check: "age > 20"},
	// 			},
	// 			Checks: []string{"age < 50"},
	// 			ForeignKeys: []ForeignKey{
	// 				{Referer: "department_id", Reference: "department(id)"},
	// 			},
	// 		},
	// 	},
	// }
	schema := new(Schema).WithName("company").OnPlatform(platform).WithComment("The Company Schema")

	department := new(Table).WithName("department").WithComment("Departments of company")
	department.AddColumn(*new(Column).WithName("id").WithType(INT).IsNotNull().IsUnsigned().IsAutoIncrement())
	department.AddColumn(*new(Column).WithName("name").WithType(TEXT).WithLength(2).IsNotNull())
	department.AddColumn(*new(Column).WithName("revenue").WithType(FLOAT).IsNotNull().IsUnsigned().WithDefault("1.01"))
	department.AddColumn(*new(Column).WithName("position").WithType(SMALLINT).WithLength(1).IsNotNull().IsUnsigned().IsUnique())
	department.AddPrimaryKey([]string{"id"})

	employee := new(Table).WithName("employee")
	employee.AddColumn(*new(Column).WithName("id").WithType(INT).IsNotNull().IsUnsigned().IsAutoIncrement())
	employee.AddColumn(*new(Column).WithName("name").WithType(TEXT).IsNotNull())
	employee.AddColumn(*new(Column).WithName("department_id").WithType(INT).IsUnsigned())
	employee.AddColumn(*new(Column).WithName("valid").WithType(SMALLINT).WithDefault("1").WithComment("Indicate employee status"))
	employee.AddColumn(*new(Column).WithName("age").WithType(SMALLINT).IsNotNull().IsUnsigned().WithLength(2).AddCheck("age > 20"))

	employee.AddPrimaryKey([]string{"id"})
	employee.AddCheck("age < 50")
	employee.AddForeignKey("department_id", "department(id)")

	schema.AddTable(*department)
	schema.AddTable(*employee)

	return schema
}

func setupDB(t *testing.T, dbPlatform Platform, dbSchema *Schema) (*sql.DB, error) {
	db, err := sql.Open(
		dbPlatform.GetDriverName(),
		dbPlatform.GetDBConnectionString(serverName, 3306, user, password, dbName),
	)
	assertNotHasError(t, err)
	assertNotHasError(t, dbSchema.Drop(db))
	assertNotHasError(t, dbSchema.Install(db))

	return db, err
}

func TestSchemaInstall(t *testing.T) {
	dbSchema := getSchema(platform)
	dbPlatform := GetPlatform(platform)

	db, err := setupDB(t, dbPlatform, dbSchema)

	employee := dbPlatform.GetSchemaAccessName(dbSchema.Name, "employee")
	department := dbPlatform.GetSchemaAccessName(dbSchema.Name, "department")

	_, err = db.Exec(fmt.Sprintf("INSERT INTO %s (name, position) VALUES ('Luan Phan Corps', 1)", department))
	assertNotHasError(t, err)
	// Checks constraint is parsed but will be ignored in mysql5.7
	// @TODO query builder will help to create query across platforms
	if platform != MYSQL57 {
		_, err = db.Exec(fmt.Sprintf("INSERT INTO %s (name, age, department_id) VALUES ('Luan Phan', 5, 1)", employee))
		assertHasError(t, err)

		_, err = db.Exec(fmt.Sprintf("INSERT INTO %s (name, age, department_id) VALUES ('Luan Phan', 51, 1)", employee))
		assertHasError(t, err)

		// Some check is different across platforms eg: length(name) and LEN(name) in mssql
		// _, err = db.Exec(fmt.Sprintf("INSERT INTO %s (name, age) VALUES ('Luan Phan Wrong', 22)", employee))
		// assertHasError(t, err)

		_, err = db.Exec(fmt.Sprintf("INSERT INTO %s (name, age) VALUES (NULL, 22)", employee))
		assertHasError(t, err)

		_, err = db.Exec(fmt.Sprintf("INSERT INTO %s (name, age) VALUES ('Luan Phan', 22)", employee))
		assertNotHasError(t, err)
	}

	_, err = db.Exec(fmt.Sprintf("INSERT INTO %s (name, age, department_id) VALUES ('Luan Phan', 22, 1)", employee))
	assertNotHasError(t, err)

	var valid, age, position int
	var name string
	var revenue float32
	err = db.QueryRow(fmt.Sprintf("select valid, name, age from %s", employee)).Scan(&valid, &name, &age)
	assertNotHasError(t, err)
	assertStringEquals(t, "Luan Phan", name)
	assertIntEquals(t, 22, age)
	assertIntEquals(t, 1, valid)

	err = db.QueryRow(fmt.Sprintf("select name, position, revenue from %s", department)).Scan(&name, &position, &revenue)
	assertNotHasError(t, err)
	assertStringEquals(t, "Luan Phan Corps", name)
	assertIntEquals(t, 1, position)
	assertFloatEquals(t, 1.01, revenue)
}

func TestAutoIncrement(t *testing.T) {
	dbSchema := getSchema(platform)
	dbPlatform := GetPlatform(platform)

	db, err := setupDB(t, dbPlatform, dbSchema)

	employee := dbPlatform.GetSchemaAccessName(dbSchema.Name, "employee")
	department := dbPlatform.GetSchemaAccessName(dbSchema.Name, "department")

	_, err = db.Exec(fmt.Sprintf("INSERT INTO %s (name, position) VALUES ('Luan Phan Corps', 1)", department))
	assertNotHasError(t, err)

	_, err = db.Exec(fmt.Sprintf("INSERT INTO %s (name, age, department_id) VALUES ('Luan Phan', 22, 1)", employee))
	assertNotHasError(t, err)

	var valid, age, id int
	var name string
	err = db.QueryRow(fmt.Sprintf("select id, valid, name, age from %s", employee)).Scan(&id, &valid, &name, &age)
	assertIntEquals(t, 1, id)
	assertNotHasError(t, err)

	_, err = db.Exec(fmt.Sprintf("INSERT INTO %s (name, age, department_id) VALUES ('Luan Phan', 22, 1)", employee))
	assertNotHasError(t, err)
	err = db.QueryRow(fmt.Sprintf("select id, valid, name, age from %s where id = 2", employee)).Scan(&id, &valid, &name, &age)
	assertIntEquals(t, 2, id)
	assertNotHasError(t, err)
}
