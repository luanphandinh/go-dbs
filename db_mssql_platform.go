package dbs

import (
	"database/sql"
	"log"
	"strconv"
	"strings"
)

const mssql string = "sqlserver"

type dbMsSQLPlatform struct{}

func (platform *dbMsSQLPlatform) getDriverName() string {
	return mssql
}

func (platform *dbMsSQLPlatform) getDBConnectionString(server string, port int, user string, password string, dbName string) string {
	info := make([]string, 0)
	info = append(info, "server="+server)
	info = append(info, "user id="+user)
	info = append(info, "password="+password)
	info = append(info, "database="+dbName)

	return concatStrings(info, ";")
}

func (platform *dbMsSQLPlatform) chainCommands(commands ...string) string {
	return concatStrings(commands, ";\nGO\n")
}

func (platform *dbMsSQLPlatform) getTypeDeclaration(col *Column) string {
	if col.length > 0 {
		return col.dbType + "(" + strconv.Itoa(col.length) + ")"
	}

	return col.dbType
}

func (platform *dbMsSQLPlatform) getUniqueDeclaration() string {
	return _getUniqueDeclaration()
}

func (platform *dbMsSQLPlatform) getNotNullDeclaration() string {
	return _getNotNullDeclaration()
}

func (platform *dbMsSQLPlatform) getPrimaryDeclaration(key []string) string {
	return _getPrimaryDeclaration(key)
}

func (platform *dbMsSQLPlatform) getAutoIncrementDeclaration() string {
	return "IDENTITY(1,1)"
}

func (platform *dbMsSQLPlatform) getUnsignedDeclaration() string {
	return ""
}

func (platform *dbMsSQLPlatform) buildColumnDefinitionSQL(col *Column) string {
	return _buildColumnDefinitionSQL(platform, col)
}

func (platform *dbMsSQLPlatform) buildColumnsDefinitionSQL(cols []*Column) []string {
	return _buildColumnsDefinitionSQL(platform, cols)
}

func (platform *dbMsSQLPlatform) getColumnCommentDeclaration(expression string) string {
	return ""
}

func (platform *dbMsSQLPlatform) getColumnsCommentDeclaration(schema string, table *Table) []string {
	return make([]string, 0)
}

func (platform *dbMsSQLPlatform) getColumnCheckDeclaration(expression string) string {
	return _getColumnCheckDeclaration(expression)
}

func (platform *dbMsSQLPlatform) buildSchemaCreateSQL(schema *Schema) string {
	return platform.getSchemaCreateDeclarationSQL(schema.name)
}

func (platform *dbMsSQLPlatform) getSchemaCreateDeclarationSQL(schema string) string {
	return "CREATE SCHEMA " + schema
}

func (platform *dbMsSQLPlatform) getSchemaDropDeclarationSQL(schema string) string {
	return "DROP SCHEMA IF EXISTS " + schema
}

func (platform *dbMsSQLPlatform) getDefaultDeclaration(expression string) string {
	return _getDefaultDeclaration(expression)
}

func (platform *dbMsSQLPlatform) getSchemaAccessName(schema string, name string) string {
	return schema + "." + name
}

func (platform *dbMsSQLPlatform) getSchemaCommentDeclaration(schema string, expression string) string {
	return ""
}

func (platform *dbMsSQLPlatform) getTableChecksDeclaration(expressions []string) []string {
	return _getTableChecksDeclaration(expressions)
}

func (platform *dbMsSQLPlatform) getTableReferencesDeclarationSQL(schema string, foreignKeys []*ForeignKey) []string {
	return _getTableReferencesDeclarationSQL(platform, schema, foreignKeys)
}

func (platform *dbMsSQLPlatform) getTableIndexesDeclarationSQL(schema string, table string, indexes []*TableIndex) []string {
	return _getTableIndexesDeclarationSQL(platform, schema, table, indexes)
}

func (platform *dbMsSQLPlatform) getTableCommentDeclarationSQL(name string, expression string) string {
	return ""
}

func (platform *dbMsSQLPlatform) buildTableCreateSQL(schema string, table *Table) string {
	return _buildTableCreateSQL(platform, schema, table)
}

func (platform *dbMsSQLPlatform) buildTableAddColumnSQL(schema string, table string, col *Column) string {
	return "ALTER TABLE " + platform.getSchemaAccessName(schema, table) + " ADD " + platform.buildColumnDefinitionSQL(col)
}

func (platform *dbMsSQLPlatform) getTableDropSQL(schema string, table string) string {
	return _getTableDropSQL(platform, schema, table)
}

func (platform *dbMsSQLPlatform) getSequenceCreateSQL(sequence string) string {
	return "CREATE SEQUENCE " + sequence
}

func (platform *dbMsSQLPlatform) getSequenceDropSQL(sequence string) string {
	return "DROP SEQUENCE " + sequence
}

func (platform *dbMsSQLPlatform) checkSchemaExistSQL(schema string) string {
	return "SELECT name FROM sys.schemas WHERE name = '" + schema + "'"
}

func (platform *dbMsSQLPlatform) checkSchemaHasTableSQL(schema string, table string) string {
	return "SELECT TABLE_NAME FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_NAME = '" + table + "' AND TABLE_SCHEMA = '" + schema + "'"
}

func (platform *dbMsSQLPlatform) getSchemaTablesSQL(schema string) string {
	return "SELECT TABLE_NAME FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_SCHEMA = '" + schema + "'"
}

func (platform *dbMsSQLPlatform) getTableColumnNamesSQL(schema string, table string) string {
	return "SELECT COLUMN_NAME FROM INFORMATION_SCHEMA.COLUMNS WHERE TABLE_NAME = '" + table +
		"' AND TABLE_SCHEMA = '" + schema + "'"
}

// https://docs.microsoft.com/en-us/sql/relational-databases/system-information-schema-views/columns-transact-sql?view=sql-server-2017
// ORDINAL_POSITION     COLUMN_NAME 	DATA_TYPE    	IS_NULLABLE 	COLUMN_DEFAULT
// ----------  			----------  	-----------  	----------  	----------
// 0           			id          	int     		NO				NULL
// 1           			name        	bit				YES				((0))
func (platform *dbMsSQLPlatform) getTableColumnsSQL(schema string, table string) string {
	return "SELECT ORDINAL_POSITION, COLUMN_NAME, DATA_TYPE, IS_NULLABLE, COLUMN_DEFAULT" +
		" FROM INFORMATION_SCHEMA.COLUMNS WHERE TABLE_NAME = '" + table +
		"' AND TABLE_SCHEMA = '" + schema + "'"
}

func (platform *dbMsSQLPlatform) parseTableColumns(rows *sql.Rows) []*Column {
	columns := make([]*Column, 0)

	var cid, field, dbType, notnull string
	var dfltValue sql.NullString
	for rows.Next() {
		err := rows.Scan(&cid, &field, &dbType, &notnull, &dfltValue)
		if err != nil {
			log.Fatal(err)
		}

		dVal := ""
		if dfltValue.Valid {
			dVal = dfltValue.String
		}

		columns = append(columns, _parseColumnMSSQL(field, dbType, notnull, dVal))
	}

	return columns
}

func _parseColumnMSSQL(field string, dbType string, notnull string, dVal string) *Column {
	col := new(Column).WithName(field)

	if dbTypeVal := strings.ToUpper(dbType); inStringArray(dbTypeVal, allTypes) {
		col.WithType(dbTypeVal)
	}

	if notnull != "NO" {
		col.IsNotNull()
	}

	col.WithDefault(dVal)

	return col
}

func (platform *dbMsSQLPlatform) columnDiff(col1 *Column, col2 *Column) bool {
	return false
}

// Query
func (platform *dbMsSQLPlatform) getQueryOffsetDeclaration(offset string) string {
	return "OFFSET " + offset + " ROWS"
}

func (platform *dbMsSQLPlatform) getQueryLimitDeclaration(limit string) string {
	return "FETCH NEXT " + limit + " ROWS ONLY"
}

func (platform *dbMsSQLPlatform) getPagingDeclaration(limit string, offset string) []byte {
	// query := make([]string, 0)
	//
	// if offset != "" {
	// 	query = append(query, platform.getQueryOffsetDeclaration(offset))
	// }
	//
	// if limit != "" {
	// 	query = append(query, platform.getQueryLimitDeclaration(limit))
	// }
	//
	// return concatStrings(query, " ")
	return nil
}
