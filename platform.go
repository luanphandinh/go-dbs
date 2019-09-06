package dbs

type Platform interface {
	GetUniqueDeclaration() string
	GetNotNullDeclaration() string
	GetPrimaryDeclaration() string
	GetAutoIncrementDeclaration() string
	GetUnsignedDeclaration() string
}

func GetPlatform(platform string) Platform {
	if platform == MYSQL {
		return &MySqlPlatform{}
	}

	if platform == SQLITE3 {
		return &SqlitePlatform{}
	}

	return nil
}