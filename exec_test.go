package dbs

import (
	"database/sql"
	"testing"
)

func BenchmarkSequential(b *testing.B) {
	db, _ := sql.Open(
		_getPlatform(platform).getDriverName(),
		_getPlatform(platform).getDBConnectionString(serverName, 3306, user, password, dbName),
	)
	SetPlatform(platform, db)
	schema := getSchema()

	for i := 0; i < b.N; i++ {
		drop(schema)
		install(schema)
	}
}
