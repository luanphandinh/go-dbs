package dbs

import "testing"

func TestColumnDeclaration(t *testing.T) {
	mysqlPlatform := getPlatform(MYSQL80)
	mysql57Platform := getPlatform(MYSQL57)
	sqlitePlatform := getPlatform(SQLITE3)
	postgresPlatform := getPlatform(POSTGRES)
	mssqlPlatform := getPlatform(MSSQL)

	// id := Column{
	// 	Name:          "id",
	// 	Type:          INT,
	// 	NotNull:       true,
	// 	AutoIncrement: true,
	// }
	id := new(Column).WithName("id").WithType(INT).IsNotNull().IsAutoIncrement()
	assertStringEquals(t, "id INT NOT NULL AUTO_INCREMENT", mysqlPlatform.buildColumnDeclarationSQL(id))
	assertStringEquals(t, "id INT NOT NULL AUTO_INCREMENT", mysql57Platform.buildColumnDeclarationSQL(id))
	assertStringEquals(t, "id INTEGER NOT NULL", sqlitePlatform.buildColumnDeclarationSQL(id))
	assertStringEquals(t, "id INT NOT NULL", postgresPlatform.buildColumnDeclarationSQL(id))
	assertStringEquals(t, "id INT NOT NULL IDENTITY(1,1)", mssqlPlatform.buildColumnDeclarationSQL(id))

	name := new(Column).WithName("name").WithType(TEXT).IsNotNull()
	assertStringEquals(t, "name TEXT NOT NULL", mysqlPlatform.buildColumnDeclarationSQL(name))
	assertStringEquals(t, "name TEXT NOT NULL", mysql57Platform.buildColumnDeclarationSQL(name))
	assertStringEquals(t, "name TEXT NOT NULL", sqlitePlatform.buildColumnDeclarationSQL(name))
	assertStringEquals(t, "name TEXT NOT NULL", postgresPlatform.buildColumnDeclarationSQL(name))
	assertStringEquals(t, "name TEXT NOT NULL", mssqlPlatform.buildColumnDeclarationSQL(name))

	// age := Column{}
	// age.Name = "age"
	// age.Type = "INT"
	// age.Unsigned = true
	age := new(Column).WithName("age").WithType(INT).IsUnsigned()
	assertStringEquals(t, "age INT UNSIGNED", mysqlPlatform.buildColumnDeclarationSQL(age))
	assertStringEquals(t, "age INT UNSIGNED", mysql57Platform.buildColumnDeclarationSQL(age))
	assertStringEquals(t, "age INTEGER", sqlitePlatform.buildColumnDeclarationSQL(age))
	assertStringEquals(t, "age INT", postgresPlatform.buildColumnDeclarationSQL(age))
	assertStringEquals(t, "age INT", mssqlPlatform.buildColumnDeclarationSQL(age))

	age.WithLength(2)
	assertStringEquals(t, "age INT(2) UNSIGNED", mysqlPlatform.buildColumnDeclarationSQL(age))
	assertStringEquals(t, "age INT(2) UNSIGNED", mysql57Platform.buildColumnDeclarationSQL(age))
	assertStringEquals(t, "age INTEGER(2)", sqlitePlatform.buildColumnDeclarationSQL(age))
	assertStringEquals(t, "age INT(2)", postgresPlatform.buildColumnDeclarationSQL(age))
	assertStringEquals(t, "age INT(2)", mssqlPlatform.buildColumnDeclarationSQL(age))

	age.WithDefault("10")
	assertStringEquals(t, "age INT(2) UNSIGNED DEFAULT 10", mysqlPlatform.buildColumnDeclarationSQL(age))
	assertStringEquals(t, "age INT(2) UNSIGNED DEFAULT 10", mysql57Platform.buildColumnDeclarationSQL(age))
	assertStringEquals(t, "age INTEGER(2) DEFAULT 10", sqlitePlatform.buildColumnDeclarationSQL(age))
	assertStringEquals(t, "age INT(2) DEFAULT 10", postgresPlatform.buildColumnDeclarationSQL(age))
	assertStringEquals(t, "age INT(2) DEFAULT 10", mssqlPlatform.buildColumnDeclarationSQL(age))

	age.AddCheck("age < 150")
	assertStringEquals(t, "age INT(2) UNSIGNED DEFAULT 10 CHECK (age < 150)", mysqlPlatform.buildColumnDeclarationSQL(age))
	assertStringEquals(t, "age INT(2) UNSIGNED DEFAULT 10 CHECK (age < 150)", mysql57Platform.buildColumnDeclarationSQL(age))
	assertStringEquals(t, "age INTEGER(2) DEFAULT 10 CHECK (age < 150)", sqlitePlatform.buildColumnDeclarationSQL(age))
	assertStringEquals(t, "age INT(2) DEFAULT 10 CHECK (age < 150)", postgresPlatform.buildColumnDeclarationSQL(age))
	assertStringEquals(t, "age INT(2) DEFAULT 10 CHECK (age < 150)", mssqlPlatform.buildColumnDeclarationSQL(age))

	age.WithComment("age should be less than 150")
	assertStringEquals(t, "age INT(2) UNSIGNED DEFAULT 10 CHECK (age < 150) COMMENT 'age should be less than 150'", mysqlPlatform.buildColumnDeclarationSQL(age))
	assertStringEquals(t, "age INT(2) UNSIGNED DEFAULT 10 CHECK (age < 150) COMMENT 'age should be less than 150'", mysql57Platform.buildColumnDeclarationSQL(age))
	assertStringEquals(t, "age INTEGER(2) DEFAULT 10 CHECK (age < 150)", sqlitePlatform.buildColumnDeclarationSQL(age))
	assertStringEquals(t, "age INT(2) DEFAULT 10 CHECK (age < 150)", postgresPlatform.buildColumnDeclarationSQL(age))
	assertStringEquals(t, "age INT(2) DEFAULT 10 CHECK (age < 150)", mssqlPlatform.buildColumnDeclarationSQL(age))
}
