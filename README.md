# go-dbs [![Build Status](https://travis-ci.org/luanphandinh/go-dbs.svg?branch=master)](https://travis-ci.org/luanphandinh/go-dbs)
```
Manage databse schema, intall database, mirgrate database, ...
```

# Usage
```go
dbSchema := &Schema{
    Name: "workspace",
    Tables: []Table{
        {
            "user",
            []Column{
                {Name: "id", Type: "INT", Primary: true, NotNull: true, AutoIncrement: true},
                {Name: "name", Type: "NVARCHAR(50)", NotNull: true},
            },
        },
    },
}

dbSchema.Install(db)
```
# TODO

* Query Builder
* Column Types
* Support Migrate Schema
* Support check current database schema
* Support get, set, create function for tables, columns