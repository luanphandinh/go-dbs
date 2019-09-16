package dbs

import "strconv"

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
	colType := col.Type

	// @TODO: make some type reference that centralized all types together across platforms
	if colType == NVARCHAR {
		 colType = VARCHAR
	}

	if col.Length > 0 {
		return colType + "(" + strconv.Itoa(col.Length) + ")"
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

func (platform *dbPostgresPlatform) buildColumnDeclarationSQL(col *Column) string {
	return _buildColumnDeclarationSQL(platform, col)
}

func (platform *dbPostgresPlatform) buildColumnsDeclarationSQL(cols []*Column) []string {
	return _buildColumnsDeclarationSQL(platform, cols)
}

func (platform *dbPostgresPlatform) getColumnCommentDeclaration(expression string) string {
	return ""
}

func (platform *dbPostgresPlatform) getColumnsCommentDeclaration(schema string, table *Table) []string {
	comments := make([]string, 0)
	tableName := platform.getSchemaAccessName(schema, table.Name)
	for _, col := range table.Columns {
		if col.Comment != "" {
			colName := tableName + "." + col.Name
			comment := " IS '" + col.Comment + "'"
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
	commands = append(commands, platform.getSchemaCreateDeclarationSQL(schema.Name))
	if schema.Comment != "" {
		commands = append(commands, platform.getSchemaCommentDeclaration(schema.Name, schema.Comment))
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

func (platform *dbPostgresPlatform) getTableReferencesDeclarationSQL(schema string, foreignKeys []ForeignKey) []string {
	return _getTableReferencesDeclarationSQL(platform, schema, foreignKeys)
}

func (platform *dbPostgresPlatform) getTableCommentDeclarationSQL(name string, expression string) string {
	return "COMMENT ON TABLE " + name + " IS '" + expression + "'"
}

func (platform *dbPostgresPlatform) buildTableCreateSQL(schema string, table *Table) (tableString string) {
	tableName := platform.getSchemaAccessName(schema, table.Name)

	commands := make([]string, 0)
	commands = append(commands, _buildTableCreateSQL(platform, schema, table))
	// Auto increment
	for _, col := range table.Columns {
		if col.AutoIncrement {
			seqName := platform.getSchemaAccessName(schema, table.Name + "_" + col.Name + "_seq")
			alterTableCommand := "ALTER TABLE " + tableName + " ALTER " + col.Name + " SET DEFAULT NEXTVAL('" + seqName + "')"
			commands = append(commands, platform.getSequenceCreateSQL(seqName))
			commands = append(commands, alterTableCommand)
		}
	}

	return platform.chainCommands(commands...)
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

func (platform *dbPostgresPlatform) checkSchemaHasTableSQL(schema string, table string) string {
	return "SELECT '" + platform.getSchemaAccessName(schema, table) + "'::regclass"
}
