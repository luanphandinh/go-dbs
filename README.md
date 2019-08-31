# go-dbs [![Build Status](https://travis-ci.org/luanphandinh/go-dbs.svg?branch=master)](https://travis-ci.org/luanphandinh/go-dbs)
```
Manage databse schema, intall database, mirgrate database, ...
```

#Usage
```go
	dbSchema := &dbs.Schema{
		Name: "workspace",
		Tables: []dbs.Table{
			{
				"user",
				[]dbs.Column{
					{"id", "INT", true, true, true},
					{"name", "NVARCHAR(50)", true, false, false},
				},
			},
		},
	}

	dbSchema.Install(db)
```
#TODO

* Query Builder
* Use sqlite for testing installation on CI/CD
* Column Types