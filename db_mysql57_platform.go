package dbs

import (
	"database/sql"
	"log"
	"regexp"
	"strconv"
	"strings"
)

const (
	mysql   string = "mysql"
	mysql57 string = "mysql:5.7"
)

type dbMySQL57Platform struct{}

func (platform *dbMySQL57Platform) getDriverName() string {
	return mysql
}

func (platform *dbMySQL57Platform) getDBConnectionString(server string, port int, user string, password string, dbName string) string {
	return user + ":" + password + "@tcp(" + server + ")/" + dbName
}

func (platform *dbMySQL57Platform) chainCommands(commands ...string) string {
	return concatStrings(commands, "\n")
}

func (platform *dbMySQL57Platform) getTypeDeclaration(col *Column) string {
	if col.Length > 0 {
		return col.Type + "(" + strconv.Itoa(col.Length) + ")"
	}

	return col.Type
}

func (platform *dbMySQL57Platform) getUniqueDeclaration() string {
	return _getUniqueDeclaration()
}

func (platform *dbMySQL57Platform) getNotNullDeclaration() string {
	return _getNotNullDeclaration()
}

func (platform *dbMySQL57Platform) getPrimaryDeclaration(key []string) string {
	return _getPrimaryDeclaration(key)
}

func (platform *dbMySQL57Platform) getAutoIncrementDeclaration() string {
	return "AUTO_INCREMENT"
}

func (platform *dbMySQL57Platform) getUnsignedDeclaration() string {
	return "UNSIGNED"
}

func (platform *dbMySQL57Platform) getDefaultDeclaration(expression string) string {
	return _getDefaultDeclaration(expression)
}

func (platform *dbMySQL57Platform) getColumnCommentDeclaration(expression string) string {
	return "COMMENT '" + expression + "'"
}

func (platform *dbMySQL57Platform) getColumnsCommentDeclaration(schema string, table *Table) []string {
	return make([]string, 0)
}

func (platform *dbMySQL57Platform) getColumnCheckDeclaration(expression string) string {
	return _getColumnCheckDeclaration(expression)
}

func (platform *dbMySQL57Platform) buildColumnDeclarationSQL(col *Column) string {
	return _buildColumnDeclarationSQL(platform, col)
}

func (platform *dbMySQL57Platform) buildColumnsDeclarationSQL(cols []*Column) []string {
	return _buildColumnsDeclarationSQL(platform, cols)
}

func (platform *dbMySQL57Platform) buildSchemaCreateSQL(schema *Schema) string {
	return ""
}

func (platform *dbMySQL57Platform) getSchemaCreateDeclarationSQL(schema string) string {
	return ""
}

func (platform *dbMySQL57Platform) getSchemaDropDeclarationSQL(schema string) string {
	return ""
}

func (platform *dbMySQL57Platform) getSchemaAccessName(schema string, name string) string {
	return name
}

func (platform *dbMySQL57Platform) getSchemaCommentDeclaration(schema string, expression string) string {
	return ""
}

func (platform *dbMySQL57Platform) getTableChecksDeclaration(expressions []string) []string {
	return _getTableChecksDeclaration(expressions)
}

func (platform *dbMySQL57Platform) getTableReferencesDeclarationSQL(schema string, foreignKeys []ForeignKey) []string {
	return _getTableReferencesDeclarationSQL(platform, schema, foreignKeys)
}

func (platform *dbMySQL57Platform) getTableCommentDeclarationSQL(name string, expression string) string {
	return "COMMENT '" + expression + "'"
}

func (platform *dbMySQL57Platform) buildTableCreateSQL(schema string, table *Table) (tableString string) {
	return _buildTableCreateSQL(platform, schema, table)
}

func (platform *dbMySQL57Platform) getTableDropSQL(schema string, table string) (tableString string) {
	return _getTableDropSQL(platform, schema, table)
}

func (platform *dbMySQL57Platform) getSequenceCreateSQL(sequence string) string {
	return ""
}

func (platform *dbMySQL57Platform) getSequenceDropSQL(sequence string) string {
	return ""
}

func (platform *dbMySQL57Platform) checkSchemaExistSQL(schema string) string {
	return ""
}

func (platform *dbMySQL57Platform) checkSchemaHasTableSQL(schema string, table string) string {
	return "SHOW TABLES LIKE '" + platform.getSchemaAccessName(schema, table) + "'"
}

func (platform *dbMySQL57Platform) getSchemaTablesSQL(schema string) string {
	return "SHOW TABLES"
}

// This will query data from mysql and return format of
// Field | Type 			| Null 	| Key 	| Default 	| Extra
// id    | int(10) unsigned	| NO	| PRI	| NULL		| auto_increment
//		 |					| YES	| UNI	| 1			| ""
func (platform *dbMySQL57Platform) getTableColumnsSQL(schema string , table string) string {
	return "SHOW COLUMNS FROM " + platform.getSchemaAccessName(schema, table)
}

func (platform *dbMySQL57Platform) parseTableColumns(rows *sql.Rows) []*Column {
	columns := make([]*Column, 0)

	var field, dbType, nullable, key, extra string
	var defaultVal sql.NullString
	for rows.Next() {
		err := rows.Scan(&field, &dbType, &nullable, &key, &defaultVal, &extra)
		if err != nil {
			log.Fatal(err)
		}
		dVal := ""
		if defaultVal.Valid {
			dVal = defaultVal.String
		}

		columns = append(columns, _parseColumnMySQL(field, dbType, nullable, key, dVal, extra))
	}

	return columns
}

func _parseColumnMySQL(field string, dbType string, nullable string, key string, dVal string, extra string) *Column {
	col := new(Column).WithName(field)

	dbTypes := regexp.MustCompile(`\(|\)|\s`).Split(dbType, -1)

	if key == "UNI" {
		col.IsUnique()
	}

	for _, val := range dbTypes {
		if val == "unsigned" {
			col.IsUnsigned()
		}

		if dbType := strings.ToUpper(val); inStringArray(dbType, allTypes) {
			col.WithType(dbType)
		}

		length, err := strconv.Atoi(val)
		if err == nil {
			col.WithLength(length)
		}
	}

	if nullable == "NO" {
		col.IsNotNull()
	}

	if extra == "auto_increment" {
		col.IsAutoIncrement()
	}

	col.WithDefault(dVal)

	return col
}
