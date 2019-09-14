package dbs

import "fmt"

const POSTGRES string = "postgres"

type PostgresPlatform struct {
}

func (platform *PostgresPlatform) GetDriverName() string {
	return POSTGRES
}

func (platform *PostgresPlatform) GetDBConnectionString(server string, port int, user string, password string, dbName string) string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s sslmode=disable",
		server,
		user,
		password,
		dbName,
	)
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

func (platform *PostgresPlatform) BuildColumnsDeclarationSQL(cols []Column) []string {
	return _buildColumnsDeclarationSQL(platform, cols)
}

func (platform *PostgresPlatform) GetColumnCommentDeclaration(expression string) string {
	return ""
}

func (platform *PostgresPlatform) GetColumnsCommentDeclaration(schema string, table *Table) []string {
	comments := make([]string, 0)
	for _, col := range table.Columns {
		if col.Comment != "" {
			comments = append(
				comments,
				fmt.Sprintf(
					"COMMENT ON COLUMN %s.%s IS '%s'",
					platform.GetSchemaAccessName(schema, table.Name),
					col.Name,
					col.Comment,
				),
			)
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
	return fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS %s", schema)
}

func (platform *PostgresPlatform) GetSchemaDropDeclarationSQL(schema string) string {
	return _getSchemaDropDeclarationSQL(schema)
}

func (platform *PostgresPlatform) GetDefaultDeclaration(expression string) string {
	return _getDefaultDeclaration(expression)
}

func (platform *PostgresPlatform) GetSchemaAccessName(schema string, name string) string {
	return fmt.Sprintf("%s.%s", schema, name)
}

func (platform *PostgresPlatform) GetSchemaCommentDeclaration(schema string, expression string) string {
	return fmt.Sprintf("COMMENT ON SCHEMA %s IS '%s'", schema, expression)
}

func (platform *PostgresPlatform) GetTableChecksDeclaration(expressions []string) []string {
	return _getTableChecksDeclaration(expressions)
}

func (platform *PostgresPlatform) GetTableCommentDeclarationSQL(name string, expression string) string {
	return fmt.Sprintf("COMMENT ON TABLE %s IS '%s'", name, expression)
}

func (platform *PostgresPlatform) BuildTableCreateSQL(schema string, table *Table) (tableString string) {
	commands := make([]string, 0)
	commands = append(commands, _buildTableCreateSQL(platform, schema, table))
	// Auto increment
	for _, col := range table.Columns {
		if col.AutoIncrement {
			seqName := platform.GetSchemaAccessName(schema, fmt.Sprintf("%s_%s_seq", table.Name, col.Name))
			commands = append(
				commands,
				fmt.Sprintf(
					"%s; ALTER TABLE %s ALTER %s SET DEFAULT NEXTVAL('%s')",
					platform.GetSequenceCreateSQL(seqName),
					platform.GetSchemaAccessName(schema, table.Name),
					col.Name,
					seqName,
				),
			)
		}
	}

	return platform.ChainCommands(commands...)
}

func (platform *PostgresPlatform) GetTableDropSQL(schema string, table string) (tableString string) {
	return _getTableDropSQL(platform, schema, table)
}

func (platform *PostgresPlatform) GetSequenceCreateSQL(sequence string) string {
	return fmt.Sprintf("CREATE SEQUENCE %s", sequence)
}

func (platform *PostgresPlatform) GetSequenceDropSQL(sequence string) string {
	return fmt.Sprintf("DROP SEQUENCE %s", sequence)
}
