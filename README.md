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
* Use sqlite for testing installation on CI/CD
* Column Types