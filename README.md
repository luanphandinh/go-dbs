# go-dbs [![Build Status](https://travis-ci.org/luanphandinh/go-dbs.svg?branch=master)](https://travis-ci.org/luanphandinh/go-dbs)
```
Manage databse(mysql, postgresql, sqlite3) schema.
```

# Usage
Normal declaration:
```go
dbSchema := &Schema{
    Name:     "company",
    Platform: "mysql80", // or mysql57, sqlite, postgresql, sqlserver
    Tables: []*Table{
        {
            Name:       "department",
            PrimaryKey: []string{"id"},
            Columns: []*Column{
                {Name: "id", Type: INT, NotNull: true, Unsigned: true, AutoIncrement: true},
                {Name: "name", Type: NVARCHAR, NotNull: true, Length: 20},
                {Name: "revenue", Type: FLOAT, NotNull: true, Default: "1.01"},
                {Name: "position", Type: SMALLINT, NotNull: true, Unsigned: true, Unique: true},
            },
            Comment: "Departments of company",
        },
        {
            Name:       "employee",
            PrimaryKey: []string{"id"},
            Columns: []*Column{
                {Name: "id", Type: INT, NotNull: true, Unsigned: true, AutoIncrement: true},
                {Name: "name", Type: NVARCHAR, NotNull: true, Length: 20},
                {Name: "department_id", Type: INT, Unsigned: true},
                {Name: "valid", Type: SMALLINT, Default: "1", Comment: "Indicate employee status"},
                {Name: "age", Type: SMALLINT, NotNull: true, Unsigned: true, Check: "age > 20"},
            },
            Checks: []string{"age < 50"},
            ForeignKeys: []ForeignKey{
                {Referer: "department_id", Reference: "department(id)"},
            },
        },
        {
            Name:       "storage",
            Columns: []*Column{
                {Name: "room", Type: NVARCHAR, NotNull: true, Length: 50},
                {Name: "description", Type: TEXT},
            },
        },
    },
}
```

Or using builders:
```go
    schema := new(Schema).WithName("company").OnPlatform(platform).WithComment("The Company Schema")

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
```

With db is the connection(`*sql.DB`) to your database
```go
dbSchema.Install(db)
```

* Since Database and Schema a mostly the same stuff in MySQL, we will just care about tables.
* Schema name will be used by postgresql.

# TODO
* Query Builder
* Support Migrate Schema
* Support check current database schema
* Support get, set, create function for tables, columns
