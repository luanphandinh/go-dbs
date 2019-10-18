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

type dbMySQLPlatform struct{}

func (platform *dbMySQLPlatform) getDriverName() string {
	return mysql
}

func (platform *dbMySQLPlatform) getDBConnectionString(server string, port int, user string, password string, dbName string) string {
	return user + ":" + password + "@tcp(" + server + ")/" + dbName
}

func (platform *dbMySQLPlatform) chainCommands(commands ...string) string {
	return concatStrings(commands, "\n")
}

func (platform *dbMySQLPlatform) getTypeDeclaration(col *Column) string {
	if col.length > 0 {
		return col.dbType + "(" + strconv.Itoa(col.length) + ")"
	}

	return col.dbType
}

func (platform *dbMySQLPlatform) getUniqueDeclaration() string {
	return _getUniqueDeclaration()
}

func (platform *dbMySQLPlatform) getNotNullDeclaration() string {
	return _getNotNullDeclaration()
}

func (platform *dbMySQLPlatform) getPrimaryDeclaration(key []string) string {
	return _getPrimaryDeclaration(key)
}

func (platform *dbMySQLPlatform) getAutoIncrementDeclaration() string {
	return "AUTO_INCREMENT"
}

func (platform *dbMySQLPlatform) getUnsignedDeclaration() string {
	return "UNSIGNED"
}

func (platform *dbMySQLPlatform) getDefaultDeclaration(expression string) string {
	return _getDefaultDeclaration(expression)
}

func (platform *dbMySQLPlatform) getColumnCommentDeclaration(expression string) string {
	return "COMMENT '" + expression + "'"
}

func (platform *dbMySQLPlatform) getColumnsCommentDeclaration(schema string, table *Table) []string {
	return make([]string, 0)
}

func (platform *dbMySQLPlatform) getColumnCheckDeclaration(expression string) string {
	return _getColumnCheckDeclaration(expression)
}

func (platform *dbMySQLPlatform) buildColumnDefinitionSQL(col *Column) string {
	return _buildColumnDefinitionSQL(platform, col)
}

func (platform *dbMySQLPlatform) buildColumnsDefinitionSQL(cols []*Column) []string {
	return _buildColumnsDefinitionSQL(platform, cols)
}

func (platform *dbMySQLPlatform) buildSchemaCreateSQL(schema *Schema) string {
	return ""
}

func (platform *dbMySQLPlatform) getSchemaCreateDeclarationSQL(schema string) string {
	return ""
}

func (platform *dbMySQLPlatform) getSchemaDropDeclarationSQL(schema string) string {
	return ""
}

func (platform *dbMySQLPlatform) getSchemaAccessName(schema string, name string) string {
	return name
}

func (platform *dbMySQLPlatform) getSchemaCommentDeclaration(schema string, expression string) string {
	return ""
}

func (platform *dbMySQLPlatform) getTableChecksDeclaration(expressions []string) []string {
	return _getTableChecksDeclaration(expressions)
}

func (platform *dbMySQLPlatform) getTableReferencesDeclarationSQL(schema string, foreignKeys []*ForeignKey) []string {
	return _getTableReferencesDeclarationSQL(platform, schema, foreignKeys)
}

func (platform *dbMySQLPlatform) getTableIndexesDeclarationSQL(schema string, table string, indexes []*TableIndex) []string {
	return _getTableIndexesDeclarationSQL(platform, schema, table, indexes)
}

func (platform *dbMySQLPlatform) getTableCommentDeclarationSQL(name string, expression string) string {
	return "COMMENT '" + expression + "'"
}

func (platform *dbMySQLPlatform) buildTableCreateSQL(schema string, table *Table) (tableString string) {
	return _buildTableCreateSQL(platform, schema, table)
}

func (platform *dbMySQLPlatform) buildTableAddColumnSQL(schema string, table string, col *Column) string {
	return "ALTER TABLE " + platform.getSchemaAccessName(schema, table) + " ADD " + platform.buildColumnDefinitionSQL(col)
}

func (platform *dbMySQLPlatform) getTableDropSQL(schema string, table string) (tableString string) {
	return _getTableDropSQL(platform, schema, table)
}

func (platform *dbMySQLPlatform) getSequenceCreateSQL(sequence string) string {
	return ""
}

func (platform *dbMySQLPlatform) getSequenceDropSQL(sequence string) string {
	return ""
}

func (platform *dbMySQLPlatform) checkSchemaExistSQL(schema string) string {
	return ""
}

func (platform *dbMySQLPlatform) checkSchemaHasTableSQL(schema string, table string) string {
	return "SHOW TABLES LIKE '" + platform.getSchemaAccessName(schema, table) + "'"
}

func (platform *dbMySQLPlatform) getSchemaTablesSQL(schema string) string {
	return "SHOW TABLES"
}

func (platform *dbMySQLPlatform) getTableColumnNamesSQL(schema string, table string) string {
	return "SELECT COLUMN_NAME FROM INFORMATION_SCHEMA.COLUMNS WHERE table_name = '" + table + "'" +
		" ORDER BY ORDINAL_POSITION ASC"
}

// This will query data from mysql and return format of
// Field | dbType 			| Null 	| Key 	| defaultValue 	| Extra
// id    | int(10) unsigned	| NO	| PRI	| NULL		| auto_increment
//		 |					| YES	| UNI	| 1			| ""
func (platform *dbMySQLPlatform) getTableColumnsSQL(schema string, table string) string {
	return "SHOW COLUMNS FROM " + platform.getSchemaAccessName(schema, table)
}

func (platform *dbMySQLPlatform) parseTableColumns(rows *sql.Rows) []*Column {
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

func (platform *dbMySQLPlatform) columnDiff(col1 *Column, col2 *Column) bool {
	return false
}

// Query
func (platform *dbMySQLPlatform) getQueryOffsetDeclaration(offset string) string {
	return "OFFSET " + offset
}

func (platform *dbMySQLPlatform) getQueryLimitDeclaration(limit string) string {
	return "LIMIT " + limit
}

func (platform *dbMySQLPlatform) getPagingDeclaration(limit string, offset string) string {
	return _getPagingDeclaration(platform, limit, offset)
}
