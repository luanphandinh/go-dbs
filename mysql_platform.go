package dbs

type MySqlPlatform struct {
}

func (platform *MySqlPlatform) GetUniqueDeclaration() string {
	return "UNIQUE"
}

func (platform *MySqlPlatform) GetNotNullDeclaration() string {
	return "NOT NULL"
}

func (platform *MySqlPlatform) GetPrimaryDeclaration() string {
	return "PRIMARY KEY"
}

func (platform *MySqlPlatform) GetAutoIncrementDeclaration() string {
	return "AUTO_INCREMENT"
}

func (platform *MySqlPlatform) GetUnsignedDeclaration() string {
	return "UNSIGNED"
}
