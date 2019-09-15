package dbs

// https://dev.mysql.com/doc/refman/8.0/en/data-types.html
const (
	TINYINT    string = "TINYINT"
	SMALLINT   string = "SMALLINT"
	MEDIUMINT  string = "MEDIUMINT"
	INT        string = "INT"
	BIGINT     string = "BIGINT"
	DECIMAL    string = "DECIMAL"
	NUMERIC    string = "NUMERIC"
	FLOAT      string = "FLOAT"
	DOUBLE     string = "DOUBLE"
	BIT        string = "BIT"
	CHAR       string = "CHAR"
	VARCHAR    string = "VARCHAR"
	NVARCHAR   string = "NVARCHAR"
	BINARY     string = "BINARY"
	VARBINARY  string = "VARBINARY"
	TINYBLOB   string = "TINYBLOB"
	BLOB       string = "BLOB"
	MEDIUMBLOB string = "MEDIUMBLOB"
	LONGBLOB   string = "LONGBLOB"
	TINYTEXT   string = "TINYTEXT"
	TEXT       string = "TEXT"
	MEDIUMTEXT string = "MEDIUMTEXT"
	LONGTEXT   string = "LONGTEXT"
	ENUM       string = "ENUM"
	SET        string = "SET"
	DATE       string = "DATE"
	TIME       string = "TIME"
	DATETIME   string = "DATETIME"
	TIMESTAMP  string = "TIMESTAMP"
	YEAR       string = "YEAR"
)

var integerTypes = []string{
	TINYINT,
	SMALLINT,
	MEDIUMINT,
	INT,
	BIGINT,
}

var floatingTypes = []string{
	DOUBLE,
	FLOAT,
}
