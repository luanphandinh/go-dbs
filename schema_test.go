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
	// return &Schema{
	// 	name:     "company",
	// 	dbPlatform: platform,
	// 	comment:  "The Company Schema",
	// 	Tables: []*Table{
	// 		{
	// 			name:       "department",
	// 			primaryKey: []string{"id"},
	// 			columns: []*Column{
	// 				{name: "id", dbType: INT, notNull: true, unsigned: true, autoIncrement: true},
	// 				{name: "name", dbType: NVARCHAR, notNull: true, length: 20},
	// 				{name: "revenue", dbType: FLOAT, notNull: true, defaultValue: "1.01"},
	// 				{name: "position", dbType: SMALLINT, notNull: true, unsigned: true, unique: true},
	// 			},
	// 			comment: "Departments of company",
	// 		},
	// 		{
	// 			name:       "employee",
	// 			primaryKey: []string{"id"},
	// 			columns: []*Column{
	// 				{name: "id", dbType: INT, notNull: true, unsigned: true, autoIncrement: true},
	// 				{name: "name", dbType: NVARCHAR, notNull: true, length: 20},
	// 				{name: "department_id", dbType: INT, unsigned: true},
	// 				{name: "valid", dbType: SMALLINT, defaultValue: "1", comment: "Indicate employee status"},
	// 				{name: "age", dbType: SMALLINT, notNull: true, unsigned: true, check: "age > 20"},
	// 			},
	// 			checks: []string{"age < 50"},
	// 			foreignKeys: []ForeignKey{
	// 				{referer: "department_id", reference: "department(id)"},
	// 			},
	// 		},
	// 		{
	// 			name:       "storage",
	// 			columns: []*Column{
	// 				{name: "room", dbType: NVARCHAR, notNull: true, length: 50},
	// 				{name: "description", dbType: TEXT},
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

func setupDB(t *testing.T, dbSchema *Schema) *sql.DB {
	db, err := sql.Open(
		_getPlatform(platform).getDriverName(),
		_getPlatform(platform).getDBConnectionString(serverName, 3306, user, password, dbName),
	)
	SetPlatform(platform, db)
	assertNotHasError(t, err)

	assertNotHasError(t, Drop(dbSchema))
	if platform == postgres || platform == mssql {
		assertFalse(t, checkSchemaExists(dbSchema.Name))
	}
	assertFalse(t, checkSchemaHasTableSQL(dbSchema.Name, "department"))
	assertFalse(t, checkSchemaHasTableSQL(dbSchema.Name, "employee"))
	assertFalse(t, checkSchemaHasTableSQL(dbSchema.Name, "storage"))

	assertNotHasError(t, Install(dbSchema))

	return db
}

func TestSchemaInstall(t *testing.T) {
	dbSchema := getSchema()
	setupDB(t, dbSchema)

	assertTrue(t, checkSchemaExists(dbSchema.Name))
	assertTrue(t, checkSchemaHasTableSQL(dbSchema.Name, "employee"))
	assertTrue(t, checkSchemaHasTableSQL(dbSchema.Name, "department"))
	assertTrue(t, checkSchemaHasTableSQL(dbSchema.Name, "storage"))
	assertArrayStringEquals(
		t,
		[]string{"department", "employee" , "storage"},
		 fetchTables(dbSchema.Name),
	)

	assertArrayStringEquals(
		t,
		[]string{"id", "name", "revenue", "position"},
		 fetchTableColumnNames(dbSchema.Name, "department"),
	)

	assertArrayStringEquals(
		t,
		[]string{"id", "name", "department_id", "valid", "age"},
		 fetchTableColumnNames(dbSchema.Name, "employee"),
	)

	assertArrayStringEquals(
		t,
		[]string{"room", "description"},
		 fetchTableColumnNames(dbSchema.Name, "storage"),
	)

	schemaDepartmentCols := dbSchema.Tables[0].columns
	departmentCols := fetchTableColumns(dbSchema.Name, "department")
	assertIntEquals(t, len(departmentCols), len(schemaDepartmentCols))
	for index, col := range departmentCols {
		assertFalse(t, schemaDepartmentCols[index].diff(col))
	}

	schemaEmployeeCols := dbSchema.Tables[1].columns
	employeeCols := fetchTableColumns(dbSchema.Name, "employee")
	assertIntEquals(t, len(employeeCols), len(schemaEmployeeCols))
	for index, col := range employeeCols {
		assertFalse(t, schemaEmployeeCols[index].diff(col))
	}

	schemaStorageCols := dbSchema.Tables[2].columns
	storageCols := fetchTableColumns(dbSchema.Name, "storage")
	assertIntEquals(t, len(storageCols), len(schemaStorageCols))
	for index, col := range storageCols {
		assertFalse(t, schemaStorageCols[index].diff(col))
	}

	// Migrate
	employee := dbSchema.GetTables("employee")
	employee.AddColumn(new(Column).WithName("health_check").WithType(SMALLINT))
	assertNotHasError(t, Install(dbSchema))

	assertTrue(t, checkSchemaExists(dbSchema.Name))
	assertTrue(t, checkSchemaHasTableSQL(dbSchema.Name, "employee"))
	assertTrue(t, checkSchemaHasTableSQL(dbSchema.Name, "department"))
	assertTrue(t, checkSchemaHasTableSQL(dbSchema.Name, "storage"))
	assertArrayStringEquals(
		t,
		[]string{"id", "name", "department_id", "valid", "age", "health_check"},
		fetchTableColumnNames(dbSchema.Name, "employee"),
	)
}

func TestSchemaWorks(t *testing.T) {
	dbSchema := getSchema()
	db := setupDB(t, dbSchema)

	employee := _platform().getSchemaAccessName(dbSchema.Name, "employee")
	department := _platform().getSchemaAccessName(dbSchema.Name, "department")
	storage := _platform().getSchemaAccessName(dbSchema.Name, "storage")

	_, err := db.Exec(fmt.Sprintf("INSERT INTO %s (name, position) VALUES ('Luan Phan Corps', 1)", department))
	assertNotHasError(t, err)
	// checks constraint is parsed but will be ignored in mysql5.7
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

	assertNotHasError(t, Install(dbSchema))
}
