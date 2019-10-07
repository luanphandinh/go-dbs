# go-dbs [![Build Status](https://travis-ci.org/luanphandinh/go-dbs.svg?branch=master)](https://travis-ci.org/luanphandinh/go-dbs)
```
Manage databse(sqlite3, mysql, postgres, sqlserver) schema.
```

# Usage
#### 1. Set platform:
```go
    // supported platforms: "sqlite3", "mysql", "postgres", "sqlserver"
    // db is your `*sql.DB`
    dbs.SetPlatform("sqlite3", db)

```

#### 2. Define schema
```go
    schema := new(dbs.Schema).WithName("company").WithComment("The Company Schema")

    department := new(dbs.Table).WithName("department").WithComment("Departments of company")
    department.AddColumn(new(dbs.Column).WithName("id").WithType(dbs.INT).IsNotNull().IsUnsigned().IsAutoIncrement())
    department.AddColumn(new(dbs.Column).WithName("name").WithType(dbs.NVARCHAR).WithLength(20).IsNotNull())
    department.AddColumn(new(dbs.Column).WithName("revenue").WithType(dbs.FLOAT).IsNotNull().IsUnsigned().WithDefault("1.01"))
    department.AddColumn(new(dbs.Column).WithName("position").WithType(dbs.SMALLINT).IsNotNull().IsUnsigned().IsUnique())
    department.AddPrimaryKey("id", "name")
    department.AddIndex("name", "position")
    department.AddIndex("id", "name", "position")

    employee := new(dbs.Table).WithName("employee")
    employee.AddColumn(new(dbs.Column).WithName("id").WithType(dbs.INT).IsNotNull().IsUnsigned().IsAutoIncrement())
    employee.AddColumn(new(dbs.Column).WithName("name").WithType(dbs.NVARCHAR).WithLength(20).IsNotNull())
    employee.AddColumn(new(dbs.Column).WithName("department_id").WithType(dbs.INT).IsUnsigned())
    employee.AddColumn(new(dbs.Column).WithName("valid").WithType(dbs.SMALLINT).WithDefault("1").WithComment("Indicate employee status"))
    employee.AddColumn(new(dbs.Column).WithName("age").WithType(dbs.SMALLINT).IsNotNull().IsUnsigned().AddCheck("age > 20"))

    employee.AddPrimaryKey("id")
    employee.AddCheck("age < 50")
    employee.AddForeignKey("department_id", "department(id)")

    storage := new(dbs.Table).WithName("storage").WithComment("Storage for fun")
    storage.AddColumn(new(dbs.Column).WithName("room").WithType(dbs.NVARCHAR).WithLength(50))
    storage.AddColumn(new(dbs.Column).WithName("description").WithType(dbs.TEXT))

    schema.AddTable(department)
    schema.AddTable(employee)
    schema.AddTable(storage)
```

#### 3. Install
```go
    dbs.Install(dbSchema)
```

* Since Database and Schema a mostly the same stuff in MySQL, we will just care about tables.
* Schema name will be used by `postgres` and `sqlserver`.

# ISSUES
* Please refer your data types as your database platform
* Currently `go-dbs` *doest not support* centralized data types across platforms
* Any error or failure will resulted int `log.Fatal()` since the schema installation is important,
so thing need to be failed as soon as possible 

# TODO
* Query Builder
* Support Migrate Schema
* Support check current database schema
