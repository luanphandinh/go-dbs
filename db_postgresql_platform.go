package dbs

import "strconv"

const POSTGRES string = "postgres"

type PostgresPlatform struct{}

func (platform *PostgresPlatform) getDriverName() string {
	return POSTGRES
}

func (platform *PostgresPlatform) getDBConnectionString(server string, port int, user string, password string, dbName string) string {
	info := make([]string, 0)
	info = append(info, "host=" + server)
	info = append(info, "user=" + user)
	info = append(info, "password=" + password)
	info = append(info, "dbname=" + dbName)
	info = append(info, "sslmode=disable")

	return concatStrings(info, " ")
}

func (platform *PostgresPlatform) chainCommands(commands ...string) string {
	return concatStrings(commands, ";\n")
}

func (platform *PostgresPlatform) getTypeDeclaration(col *Column) string {
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

func (platform *PostgresPlatform) getUniqueDeclaration() string {
	return _getUniqueDeclaration()
}

func (platform *PostgresPlatform) getNotNullDeclaration() string {
	return _getNotNullDeclaration()
}

func (platform *PostgresPlatform) getPrimaryDeclaration(key []string) string {
	return _getPrimaryDeclaration(key)
}

func (platform *PostgresPlatform) getAutoIncrementDeclaration() string {
	return ""
}

func (platform *PostgresPlatform) getUnsignedDeclaration() string {
	return ""
}

func (platform *PostgresPlatform) buildColumnDeclarationSQL(col *Column) string {
	return _buildColumnDeclarationSQL(platform, col)
}

func (platform *PostgresPlatform) buildColumnsDeclarationSQL(cols []*Column) []string {
	return _buildColumnsDeclarationSQL(platform, cols)
}

func (platform *PostgresPlatform) getColumnCommentDeclaration(expression string) string {
	return ""
}

func (platform *PostgresPlatform) getColumnsCommentDeclaration(schema string, table *Table) []string {
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

func (platform *PostgresPlatform) getColumnCheckDeclaration(expression string) string {
	return _getColumnCheckDeclaration(expression)
}

func (platform *PostgresPlatform) buildSchemaCreateSQL(schema *Schema) string {
	commands := make([]string, 0)
	commands = append(commands, platform.getSchemaCreateDeclarationSQL(schema.Name))
	if schema.Comment != "" {
		commands = append(commands, platform.getSchemaCommentDeclaration(schema.Name, schema.Comment))
	}

	return platform.chainCommands(commands...)
}

func (platform *PostgresPlatform) getSchemaCreateDeclarationSQL(schema string) string {
	return "CREATE SCHEMA IF NOT EXISTS " + schema
}

func (platform *PostgresPlatform) getSchemaDropDeclarationSQL(schema string) string {
	return _getSchemaDropDeclarationSQL(schema)
}

func (platform *PostgresPlatform) getDefaultDeclaration(expression string) string {
	return _getDefaultDeclaration(expression)
}

func (platform *PostgresPlatform) getSchemaAccessName(schema string, name string) string {
	return schema + "." + name
}

func (platform *PostgresPlatform) getSchemaCommentDeclaration(schema string, expression string) string {
	return "COMMENT ON SCHEMA " + schema + " IS '" + expression + "'"
}

func (platform *PostgresPlatform) getTableChecksDeclaration(expressions []string) []string {
	return _getTableChecksDeclaration(expressions)
}

func (platform *PostgresPlatform) getTableReferencesDeclarationSQL(schema string, foreignKeys []ForeignKey) []string {
	return _getTableReferencesDeclarationSQL(platform, schema, foreignKeys)
}

func (platform *PostgresPlatform) getTableCommentDeclarationSQL(name string, expression string) string {
	return "COMMENT ON TABLE " + name + " IS '" + expression + "'"
}

func (platform *PostgresPlatform) buildTableCreateSQL(schema string, table *Table) (tableString string) {
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

func (platform *PostgresPlatform) getTableDropSQL(schema string, table string) (tableString string) {
	return _getTableDropSQL(platform, schema, table)
}

func (platform *PostgresPlatform) getSequenceCreateSQL(sequence string) string {
	return "CREATE SEQUENCE " + sequence
}

func (platform *PostgresPlatform) getSequenceDropSQL(sequence string) string {
	return "DROP SEQUENCE " + sequence
}
