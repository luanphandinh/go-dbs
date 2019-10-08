package dbs

import (
	"database/sql"
	"log"
	"regexp"
	"strconv"
)

const sqlite3 string = "sqlite3"

type dbSqlitePlatform struct{}

func (platform *dbSqlitePlatform) getDriverName() string {
	return sqlite3
}

func (platform *dbSqlitePlatform) getDBConnectionString(server string, port int, user string, password string, dbName string) string {
	return dbName
}

func (platform *dbSqlitePlatform) chainCommands(commands ...string) string {
	return concatStrings(commands, ";\n")
}

func (platform *dbSqlitePlatform) getTypeDeclaration(col *Column) string {
	dbType := col.dbType

	// @TODO: make some type reference that centralized all types together across platforms
	if inStringArray(col.dbType, integerTypes) {
		dbType = "INTEGER"
	}

	if col.length > 0 {
		return dbType + "(" + strconv.Itoa(col.length) + ")"
	}

	return dbType
}

func (platform *dbSqlitePlatform) getUniqueDeclaration() string {
	return _getUniqueDeclaration()
}

func (platform *dbSqlitePlatform) getNotNullDeclaration() string {
	return _getNotNullDeclaration()
}

func (platform *dbSqlitePlatform) getPrimaryDeclaration(key []string) string {
	return _getPrimaryDeclaration(key)
}

func (platform *dbSqlitePlatform) getAutoIncrementDeclaration() string {
	return ""
}

func (platform *dbSqlitePlatform) getUnsignedDeclaration() string {
	return ""
}

func (platform *dbSqlitePlatform) getDefaultDeclaration(expression string) string {
	return _getDefaultDeclaration(expression)
}

func (platform *dbSqlitePlatform) getColumnCommentDeclaration(expression string) string {
	return ""
}

func (platform *dbSqlitePlatform) getColumnsCommentDeclaration(schema string, table *Table) []string {
	return make([]string, 0)
}

func (platform *dbSqlitePlatform) getColumnCheckDeclaration(expression string) string {
	return _getColumnCheckDeclaration(expression)
}

func (platform *dbSqlitePlatform) buildColumnDefinitionSQL(col *Column) string {
	return _buildColumnDefinitionSQL(platform, col)
}

func (platform *dbSqlitePlatform) buildColumnsDefinitionSQL(cols []*Column) []string {
	return _buildColumnsDefinitionSQL(platform, cols)
}

func (platform *dbSqlitePlatform) buildSchemaCreateSQL(schema *Schema) string {
	return ""
}

func (platform *dbSqlitePlatform) getSchemaCreateDeclarationSQL(schema string) string {
	return ""
}

func (platform *dbSqlitePlatform) getSchemaDropDeclarationSQL(schema string) string {
	return ""
}

func (platform *dbSqlitePlatform) getSchemaAccessName(schema string, name string) string {
	return name
}

func (platform *dbSqlitePlatform) getSchemaCommentDeclaration(schema string, expression string) string {
	return ""
}

func (platform *dbSqlitePlatform) getTableChecksDeclaration(expressions []string) []string {
	return _getTableChecksDeclaration(expressions)
}

func (platform *dbSqlitePlatform) getTableReferencesDeclarationSQL(schema string, foreignKeys []*ForeignKey) []string {
	return _getTableReferencesDeclarationSQL(platform, schema, foreignKeys)
}

func (platform *dbSqlitePlatform) getTableIndexesDeclarationSQL(schema string, table string, indexes []*TableIndex) []string {
	return _getTableIndexesDeclarationSQL(platform, schema, table, indexes)
}

func (platform *dbSqlitePlatform) getTableCommentDeclarationSQL(name string, expression string) string {
	return ""
}

func (platform *dbSqlitePlatform) buildTableCreateSQL(schema string, table *Table) (tableString string) {
	return _buildTableCreateSQL(platform, schema, table)
}

func (platform *dbSqlitePlatform) buildTableAddColumnSQL(schema string, table string, col *Column) string {
	return "ALTER TABLE " + platform.getSchemaAccessName(schema, table) + " ADD COLUMN " + platform.buildColumnDefinitionSQL(col)
}

func (platform *dbSqlitePlatform) getTableDropSQL(schema string, table string) (tableString string) {
	return _getTableDropSQL(platform, schema, table)
}

func (platform *dbSqlitePlatform) getSequenceCreateSQL(sequence string) string {
	return ""
}

func (platform *dbSqlitePlatform) getSequenceDropSQL(sequence string) string {
	return ""
}

func (platform *dbSqlitePlatform) checkSchemaExistSQL(schema string) string {
	return ""
}

func (platform *dbSqlitePlatform) checkSchemaHasTableSQL(schema string, table string) string {
	return "SELECT name FROM sqlite_master WHERE type='table' AND name='" + platform.getSchemaAccessName(schema, table) + "'"
}

func (platform *dbSqlitePlatform) getSchemaTablesSQL(schema string) string {
	return "SELECT name FROM sqlite_master WHERE type='table'"
}

func (platform *dbSqlitePlatform) getTableColumnNamesSQL(schema string, table string) string {
	return "SELECT name from pragma_table_info('" + platform.getSchemaAccessName(schema, table) + "')"
}

// cid         name        type        	notnull     dflt_value	pk
// ----------  ----------  -----------  ----------  ----------	----------
// 0           id          INTEGER     	1                      	1
// 1           name        NVARCHAR(20)	0           1          	0
func (platform *dbSqlitePlatform) getTableColumnsSQL(schema string, table string) string {
	return "PRAGMA table_info(" + platform.getSchemaAccessName(schema, table) + ")"
}

func (platform *dbSqlitePlatform) parseTableColumns(rows *sql.Rows) []*Column {
	columns := make([]*Column, 0)

	var notnull, pk bool
	var cid, field, dbType string
	var dfltValue sql.NullString
	for rows.Next() {
		err := rows.Scan(&cid, &field, &dbType, &notnull, &dfltValue, &pk)
		if err != nil {
			log.Fatal(err)
		}

		dVal := ""
		if dfltValue.Valid {
			dVal = dfltValue.String
		}

		columns = append(columns, _parseColumnMySQLite(field, dbType, notnull, dVal))
	}

	return columns
}

func _parseColumnMySQLite(field string, dbType string, notnull bool, dVal string) *Column {
	col := new(Column).WithName(field)

	dbTypes := regexp.MustCompile(`\(|\)|\s`).Split(dbType, -1)

	for _, val := range dbTypes {
		if inStringArray(val, allTypes) {
			col.WithType(val)
		}

		length, err := strconv.Atoi(val)
		if err == nil {
			col.WithLength(length)
		}
	}

	if notnull {
		col.IsNotNull()
	}

	col.WithDefault(dVal)

	return col
}

func (platform *dbSqlitePlatform) columnDiff(col1 *Column, col2 *Column) bool {
	return false
}
