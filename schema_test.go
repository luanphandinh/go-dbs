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

func getSchema() *Schema {
	SetPlatform(platform)
	// return &Schema{
	// 	Name:     "company",
	// 	dbPlatform: platform,
	// 	Comment:  "The Company Schema",
	// 	Tables: []*Table{
	// 		{
	// 			Name:       "department",
	// 			PrimaryKey: []string{"id"},
	// 			Columns: []*Column{
	// 				{Name: "id", Type: INT, NotNull: true, Unsigned: true, AutoIncrement: true},
	// 				{Name: "name", Type: NVARCHAR, NotNull: true, Length: 20},
	// 				{Name: "revenue", Type: FLOAT, NotNull: true, Default: "1.01"},
	// 				{Name: "position", Type: SMALLINT, NotNull: true, Unsigned: true, Unique: true},
	// 			},
	// 			Comment: "Departments of company",
	// 		},
	// 		{
	// 			Name:       "employee",
	// 			PrimaryKey: []string{"id"},
	// 			Columns: []*Column{
	// 				{Name: "id", Type: INT, NotNull: true, Unsigned: true, AutoIncrement: true},
	// 				{Name: "name", Type: NVARCHAR, NotNull: true, Length: 20},
	// 				{Name: "department_id", Type: INT, Unsigned: true},
	// 				{Name: "valid", Type: SMALLINT, Default: "1", Comment: "Indicate employee status"},
	// 				{Name: "age", Type: SMALLINT, NotNull: true, Unsigned: true, Check: "age > 20"},
	// 			},
	// 			Checks: []string{"age < 50"},
	// 			ForeignKeys: []ForeignKey{
	// 				{Referer: "department_id", Reference: "department(id)"},
	// 			},
	// 		},
	// 		{
	// 			Name:       "storage",
	// 			Columns: []*Column{
	// 				{Name: "room", Type: NVARCHAR, NotNull: true, Length: 50},
	// 				{Name: "description", Type: TEXT},
	// 			},
	// 		},
	// 	},
	// }
	schema := new(Schema).WithName("company").WithComment("The Company Schema")

	department := new(Table).WithName("department").WithComment("Departments of company")
	department.AddColumn(new(Column).WithName("id").WithType(INT).IsNotNull().IsUnsigned().IsAutoIncrement())
	department.AddColumn(new(Column).WithName("name").WithType(NVARCHAR).WithLength(20).IsNotNull())
	department.AddColumn(new(Column).WithName("revenue").WithType(FLOAT).IsNotNull().IsUnsigned().WithDefault("1.01"))
	department.AddColumn(new(Column).WithName("position").WithType(SMALLINT).IsNotNull().IsUnsigned().IsUnique())
	department.AddPrimaryKey([]string{"id"})

	employee := new(Table).WithName("employee")
	employee.AddColumn(new(Column).WithName("id").WithType(INT).IsNotNull().IsUnsigned().IsAutoIncrement())
	employee.AddColumn(new(Column).WithName("name").WithType(NVARCHAR).WithLength(20).IsNotNull())
	employee.AddColumn(new(Column).WithName("department_id").WithType(INT).IsUnsigned())
	employee.AddColumn(new(Column).WithName("valid").WithType(SMALLINT).WithDefault("1").WithComment("Indicate employee status"))
	employee.AddColumn(new(Column).WithName("age").WithType(SMALLINT).IsNotNull().IsUnsigned().AddCheck("age > 20"))

	employee.AddPrimaryKey([]string{"id"})
	employee.AddCheck("age < 50")
	employee.AddForeignKey("department_id", "department(id)")

	storage := new(Table).WithName("storage").WithComment("Storage for fun")
	storage.AddColumn(new(Column).WithName("room").WithType(NVARCHAR).WithLength(50))
	storage.AddColumn(new(Column).WithName("description").WithType(TEXT))

	schema.AddTable(department)
	schema.AddTable(employee)
	schema.AddTable(storage)

	return schema
}

func setupDB(t *testing.T, dbSchema *Schema) (*sql.DB, error) {
	db, err := sql.Open(
		_platform().getDriverName(),
		_platform().getDBConnectionString(serverName, 3306, user, password, dbName),
	)
	dbSchema.SetDB(db)
	assertNotHasError(t, err)

	assertNotHasError(t, dbSchema.Drop())
	if platform == postgres || platform == mssql {
		assertFalse(t, dbSchema.IsExists())
	}
	assertFalse(t, dbSchema.HasTable("employee"))
	assertFalse(t, dbSchema.HasTable("department"))
	assertFalse(t, dbSchema.HasTable("storage"))

	assertNotHasError(t, dbSchema.Install())

	return db, err
}

func TestSchemaInstall(t *testing.T) {
	dbSchema := getSchema()
	db, err := setupDB(t, dbSchema)

	employee := _platform().getSchemaAccessName(dbSchema.Name, "employee")
	department := _platform().getSchemaAccessName(dbSchema.Name, "department")
	storage := _platform().getSchemaAccessName(dbSchema.Name, "storage")

	assertTrue(t, dbSchema.IsExists())
	assertTrue(t, dbSchema.HasTable("employee"))
	assertTrue(t, dbSchema.HasTable("department"))
	assertTrue(t, dbSchema.HasTable("storage"))

	_, err = db.Exec(fmt.Sprintf("INSERT INTO %s (name, position) VALUES ('Luan Phan Corps', 1)", department))
	assertNotHasError(t, err)
	// Checks constraint is parsed but will be ignored in mysql5.7
	// @TODO query builder will help to create query across platforms
	if platform != mysql57 {
		_, err = db.Exec(fmt.Sprintf("INSERT INTO %s (name, age, department_id) VALUES ('Luan Phan', 5, 1)", employee))
		assertHasError(t, err)

		_, err = db.Exec(fmt.Sprintf("INSERT INTO %s (name, age, department_id) VALUES ('Luan Phan', 51, 1)", employee))
		assertHasError(t, err)

		// SQLITE have type affinity, so hard to apply the text range of NVARCHAR(20) onto it
		if platform != sqlite3 {
			_, err = db.Exec(fmt.Sprintf("INSERT INTO %s (name, age) VALUES ('Luan Phan Wrong Too Looooong', 22)", employee))
			assertHasError(t, err)
		}

		_, err = db.Exec(fmt.Sprintf("INSERT INTO %s (name, age) VALUES (NULL, 22)", employee))
		assertHasError(t, err)

		_, err = db.Exec(fmt.Sprintf("INSERT INTO %s (name, age) VALUES ('Luan Phan', 22)", employee))
		assertNotHasError(t, err)
	}

	_, err = db.Exec(fmt.Sprintf("INSERT INTO %s (name, age, department_id) VALUES ('Luan Phan', 22, 1)", employee))
	assertNotHasError(t, err)

	_, err = db.Exec(fmt.Sprintf("INSERT INTO %s (room, description) VALUES ('ROOMC1', 'BOOM')", storage))
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

	assertNotHasError(t, dbSchema.Install())
}

func TestAutoIncrement(t *testing.T) {
	dbSchema := getSchema()
	db, err := setupDB(t, dbSchema)

	employee := _platform().getSchemaAccessName(dbSchema.Name, "employee")
	department := _platform().getSchemaAccessName(dbSchema.Name, "department")

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
