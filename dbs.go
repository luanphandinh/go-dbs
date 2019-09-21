package dbs


var _dbPlatform dbPlatform
var _cachedPlatforms = make(map[string]dbPlatform)

// SetPlatform define the platform that entire dbs will use
// Supported platforms: sqlite3, mysql:5.7, mysql:8.0, postgres, sqlserver
func SetPlatform(platform string) {
	_dbPlatform = _getPlatform(platform)
}

func _platform() dbPlatform {
	return _dbPlatform
}

func _getPlatform(platform string) dbPlatform {
	if cached := _cachedPlatforms[platform]; cached != nil {
		return cached
	}

	cache := _makePlatform(platform)
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

