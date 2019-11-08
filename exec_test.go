package dbs

import (
	"database/sql"
	"fmt"
	"testing"
)

func getSchemaBenchmark() *Schema {
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
	// employee.AddForeignKey("department_id", "department(id)")

	for i := 0; i < 10; i++ {
		storage := new(Table).WithName(fmt.Sprintf("storage_%d", i)).WithComment("Storage for fun")
		storage.AddColumns(new(Column).WithName("room").WithType(NVARCHAR).WithLength(50))
		storage.AddColumns(new(Column).WithName("description").WithType(TEXT))
		schema.AddTables(storage)
	}

	schema.AddTables(
		department,
		employee,
	)

	return schema
}

func BenchmarkSequential(b *testing.B) {
	db, _ := sql.Open(
		_getPlatform(platform).getDriverName(),
		_getPlatform(platform).getDBConnectionString(serverName, 3306, user, password, dbName),
	)
	SetPlatform(platform, db)
	schema := getSchemaBenchmark()
	drop(schema)
	for i := 0; i < b.N; i++ {
		install(schema)
	}
}

func BenchmarkConcurrent(b *testing.B) {
	db, _ := sql.Open(
		_getPlatform(platform).getDriverName(),
		_getPlatform(platform).getDBConnectionString(serverName, 3306, user, password, dbName),
	)
	SetPlatform(platform, db)
	schema := getSchemaBenchmark()
	drop(schema)
	for i := 0; i < b.N; i++ {
		installConcurrent(schema)
	}
}
