# go-dbs [![Build Status](https://travis-ci.org/luanphandinh/go-dbs.svg?branch=master)](https://travis-ci.org/luanphandinh/go-dbs)
```
Schema install, query builder for databases(sqlite3, mysql, postgres, sqlserver).
```
## Contents
* [Schema](#schema)
    * [Set platform](#1-set-platform)
    * [Define schema](#2-define-schema)
    * [Install](#3-install)
* [Query builder](#query-builder)
    * [Select](#query-builder-select)
    * [Join](#query-builder-join)
    * [Where](#query-builder-where)
    * [Group By](#query-builder-group-by)
    * [Having](#query-builder-having)
    * [Order By](#query-builder-order-by)
    * [Offset](#query-builder-offset)
    * [Limit](#query-builder-limit)
* [Issues](#issues)
* [TODO](#todo)

<a name="schema"></a>
### Schema
<a name="1-set-platform"></a>
#### 1. Set platform:
```go
    // supported platforms: "sqlite3", "mysql", "postgres", "sqlserver"
    // db is your `*sql.DB`
    dbs.SetPlatform("sqlite3", db)

```
<a name="2-define-schema"></a>
#### 2. Define schema
```go
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
```

<a name="3-install"></a>
#### 3. Install
```go
    dbs.Install(dbSchema)
```

* Since Database and Schema a mostly the same stuff in MySQL, we will just care about tables.
* Schema name will be used by `postgres` and `sqlserver`.

<a name="query-builder"></a>
### Query Builder

<a name="query-builder-select"></a>
#### Select
```go
NewQueryBuilder().
    Select("valid", "name", "age").
    From("employee").
    GetQuery()
```

<a name="query-builder-join"></a>
#### Join
```go
NewQueryBuilder().
    Select("*").
    From("employee e").Join("department d").On("e.department_id = d.id") // or LeftJoin | RightJoin
    GetQuery()
```

<a name="query-builder-where"></a>
#### Where
* Single where
```go
NewQueryBuilder().
    Select("*").
    From("employee").
    Where("(id = %d)", 1).
    GetQuery()
```

* AndWhere | OrWhere
```go
NewQueryBuilder().
    Select("name, age").
    From("employee").
    Where("id = %d", 1).
    AndWhere("name = '%s'", "Luan Phan"). // OrWhere("name = '%s'", "Luan Phan").
    GetQuery()
```

* Mixed where query
```go
NewQueryBuilder().
    From("employee").
    Where("(id = %d AND name = ?)").
    OrWhere("department_id = ?").
    GetQuery()
```

<a name="query-builder-group-by"></a>
#### Group By
```go
NewQueryBuilder().
    Select("room, COUNT(room) as c_room").
    From("storage").
    GroupBy("room").
    GetQuery()
```

<a name="query-builder-having"></a>
#### Having
```go
NewQueryBuilder().
    Select("room, COUNT(room) as c_room").
    From("storage").
    GroupBy("room").
    Having("COUNT(room) > ?").
    GetQuery()
```

<a name="query-builder-order-by"></a>
#### Order By
```go
query = NewQueryBuilder().
    From("employee").
    Where("name = ?").
    OrderBy("id ASC, name").
    GetQuery()
```

<a name="query-builder-offset"></a>
#### Offset
This is not supported for mssql yet
```go
query = NewQueryBuilder().
    Select("*").
    From("employee").
    Offset("10").
    GetQuery()
```

<a name="query-builder-limit"></a>
#### Limit
This is not supported for mssql yet
```go
query = NewQueryBuilder().
    Select("*").
    From("employee").
    Limit("10").
    GetQuery()
```

<a name="issues"></a>
# ISSUES
* Please refer your data types as your database platform
* Currently `go-dbs` *doest not support* centralized data types across platforms
* Any error or failure will resulted int `log.Fatal()` since the schema installation is important,
so thing need to be failed as soon as possible 
* Really bad performance, need to tune up somehow when query related to params `BenchmarkQueryBuilderComplex`
```bash
go test -cpu 1 -run none -bench . -benchtime 3s
goos: darwin
goarch: amd64
pkg: github.com/luanphandinh/go-dbs
BenchmarkQueryBuilder           20000000               190 ns/op             384 B/op          2 allocs/op
BenchmarkQueryBuilderComplex    10000000               684 ns/op             624 B/op          5 allocs/op
BenchmarkRawQuery               5000000000               0.29 ns/op            0 B/op          0 allocs/op
PASS
ok      github.com/luanphandinh/go-dbs  13.690s
```
<a name="todo"></a>
# TODO
* Query Builder
* Support Migrate Schema
* Support check current database schema
