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
	schema := new(Schema).WithName("company").WithComment("The Company Schema")

	department := new(Table).WithName("department").WithComment("Departments of company")
	department.AddColumns(
		new(Column).WithName("id").WithType(INT).IsNotNull().IsUnsigned().IsAutoIncrement(),
		new(Column).WithName("name").WithType(NVARCHAR).WithLength(20).IsNotNull(),
		new(Column).WithName("revenue").WithType(FLOAT).IsNotNull().IsUnsigned().WithDefault("1.01"),
		new(Column).WithName("position").WithType(SMALLINT).IsNotNull().IsUnsigned().IsUnique(),
	)
	department.AddPrimaryKey("id")
	department.AddIndex("name", "position")
	department.AddIndex("id", "position")

	employee := new(Table).WithName("employee")
	employee.AddColumns(
		new(Column).WithName("id").WithType(INT).IsNotNull().IsUnsigned().IsAutoIncrement(),
		new(Column).WithName("name").WithType(NVARCHAR).WithLength(20).IsNotNull(),
		new(Column).WithName("department_id").WithType(INT).IsUnsigned(),
		new(Column).WithName("valid").WithType(SMALLINT).WithDefault("1").WithComment("Indicate employee status"),
		new(Column).WithName("age").WithType(SMALLINT).IsNotNull().IsUnsigned().AddCheck("age > 20"),
	)

	employee.AddPrimaryKey("id")
	employee.AddChecks("age < 50")
	employee.AddForeignKey("department_id", "department(id)")

	storage := new(Table).WithName("storage").WithComment("Storage for fun")
	storage.AddColumns(new(Column).WithName("room").WithType(NVARCHAR).WithLength(50))
	storage.AddColumns(new(Column).WithName("description").WithType(TEXT))

	schema.AddTables(
		department,
		employee,
		storage,
	)

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
		assertFalse(t, checkSchemaExists(dbSchema.name))
	}
	assertFalse(t, checkSchemaHasTableSQL(dbSchema.name, "department"))
	assertFalse(t, checkSchemaHasTableSQL(dbSchema.name, "employee"))
	assertFalse(t, checkSchemaHasTableSQL(dbSchema.name, "storage"))

	assertNotHasError(t, Install(dbSchema))

	return db
}

func TestSchemaInstall(t *testing.T) {
	dbSchema := getSchema()
	setupDB(t, dbSchema)

	assertTrue(t, checkSchemaExists(dbSchema.name))
	assertTrue(t, checkSchemaHasTableSQL(dbSchema.name, "employee"))
	assertTrue(t, checkSchemaHasTableSQL(dbSchema.name, "department"))
	assertTrue(t, checkSchemaHasTableSQL(dbSchema.name, "storage"))
	assertArrayStringEquals(
		t,
		[]string{"department", "employee", "storage"},
		fetchTables(dbSchema.name),
	)

	assertArrayStringEquals(
		t,
		[]string{"id", "name", "revenue", "position"},
		fetchTableColumnNames(dbSchema.name, "department"),
	)

	assertArrayStringEquals(
		t,
		[]string{"id", "name", "department_id", "valid", "age"},
		fetchTableColumnNames(dbSchema.name, "employee"),
	)

	assertArrayStringEquals(
		t,
		[]string{"room", "description"},
		fetchTableColumnNames(dbSchema.name, "storage"),
	)

	schemaDepartmentCols := dbSchema.tables[0].columns
	departmentCols := fetchTableColumns(dbSchema.name, "department")
	assertIntEquals(t, len(departmentCols), len(schemaDepartmentCols))
	for index, col := range departmentCols {
		assertFalse(t, schemaDepartmentCols[index].diff(col))
	}

	schemaEmployeeCols := dbSchema.tables[1].columns
	employeeCols := fetchTableColumns(dbSchema.name, "employee")
	assertIntEquals(t, len(employeeCols), len(schemaEmployeeCols))
	for index, col := range employeeCols {
		assertFalse(t, schemaEmployeeCols[index].diff(col))
	}

	schemaStorageCols := dbSchema.tables[2].columns
	storageCols := fetchTableColumns(dbSchema.name, "storage")
	assertIntEquals(t, len(storageCols), len(schemaStorageCols))
	for index, col := range storageCols {
		assertFalse(t, schemaStorageCols[index].diff(col))
	}

	// Migrate
	employee := dbSchema.GetTable("employee")
	employee.AddColumns(new(Column).WithName("health_check").WithType(SMALLINT))
	assertNotHasError(t, Install(dbSchema))

	assertTrue(t, checkSchemaExists(dbSchema.name))
	assertTrue(t, checkSchemaHasTableSQL(dbSchema.name, "employee"))
	assertTrue(t, checkSchemaHasTableSQL(dbSchema.name, "department"))
	assertTrue(t, checkSchemaHasTableSQL(dbSchema.name, "storage"))
	assertArrayStringEquals(
		t,
		[]string{"id", "name", "department_id", "valid", "age", "health_check"},
		fetchTableColumnNames(dbSchema.name, "employee"),
	)
}

func TestSchemaWorks(t *testing.T) {
	dbSchema := getSchema()
	db := setupDB(t, dbSchema)

	employee := _platform().getSchemaAccessName(dbSchema.name, "employee")
	department := _platform().getSchemaAccessName(dbSchema.name, "department")
	storage := _platform().getSchemaAccessName(dbSchema.name, "storage")

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
	}

	_, err = db.Exec(fmt.Sprintf("INSERT INTO %s (name, age, department_id) VALUES ('Luan', 22, 1)", employee))
	assertNotHasError(t, err)

	_, err = db.Exec(fmt.Sprintf("INSERT INTO %s (name, age, department_id) VALUES ('Phan', 23, 1)", employee))
	assertNotHasError(t, err)

	_, err = db.Exec(fmt.Sprintf("INSERT INTO %s (room, description) VALUES ('ROOMC1', 'BOOM')", storage))
	assertNotHasError(t, err)

	_, err = db.Exec(fmt.Sprintf("INSERT INTO %s (room, description) VALUES ('ROOMC1', 'BOOMBOOM')", storage))
	assertNotHasError(t, err)

	p1 := "?"
	p2 := "?"
	switch _platform().getDriverName() {
	case mssql:
		p1 = "@p1"
		p2 = "@p2"
	case postgres:
		p1 = "$1"
		p2 = "$2"
	}

	var valid, age int
	var name, departmentName string
	employeeQuery := NewQueryBuilder().
		Select("valid, name, age").
		From(employee).
		Where("name = " + p1).
		AndWhere("id > " + p2).
		GetQuery()

	err = db.QueryRow(employeeQuery, "Luan", 0).Scan(&valid, &name, &age)
	assertNotHasError(t, err)
	assertStringEquals(t, "Luan", name)
	assertIntEquals(t, 22, age)
	assertIntEquals(t, 1, valid)

	employeeOrderedByAgeQuery := NewQueryBuilder().
		Select("valid, name, age").
		From(employee).
		OrderBy("age DESC").
		GetQuery()

	err = db.QueryRow(employeeOrderedByAgeQuery).Scan(&valid, &name, &age)
	assertNotHasError(t, err)
	assertStringEquals(t, "Phan", name)
	assertIntEquals(t, 23, age)
	assertIntEquals(t, 1, valid)

	if _platform().getDriverName() != mssql {
		employeeOrderedByAgeWithOffsetQuery := NewQueryBuilder().
			Select("valid, name, age").
			From(employee).
			OrderBy("age DESC").
			Limit("1").
			Offset("1").
			GetQuery()

		err = db.QueryRow(employeeOrderedByAgeWithOffsetQuery).Scan(&valid, &name, &age)
		assertNotHasError(t, err)
		assertStringEquals(t, "Luan", name)
		assertIntEquals(t, 22, age)
		assertIntEquals(t, 1, valid)
	}

	// @TODO make IN query possible
	// departmentQuery := NewQueryBuilder().
	// 	Select("name").
	// 	From(department).
	// 	Where("name IN (?)").
	// 	GetQuery()
	//
	// err = db.QueryRow(departmentQuery, "Luan Phan Corps").Scan(&departmentName)
	// assertNotHasError(t, err)
	// assertStringEquals(t, "Luan Phan Corps", departmentName)

	joinQuery := NewQueryBuilder().
		Select("e.name, d.name").
		From(employee + " e").
		Join(department + " d").
		On("e.department_id = d.id").
		OrderBy("age DESC").
		GetQuery()

	err = db.QueryRow(joinQuery).Scan(&name, &departmentName)
	assertNotHasError(t, err)
	assertStringEquals(t, "Luan Phan Corps", departmentName)
	assertStringEquals(t, "Phan", name)

	joinQuery = NewQueryBuilder().
		Select("e.name, d.name").
		From(employee + " e").
		LeftJoin(department + " d").
		On("e.department_id = d.id").
		OrderBy("age DESC").
		GetQuery()

	err = db.QueryRow(joinQuery).Scan(&name, &departmentName)
	assertNotHasError(t, err)
	assertStringEquals(t, "Luan Phan Corps", departmentName)
	assertStringEquals(t, "Phan", name)

	if _platform().getDriverName() != sqlite3 { // Currently not supported RIGHT JOIN
		joinQuery = NewQueryBuilder().
			Select("e.name, d.name").
			From(employee + " e").
			RightJoin(department + " d").
			On("e.department_id = d.id").
			OrderBy("age DESC").
			GetQuery()

		err = db.QueryRow(joinQuery).Scan(&name, &departmentName)
		assertNotHasError(t, err)
		assertStringEquals(t, "Luan Phan Corps", departmentName)
		assertStringEquals(t, "Phan", name)
	}

	var storageName string
	var storageCount int
	storageQuery := NewQueryBuilder().
		Select("room, COUNT(room) as c_room").
		From(storage).
		GroupBy("room").
		GetQuery()

	err = db.QueryRow(storageQuery).Scan(&storageName, &storageCount)
	assertStringEquals(t, "ROOMC1", storageName)
	assertIntEquals(t, 2, storageCount)

	storageQuery = NewQueryBuilder().
		Select("room, COUNT(room) as c_room").
		From(storage).
		GroupBy("room").
		Having("COUNT(room) > " + p1).
		GetQuery()

	err = db.QueryRow(storageQuery, 1).Scan(&storageName, &storageCount)
	assertStringEquals(t, "ROOMC1", storageName)
	assertIntEquals(t, 2, storageCount)

	assertNotHasError(t, err)
	assertNotHasError(t, Install(dbSchema))
}
