# go-dbs [![Build Status](https://travis-ci.org/luanphandinh/go-dbs.svg?branch=master)](https://travis-ci.org/luanphandinh/go-dbs)
```
Manage databse schema.
Generate install database scripts with trasaction.
```

# Usage

The package contain some simple validation for `column`, `table`, `schema`, you could decide to use them or not, all function is contain in `_validator.go` files.

```go
dbSchema := &Schema{
    Name: "workspace",
    Tables: []Table{
        {
            "user",
            []Column{
                {Name: "id", Type: INT, Primary: true, NotNull: true, AutoIncrement: true},
                {Name: "name", Type: TEXT, NotNull: true},
            },
        },
    },
}

dbSchema.Validate() // Optional
dbSchema.Install(db)
```

Since Database and Schema a mostly the same stuff in MySQL, we will just care about tables.

# TODO

* Query Builder
* Column Types
* Support Migrate Schema
* Support check current database schema
* Support get, set, create function for tables, columns