# go-dbs [![Build Status](https://travis-ci.org/luanphandinh/go-dbs.svg?branch=master)](https://travis-ci.org/luanphandinh/go-dbs)
```
Manage databse(mysql, postgresql, sqlite3) schema.
```

# Usage

```go
dbSchema := &Schema{
    Name: "company", // Only postgresql will create a schema with this name
    Platform: "mysql", // or postgresql, sqlite3
    Tables: []Table{
        {
            Name: "employee",
            PrimaryKey: []string{"id"},
            Columns: []Column{
                {Name: "id", Type: INT, NotNull: true, Unsigned: true, AutoIncrement:true},
                {Name: "name", Type: TEXT, NotNull: true},
                {Name: "department_id", Type: INT},
                {Name: "age", Type: SMALLINT, NotNull: true, Unsigned: true, Length: 2},
            },
        },
        {
            Name: "department",
            PrimaryKey: []string{"id"},
            Columns: []Column{
                {Name: "id", Type: INT, NotNull: true, Unsigned: true, AutoIncrement:true},
                {Name: "name", Type: TEXT, NotNull: true, Length: 2},
                {Name: "rank", Type: SMALLINT, NotNull: true, Unsigned: true, Unique: true, Length: 1},
            },
        },
    },
}

dbSchema.Install(db)
```

* Since Database and Schema a mostly the same stuff in MySQL, we will just care about tables.
* Schema name will be used by postgresql.

# TODO
* Query Builder
* Support Migrate Schema
* Support check current database schema
* Support get, set, create function for tables, columns
