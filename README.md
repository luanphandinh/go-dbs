# go-dbs [![Build Status](https://travis-ci.org/luanphandinh/go-dbs.svg?branch=master)](https://travis-ci.org/luanphandinh/go-dbs)
```
Manage databse(mysql, postgresql, sqlite3) schema.
```

# Usage

```go
dbSchema := &Schema{
    Name:     "company",
    Platform: "mysql80", // or mysql57, sqlite, postgresql, sqlserver
    Tables: []Table{
        {
            Name:       "department",
            PrimaryKey: []string{"id"},
            Columns: []Column{
                {Name: "id", Type: INT, NotNull: true, Unsigned: true, AutoIncrement: true},
                {Name: "name", Type: TEXT, NotNull: true, Length: 2},
                {Name: "revenue", Type: FLOAT, NotNull: true, Default: "1.01"},
                {Name: "position", Type: SMALLINT, NotNull: true, Unsigned: true, Unique: true, Length: 1},
            },
            Comment: "Departments of company",
        },
        {
            Name:       "employee",
            PrimaryKey: []string{"id"},
            Columns: []Column{
                {Name: "id", Type: INT, NotNull: true, Unsigned: true, AutoIncrement: true},
                {Name: "name", Type: TEXT, NotNull: true},
                {Name: "department_id", Type: INT, Unsigned: true},
                {Name: "valid", Type: SMALLINT, Default: "1", Comment: "Indicate employee status"},
                {Name: "age", Type: SMALLINT, NotNull: true, Unsigned: true, Length: 2, Check: "age > 20"},
            },
            Checks: []string{"age < 50"},
            ForeignKeys: []ForeignKey{
                {Referer: "department_id", Reference: "department(id)"},
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
