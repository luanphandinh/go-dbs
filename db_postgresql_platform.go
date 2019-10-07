package dbs

import (
	"database/sql"
	"log"
	"strconv"
)

const postgres string = "postgres"

type dbPostgresPlatform struct{}

func (platform *dbPostgresPlatform) getDriverName() string {
	return postgres
}

func (platform *dbPostgresPlatform) getDBConnectionString(server string, port int, user string, password string, dbName string) string {
	info := make([]string, 0)
	info = append(info, "host=" + server)
	info = append(info, "user=" + user)
	info = append(info, "password=" + password)
	info = append(info, "dbname=" + dbName)
	info = append(info, "sslmode=disable")

	return concatStrings(info, " ")
}

func (platform *dbPostgresPlatform) chainCommands(commands ...string) string {
	return concatStrings(commands, ";\n")
}

func (platform *dbPostgresPlatform) getTypeDeclaration(col *Column) string {
	colType := col.dbType

	// @TODO: make some type reference that centralized all types together across platforms
	if colType == NVARCHAR {
		 colType = VARCHAR
	}

	if col.length > 0 {
		return colType + "(" + strconv.Itoa(col.length) + ")"
	}

	return colType
}

func (platform *dbPostgresPlatform) getUniqueDeclaration() string {
	return _getUniqueDeclaration()
}

func (platform *dbPostgresPlatform) getNotNullDeclaration() string {
	return _getNotNullDeclaration()
}

func (platform *dbPostgresPlatform) getPrimaryDeclaration(key []string) string {
	return _getPrimaryDeclaration(key)
}

func (platform *dbPostgresPlatform) getAutoIncrementDeclaration() string {
	return ""
}

func (platform *dbPostgresPlatform) getUnsignedDeclaration() string {
	return ""
}

func (platform *dbPostgresPlatform) buildColumnDefinitionSQL(col *Column) string {
	return _buildColumnDefinitionSQL(platform, col)
}

func (platform *dbPostgresPlatform) buildColumnsDefinitionSQL(cols []*Column) []string {
	return _buildColumnsDefinitionSQL(platform, cols)
}

func (platform *dbPostgresPlatform) getColumnCommentDeclaration(expression string) string {
	return ""
}

func (platform *dbPostgresPlatform) getColumnsCommentDeclaration(schema string, table *Table) []string {
	comments := make([]string, 0)
	tableName := platform.getSchemaAccessName(schema, table.name)
	for _, col := range table.columns {
		if col.comment != "" {
			colName := tableName + "." + col.name
			comment := " IS '" + col.comment + "'"
			comments = append(comments, "COMMENT ON COLUMN " + colName + comment)
		}
	}

	return comments
}

func (platform *dbPostgresPlatform) getColumnCheckDeclaration(expression string) string {
	return _getColumnCheckDeclaration(expression)
}

func (platform *dbPostgresPlatform) buildSchemaCreateSQL(schema *Schema) string {
	commands := make([]string, 0)
	commands = append(commands, platform.getSchemaCreateDeclarationSQL(schema.name))
	if schema.comment != "" {
		commands = append(commands, platform.getSchemaCommentDeclaration(schema.name, schema.comment))
	}

	return platform.chainCommands(commands...)
}

func (platform *dbPostgresPlatform) getSchemaCreateDeclarationSQL(schema string) string {
	return "CREATE SCHEMA IF NOT EXISTS " + schema
}

func (platform *dbPostgresPlatform) getSchemaDropDeclarationSQL(schema string) string {
	return _getSchemaDropDeclarationSQL(schema)
}

func (platform *dbPostgresPlatform) getDefaultDeclaration(expression string) string {
	return _getDefaultDeclaration(expression)
}

func (platform *dbPostgresPlatform) getSchemaAccessName(schema string, name string) string {
	return schema + "." + name
}

func (platform *dbPostgresPlatform) getSchemaCommentDeclaration(schema string, expression string) string {
	return "COMMENT ON SCHEMA " + schema + " IS '" + expression + "'"
}

func (platform *dbPostgresPlatform) getTableChecksDeclaration(expressions []string) []string {
	return _getTableChecksDeclaration(expressions)
}

func (platform *dbPostgresPlatform) getTableReferencesDeclarationSQL(schema string, foreignKeys []*ForeignKey) []string {
	return _getTableReferencesDeclarationSQL(platform, schema, foreignKeys)
}

func (platform *dbPostgresPlatform) getTableIndexesDeclarationSQL(schema string, table string, indexes []*TableIndex) []string {
	return _getTableIndexesDeclarationSQL(platform, schema, table, indexes)
}

func (platform *dbPostgresPlatform) getTableCommentDeclarationSQL(name string, expression string) string {
	return "COMMENT ON TABLE " + name + " IS '" + expression + "'"
}

func (platform *dbPostgresPlatform) buildTableCreateSQL(schema string, table *Table) (tableString string) {
	tableName := platform.getSchemaAccessName(schema, table.name)

	commands := make([]string, 0)
	commands = append(commands, _buildTableCreateSQL(platform, schema, table))
	// Auto increment
	for _, col := range table.columns {
		if col.autoIncrement {
			seqName := platform.getSchemaAccessName(schema, table.name + "_" + col.name+  "_seq")
			alterTableCommand := "ALTER TABLE " + tableName + " ALTER " + col.name + " SET DEFAULT NEXTVAL('" + seqName + "')"
			commands = append(commands, platform.getSequenceCreateSQL(seqName))
			commands = append(commands, alterTableCommand)
		}
	}

	return platform.chainCommands(commands...)
}

func (platform *dbPostgresPlatform) buildTableAddColumnSQL(schema string, table string, col *Column) string {
	return "ALTER TABLE " + platform.getSchemaAccessName(schema, table) + " ADD COLUMN " + platform.buildColumnDefinitionSQL(col)
}

func (platform *dbPostgresPlatform) getTableDropSQL(schema string, table string) (tableString string) {
	return _getTableDropSQL(platform, schema, table)
}

func (platform *dbPostgresPlatform) getSequenceCreateSQL(sequence string) string {
	return "CREATE SEQUENCE " + sequence
}

func (platform *dbPostgresPlatform) getSequenceDropSQL(sequence string) string {
	return "DROP SEQUENCE " + sequence
}

func (platform *dbPostgresPlatform) checkSchemaExistSQL(schema string) string {
	return "SELECT schema_name FROM information_schema.schemata WHERE schema_name = '" + schema + "'"
}

func (platform *dbPostgresPlatform) checkSchemaHasTableSQL(schema string, table string) string {
	return "SELECT '" + platform.getSchemaAccessName(schema, table) + "'::regclass"
}

func (platform *dbPostgresPlatform) getSchemaTablesSQL(schema string) string {
	return "SELECT table_name FROM information_schema.tables WHERE table_type='BASE TABLE' AND table_schema='" + schema + "'"
}

func (platform *dbPostgresPlatform) getTableColumnNamesSQL(schema string, table string) string {
	return "SELECT column_name from information_schema.columns WHERE table_name = '" + table + "'" + " AND table_schema='" + schema + "'"
}

// column_name
func (platform *dbPostgresPlatform) getTableColumnsSQL(schema string , table string) string {
	return "SELECT column_name from information_schema.columns WHERE table_name = '" + table + "'" + " AND table_schema='" + schema + "'"
}

func (platform *dbPostgresPlatform) parseTableColumns(rows *sql.Rows) []*Column {
	columns := make([]*Column, 0)

	var field string
	for rows.Next() {
		err := rows.Scan(&field)
		if err != nil {
			log.Fatal(err)
		}

		columns = append(columns, new(Column).WithName(field))
	}

	return columns
}

func (platform *dbPostgresPlatform) columnDiff(col1 *Column, col2 *Column) bool {
	return false
}
