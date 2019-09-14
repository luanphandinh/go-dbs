package dbs

import "testing"

func TestColumnDeclaration(t *testing.T) {
	mysqlPlatform := GetPlatform(MYSQL80)
	mysql57Platform := GetPlatform(MYSQL57)
	sqlitePlatform := GetPlatform(SQLITE3)
	postgresPlatform := GetPlatform(POSTGRES)
	mssqlPlatform := GetPlatform(MSSQL)

	// id := Column{
	// 	Name:          "id",
	// 	Type:          INT,
	// 	NotNull:       true,
	// 	AutoIncrement: true,
	// }
	id := new(Column).WithName("id").WithType(INT).IsNotNull().IsAutoIncrement()
	assertStringEquals(t, "id INT NOT NULL AUTO_INCREMENT", mysqlPlatform.BuildColumnDeclarationSQL(id))
	assertStringEquals(t, "id INT NOT NULL AUTO_INCREMENT", mysql57Platform.BuildColumnDeclarationSQL(id))
	assertStringEquals(t, "id INTEGER NOT NULL", sqlitePlatform.BuildColumnDeclarationSQL(id))
	assertStringEquals(t, "id INT NOT NULL", postgresPlatform.BuildColumnDeclarationSQL(id))
	assertStringEquals(t, "id INT NOT NULL IDENTITY(1,1)", mssqlPlatform.BuildColumnDeclarationSQL(id))

	name := new(Column).WithName("name").WithType(TEXT).IsNotNull()
	assertStringEquals(t, "name TEXT NOT NULL", mysqlPlatform.BuildColumnDeclarationSQL(name))
	assertStringEquals(t, "name TEXT NOT NULL", mysql57Platform.BuildColumnDeclarationSQL(name))
	assertStringEquals(t, "name TEXT NOT NULL", sqlitePlatform.BuildColumnDeclarationSQL(name))
	assertStringEquals(t, "name TEXT NOT NULL", postgresPlatform.BuildColumnDeclarationSQL(name))
	assertStringEquals(t, "name TEXT NOT NULL", mssqlPlatform.BuildColumnDeclarationSQL(name))

	// age := Column{}
	// age.Name = "age"
	// age.Type = "INT"
	// age.Unsigned = true
	age := new(Column).WithName("age").WithType(INT).IsUnsigned()
	assertStringEquals(t, "age INT UNSIGNED", mysqlPlatform.BuildColumnDeclarationSQL(age))
	assertStringEquals(t, "age INT UNSIGNED", mysql57Platform.BuildColumnDeclarationSQL(age))
	assertStringEquals(t, "age INTEGER", sqlitePlatform.BuildColumnDeclarationSQL(age))
	assertStringEquals(t, "age INT", postgresPlatform.BuildColumnDeclarationSQL(age))
	assertStringEquals(t, "age INT", mssqlPlatform.BuildColumnDeclarationSQL(age))

	age.WithLength(2)
	assertStringEquals(t, "age INT(2) UNSIGNED", mysqlPlatform.BuildColumnDeclarationSQL(age))
	assertStringEquals(t, "age INT(2) UNSIGNED", mysql57Platform.BuildColumnDeclarationSQL(age))
	assertStringEquals(t, "age INTEGER(2)", sqlitePlatform.BuildColumnDeclarationSQL(age))
	assertStringEquals(t, "age INT", postgresPlatform.BuildColumnDeclarationSQL(age))
	assertStringEquals(t, "age INT", mssqlPlatform.BuildColumnDeclarationSQL(age))

	age.WithDefault("10")
	assertStringEquals(t, "age INT(2) UNSIGNED DEFAULT 10", mysqlPlatform.BuildColumnDeclarationSQL(age))
	assertStringEquals(t, "age INT(2) UNSIGNED DEFAULT 10", mysql57Platform.BuildColumnDeclarationSQL(age))
	assertStringEquals(t, "age INTEGER(2) DEFAULT 10", sqlitePlatform.BuildColumnDeclarationSQL(age))
	assertStringEquals(t, "age INT DEFAULT 10", postgresPlatform.BuildColumnDeclarationSQL(age))
	assertStringEquals(t, "age INT DEFAULT 10", mssqlPlatform.BuildColumnDeclarationSQL(age))

	age.AddCheck("age < 150")
	assertStringEquals(t, "age INT(2) UNSIGNED DEFAULT 10 CHECK (age < 150)", mysqlPlatform.BuildColumnDeclarationSQL(age))
	assertStringEquals(t, "age INT(2) UNSIGNED DEFAULT 10 CHECK (age < 150)", mysql57Platform.BuildColumnDeclarationSQL(age))
	assertStringEquals(t, "age INTEGER(2) DEFAULT 10 CHECK (age < 150)", sqlitePlatform.BuildColumnDeclarationSQL(age))
	assertStringEquals(t, "age INT DEFAULT 10 CHECK (age < 150)", postgresPlatform.BuildColumnDeclarationSQL(age))
	assertStringEquals(t, "age INT DEFAULT 10 CHECK (age < 150)", mssqlPlatform.BuildColumnDeclarationSQL(age))

	age.WithComment("age should be less than 150")
	assertStringEquals(t, "age INT(2) UNSIGNED DEFAULT 10 CHECK (age < 150) COMMENT 'age should be less than 150'", mysqlPlatform.BuildColumnDeclarationSQL(age))
	assertStringEquals(t, "age INT(2) UNSIGNED DEFAULT 10 CHECK (age < 150) COMMENT 'age should be less than 150'", mysql57Platform.BuildColumnDeclarationSQL(age))
	assertStringEquals(t, "age INTEGER(2) DEFAULT 10 CHECK (age < 150)", sqlitePlatform.BuildColumnDeclarationSQL(age))
	assertStringEquals(t, "age INT DEFAULT 10 CHECK (age < 150)", postgresPlatform.BuildColumnDeclarationSQL(age))
	assertStringEquals(t, "age INT DEFAULT 10 CHECK (age < 150)", mssqlPlatform.BuildColumnDeclarationSQL(age))
}
