package dbs

import "testing"

func prepareTestTable() *Table {
	id := &Column{
		Name:          "id",
		Type:          INT,
		NotNull:       true,
		AutoIncrement: true,
	}

	subID := &Column{
		Name:    "sub_id",
		Type:    INT,
		NotNull: true,
	}

	name := &Column{
		Name:    "name",
		Type:    TEXT,
		NotNull: true,
	}

	age := &Column{
		Name:    "age",
		Type:    INT,
		Length:  4,
		Default: "10",
		Check:   "age < 1000",
		Comment: "age should less than 1000",
	}

	// Plain object
	// table := &Table{
	// 	Name:       "user",
	// 	PrimaryKey: []string{"id"},
	// 	Columns: []Column{
	// 		id,
	// 		subID,
	// 		name,
	// 		age,
	// 	},
	// 	Checks:  []string{"age > 50"},
	// 	Comment: "The user table",
	// 	ForeignKeys: []ForeignKey{
	// 		{Referer: "sub_id", Reference: "other_table(id)"},
	// 	},
	// }
	table := new(Table)
	table.WithName("user").WithComment("The user table")
	table.AddPrimaryKey([]string{"id"})
	table.AddColumn(id)
	table.AddColumns([]*Column{subID, name, age})
	table.AddForeignKey("sub_id", "other_table(id)")
	table.AddCheck("age > 50")

	return table
}

func TestToTableDeclaration(t *testing.T) {
	mysqlPlatform := _getPlatform(mysql)
	mysql57Platform := _getPlatform(mysql57)
	sqlitePlatform := _getPlatform(sqlite3)
	postgresPlatform := _getPlatform(postgres)
	msSQLPlatform := _getPlatform(mssql)

	table := prepareTestTable()

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
		mysqlPlatform.buildTableCreateSQL("", table),
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
		mysql57Platform.buildTableCreateSQL("", table),
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
		sqlitePlatform.buildTableCreateSQL("", table),
	)

	assertStringEquals(
		t,
`CREATE TABLE public.user (
	id INT NOT NULL,
	sub_id INT NOT NULL,
	name TEXT NOT NULL,
	age INT(4) DEFAULT 10 CHECK (age < 1000),
	PRIMARY KEY (id),
	FOREIGN KEY (sub_id) REFERENCES public.other_table(id),
	CHECK (age > 50)
);
COMMENT ON TABLE public.user IS 'The user table';
COMMENT ON COLUMN public.user.age IS 'age should less than 1000';
CREATE SEQUENCE public.user_id_seq;
ALTER TABLE public.user ALTER id SET DEFAULT NEXTVAL('public.user_id_seq')`,
		postgresPlatform.buildTableCreateSQL("public", table),
	)

	assertStringEquals(
		t,
`CREATE TABLE public.user (
	id INT NOT NULL IDENTITY(1,1),
	sub_id INT NOT NULL,
	name TEXT NOT NULL,
	age INT(4) DEFAULT 10 CHECK (age < 1000),
	PRIMARY KEY (id),
	FOREIGN KEY (sub_id) REFERENCES public.other_table(id),
	CHECK (age > 50)
)`,
		msSQLPlatform.buildTableCreateSQL("public", table),
	)

	table.PrimaryKey = []string{"id", "name"}
	assertStringEquals(t, "PRIMARY KEY (id, name)", mysqlPlatform.getPrimaryDeclaration(table.PrimaryKey))
	assertStringEquals(t, "PRIMARY KEY (id, name)", mysql57Platform.getPrimaryDeclaration(table.PrimaryKey))
	assertStringEquals(t, "PRIMARY KEY (id, name)", sqlitePlatform.getPrimaryDeclaration(table.PrimaryKey))
	assertStringEquals(t, "PRIMARY KEY (id, name)", postgresPlatform.getPrimaryDeclaration(table.PrimaryKey))
	assertStringEquals(t, "PRIMARY KEY (id, name)", msSQLPlatform.getPrimaryDeclaration(table.PrimaryKey))
}
