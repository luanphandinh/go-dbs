package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/luanphandinh/go-dbs"
	"net/http"
)

func install(db *sql.DB) error {
	dbs.SetPlatform("mysql", db)
	schema := new(dbs.Schema).WithName("company").WithComment("The Company Schema")

	department := new(dbs.Table).WithName("department").WithComment("Departments of company")
	department.AddColumns(
		new(dbs.Column).WithName("id").WithType(dbs.INT).IsNotNull().IsUnsigned().IsAutoIncrement(),
		new(dbs.Column).WithName("name").WithType(dbs.NVARCHAR).WithLength(20).IsNotNull(),
		new(dbs.Column).WithName("revenue").WithType(dbs.FLOAT).IsNotNull().IsUnsigned().WithDefault("1.01"),
		new(dbs.Column).WithName("position").WithType(dbs.SMALLINT).IsNotNull().IsUnsigned().IsUnique(),
	)

	department.AddPrimaryKey("id", "name")
	department.AddIndex("name", "position")
	department.AddIndex("id", "name", "position")

	employee := new(dbs.Table).WithName("employee")
	employee.AddColumns(
		new(dbs.Column).WithName("id").WithType(dbs.INT).IsNotNull().IsUnsigned().IsAutoIncrement(),
		new(dbs.Column).WithName("name").WithType(dbs.NVARCHAR).WithLength(20).IsNotNull(),
		new(dbs.Column).WithName("department_id").WithType(dbs.INT).IsUnsigned(),
		new(dbs.Column).WithName("valid").WithType(dbs.SMALLINT).WithDefault("1").WithComment("Indicate employee status"),
		new(dbs.Column).WithName("age").WithType(dbs.SMALLINT).IsNotNull().IsUnsigned().AddCheck("age > 20"),
	)
	employee.AddPrimaryKey("id")
	employee.AddChecks("age < 50")
	employee.AddForeignKey("department_id", "department(id)")

	storage := new(dbs.Table).WithName("storage").WithComment("Storage for fun")
	storage.AddColumns(new(dbs.Column).WithName("room").WithType(dbs.NVARCHAR).WithLength(50))
	storage.AddColumns(new(dbs.Column).WithName("description").WithType(dbs.TEXT))

	schema.AddTables(
		department,
		employee,
		storage,
	)

	if err := dbs.Drop(schema); err != nil {
		return err
	}

	return dbs.Install(schema)
}

func getFunction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-dbType", "application/json")

	server := "localhost:3306"
	user := "admin"
	password := "admin"
	dbName := "workspace"

	connection := user + ":" + password + "@tcp(" + server + ")/" + dbName
	database, err := sql.Open("mysql", connection)
	if err != nil {
		w.Write([]byte("Database connection failed: " + err.Error()))
		w.WriteHeader(500)
		return
	}

	if err := install(database); err != nil {
		w.Write([]byte("Schema install failed: " + err.Error()))
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func main() {
	http.HandleFunc("/", getFunction)

	http.ListenAndServe(":3000", nil)
}
