package dbs

import (
	"database/sql"
	"log"
)

var _dbPlatform dbPlatform
var _sqlDB *sql.DB
var _cachedPlatforms = make(map[string]dbPlatform)

// SetPlatform define the platform that entire dbs will use along with the database connection
// Supported platforms: sqlite3, mysql:5.7, mysql:8.0, postgres, sqlserver
func SetPlatform(platform string, db *sql.DB) {
	_dbPlatform = _getPlatform(platform)
	_sqlDB = db
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
	if platform == mysql57 {
		return new(dbMySQL57Platform)
	}

	if platform == mysql80 {
		return new(dbMySQL80Platform)
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

