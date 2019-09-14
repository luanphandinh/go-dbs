package dbs

import "testing"

func TestToTableDeclaration(t *testing.T) {
	mysqlPlatform := GetPlatform(MYSQL80)
	mysql57Platform := GetPlatform(MYSQL57)
	sqlitePlatform := GetPlatform(SQLITE3)
	postgresPlatform := GetPlatform(POSTGRES)
	msSqlPlatform := GetPlatform(MSSQL)

	id := Column{
		Name:          "id",
		Type:          INT,
		NotNull:       true,
		AutoIncrement: true,
	}

	subId := Column{
		Name:    "sub_id",
		Type:    INT,
		NotNull: true,
	}

	name := Column{
		Name:    "name",
		Type:    TEXT,
		NotNull: true,
	}

	age := Column{
		Name:    "age",
		Type:    INT,
		Length:  4,
		Default: "10",
		Check:   "age < 1000",
		Comment: "age should less than 1000",
	}

	table := Table{
		Name:       "user",
		PrimaryKey: []string{"id"},
		Columns: []Column{
			id,
			subId,
			name,
			age,
		},
		Checks:  []string{"age > 50"},
		Comment: "The user table",
		ForeignKeys: []ForeignKey{
			{Referer: "sub_id", Reference: "other_table(id)"},
		},
	}
	assertStringEquals(
		t,
`CREATE TABLE user (
	id INT NOT NULL AUTO_INCREMENT,
	sub_id INT NOT NULL,
	name TEXT NOT NULL,
	age INT(4) DEFAULT 10 CHECK (age < 1000) COMMENT 'age should less than 1000',
	PRIMARY KEY (id),
	FOREIGN KEY (sub_id) REFERENCES other_table(id),
	CHECK (age > 50)
)
COMMENT 'The user table'`,
		mysqlPlatform.BuildTableCreateSQL("", &table),
	)

	assertStringEquals(
		t,
`CREATE TABLE user (
	id INT NOT NULL AUTO_INCREMENT,
	sub_id INT NOT NULL,
	name TEXT NOT NULL,
	age INT(4) DEFAULT 10 CHECK (age < 1000) COMMENT 'age should less than 1000',
	PRIMARY KEY (id),
	FOREIGN KEY (sub_id) REFERENCES other_table(id),
	CHECK (age > 50)
)
COMMENT 'The user table'`,
		mysql57Platform.BuildTableCreateSQL("", &table),
	)

	assertStringEquals(
		t,
`CREATE TABLE user (
	id INTEGER NOT NULL,
	sub_id INTEGER NOT NULL,
	name TEXT NOT NULL,
	age INTEGER(4) DEFAULT 10 CHECK (age < 1000),
	PRIMARY KEY (id),
	FOREIGN KEY (sub_id) REFERENCES other_table(id),
	CHECK (age > 50)
)`,
		sqlitePlatform.BuildTableCreateSQL("", &table),
	)

	assertStringEquals(
		t,
`CREATE TABLE public.user (
	id INT NOT NULL,
	sub_id INT NOT NULL,
	name TEXT NOT NULL,
	age INT DEFAULT 10 CHECK (age < 1000),
	PRIMARY KEY (id),
	FOREIGN KEY (sub_id) REFERENCES public.other_table(id),
	CHECK (age > 50)
);
COMMENT ON TABLE public.user IS 'The user table';
COMMENT ON COLUMN public.user.age IS 'age should less than 1000';
CREATE SEQUENCE public.user_id_seq; ALTER TABLE public.user ALTER id SET DEFAULT NEXTVAL('public.user_id_seq')`,
		postgresPlatform.BuildTableCreateSQL("public", &table),
	)

	assertStringEquals(
		t,
`CREATE TABLE public.user (
	id INT NOT NULL IDENTITY(1,1),
	sub_id INT NOT NULL,
	name TEXT NOT NULL,
	age INT DEFAULT 10 CHECK (age < 1000),
	PRIMARY KEY (id),
	FOREIGN KEY (sub_id) REFERENCES public.other_table(id),
	CHECK (age > 50)
)`,
		msSqlPlatform.BuildTableCreateSQL("public", &table),
	)

	table.PrimaryKey = []string{"id", "name"}
	assertStringEquals(t, "PRIMARY KEY (id, name)", mysqlPlatform.GetPrimaryDeclaration(table.PrimaryKey))
	assertStringEquals(t, "PRIMARY KEY (id, name)", mysql57Platform.GetPrimaryDeclaration(table.PrimaryKey))
	assertStringEquals(t, "PRIMARY KEY (id, name)", sqlitePlatform.GetPrimaryDeclaration(table.PrimaryKey))
	assertStringEquals(t, "PRIMARY KEY (id, name)", postgresPlatform.GetPrimaryDeclaration(table.PrimaryKey))
	assertStringEquals(t, "PRIMARY KEY (id, name)", msSqlPlatform.GetPrimaryDeclaration(table.PrimaryKey))
}
