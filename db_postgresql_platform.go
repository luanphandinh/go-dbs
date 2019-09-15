package dbs

const POSTGRES string = "postgres"

type PostgresPlatform struct{}

func (platform *PostgresPlatform) GetDriverName() string {
	return POSTGRES
}

func (platform *PostgresPlatform) GetDBConnectionString(server string, port int, user string, password string, dbName string) string {
	info := make([]string, 0)
	info = append(info, "host=" + server)
	info = append(info, "user=" + user)
	info = append(info, "password=" + password)
	info = append(info, "dbname=" + dbName)
	info = append(info, "sslmode=disable")

	return concatStrings(info, " ")
}

func (platform *PostgresPlatform) ChainCommands(commands ...string) string {
	return concatStrings(commands, ";\n")
}

func (platform *PostgresPlatform) GetTypeDeclaration(col *Column) string {
	return col.Type
}

func (platform *PostgresPlatform) GetUniqueDeclaration() string {
	return _getUniqueDeclaration()
}

func (platform *PostgresPlatform) GetNotNullDeclaration() string {
	return _getNotNullDeclaration()
}

func (platform *PostgresPlatform) GetPrimaryDeclaration(key []string) string {
	return _getPrimaryDeclaration(key)
}

func (platform *PostgresPlatform) GetAutoIncrementDeclaration() string {
	return ""
}

func (platform *PostgresPlatform) GetUnsignedDeclaration() string {
	return ""
}

func (platform *PostgresPlatform) BuildColumnDeclarationSQL(col *Column) string {
	return _buildColumnDeclarationSQL(platform, col)
}

func (platform *PostgresPlatform) BuildColumnsDeclarationSQL(cols []*Column) []string {
	return _buildColumnsDeclarationSQL(platform, cols)
}

func (platform *PostgresPlatform) GetColumnCommentDeclaration(expression string) string {
	return ""
}

func (platform *PostgresPlatform) GetColumnsCommentDeclaration(schema string, table *Table) []string {
	comments := make([]string, 0)
	tableName := platform.GetSchemaAccessName(schema, table.Name)
	for _, col := range table.Columns {
		if col.Comment != "" {
			colName := tableName + "." + col.Name
			comment := " IS '" + col.Comment + "'"
			comments = append(comments, "COMMENT ON COLUMN " + colName + comment)
		}
	}

	return comments
}

func (platform *PostgresPlatform) GetColumnCheckDeclaration(expression string) string {
	return _getColumnCheckDeclaration(expression)
}

func (platform *PostgresPlatform) BuildSchemaCreateSQL(schema *Schema) string {
	commands := make([]string, 0)
	commands = append(commands, platform.GetSchemaCreateDeclarationSQL(schema.Name))
	if schema.Comment != "" {
		commands = append(commands, platform.GetSchemaCommentDeclaration(schema.Name, schema.Comment))
	}

	return platform.ChainCommands(commands...)
}

func (platform *PostgresPlatform) GetSchemaCreateDeclarationSQL(schema string) string {
	return "CREATE SCHEMA IF NOT EXISTS " + schema
}

func (platform *PostgresPlatform) GetSchemaDropDeclarationSQL(schema string) string {
	return _getSchemaDropDeclarationSQL(schema)
}

func (platform *PostgresPlatform) GetDefaultDeclaration(expression string) string {
	return _getDefaultDeclaration(expression)
}

func (platform *PostgresPlatform) GetSchemaAccessName(schema string, name string) string {
	return schema + "." + name
}

func (platform *PostgresPlatform) GetSchemaCommentDeclaration(schema string, expression string) string {
	return "COMMENT ON SCHEMA " + schema + " IS '" + expression + "'"
}

func (platform *PostgresPlatform) GetTableChecksDeclaration(expressions []string) []string {
	return _getTableChecksDeclaration(expressions)
}

func (platform *PostgresPlatform) GetTableReferencesDeclarationSQL(schema string, foreignKeys []ForeignKey) []string {
	return _getTableReferencesDeclarationSQL(platform, schema, foreignKeys)
}

func (platform *PostgresPlatform) GetTableCommentDeclarationSQL(name string, expression string) string {
	return "COMMENT ON TABLE " + name + " IS '" + expression + "'"
}

func (platform *PostgresPlatform) BuildTableCreateSQL(schema string, table *Table) (tableString string) {
	tableName := platform.GetSchemaAccessName(schema, table.Name)

	commands := make([]string, 0)
	commands = append(commands, _buildTableCreateSQL(platform, schema, table))
	// Auto increment
	for _, col := range table.Columns {
		if col.AutoIncrement {
			seqName := platform.GetSchemaAccessName(schema, table.Name + "_" + col.Name + "_seq")
			alterTableCommand := "ALTER TABLE " + tableName + " ALTER " + col.Name + " SET DEFAULT NEXTVAL('" + seqName + "')"
			commands = append(commands, platform.GetSequenceCreateSQL(seqName))
			commands = append(commands, alterTableCommand)
		}
	}

	return platform.ChainCommands(commands...)
}

func (platform *PostgresPlatform) GetTableDropSQL(schema string, table string) (tableString string) {
	return _getTableDropSQL(platform, schema, table)
}

func (platform *PostgresPlatform) GetSequenceCreateSQL(sequence string) string {
	return "CREATE SEQUENCE " + sequence
}

func (platform *PostgresPlatform) GetSequenceDropSQL(sequence string) string {
	return "DROP SEQUENCE " + sequence
}
