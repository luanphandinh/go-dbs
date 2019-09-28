package dbs

import (
	"database/sql"
	"log"
)

var _dbPlatform dbPlatform
var _sqlDB *sql.DB
var _cachedPlatforms = make(map[string]dbPlatform)

// SetPlatform define the platform that entire dbs will use along with the database connection
// Supported platforms: sqlite3, mysql, postgres, sqlserver
func SetPlatform(platform string, db *sql.DB) {
	_dbPlatform = _getPlatform(platform)
	_sqlDB = db
}

// Install schema
func Install(schema *Schema) error {
	return install(schema)
}

// Drop schema
func Drop(schema *Schema) error {
	return drop(schema)
}

func _platform() dbPlatform {
	return _dbPlatform
}

func _db() *sql.DB {
	return _sqlDB
}

func _getPlatform(platform string) dbPlatform {
	if cached := _cachedPlatforms[platform]; cached != nil {
		return cached
	}

	cache := _makePlatform(platform)
	if cache == nil {
		log.Fatal("platform not supported" + platform)
	}

	_cachedPlatforms[platform] = cache

	return cache
}

func _makePlatform(platform string) dbPlatform {
	if platform == mysql57 || platform == mysql {
		return new(dbMySQLPlatform)
	}

	if platform == sqlite3 {
		return new(dbSqlitePlatform)
	}

	if platform == postgres {
		return new(dbPostgresPlatform)
	}

	if platform == mssql {
		return new(dbMsSQLPlatform)
	}

	return nil
}
