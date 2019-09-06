package dbs

type SqlitePlatform struct {

}

func (platform *SqlitePlatform) GetUniqueDeclaration() string {
	return "UNIQUE"
}

func (platform *SqlitePlatform) GetNotNullDeclaration() string {
	return "NOT NULL"
}

func (platform *SqlitePlatform) GetPrimaryDeclaration() string {
	return "PRIMARY KEY"
}

func (platform *SqlitePlatform) GetAutoIncrementDeclaration() string {
	return "AUTOINCREMENT"
}

func (platform *SqlitePlatform) GetUnsignedDeclaration() string {
	return "UNSIGNED"
}

