package dbs

import "fmt"

const MSSQL string = "sqlserver"

type MsSqlPlatform struct {
}

func (platform *MsSqlPlatform) GetDriverName() string {
	return MSSQL
}

func (platform *MsSqlPlatform) GetDBConnectionString(server string, port int, user string, password string, dbName string) string {
	return fmt.Sprintf(
		"server=%s;user id=%s;password=%s;port=1433;database=%s;",
		server,
		user,
		password,
		dbName,
	)
}

func (platform *MsSqlPlatform) ChainCommands(commands ...string) string {
	return concatString(commands, ";\nGO\n")
}

func (platform *MsSqlPlatform) GetTypeDeclaration(col *Column) string {
	return col.Type
}

func (platform *MsSqlPlatform) GetUniqueDeclaration() string {
	return _getUniqueDeclaration()
}

func (platform *MsSqlPlatform) GetNotNullDeclaration() string {
	return _getNotNullDeclaration()
}

func (platform *MsSqlPlatform) GetPrimaryDeclaration(key []string) string {
	return _getPrimaryDeclaration(key)
}

func (platform *MsSqlPlatform) GetAutoIncrementDeclaration() string {
	return "IDENTITY(1,1)"
}

func (platform *MsSqlPlatform) GetUnsignedDeclaration() string {
	return _getUnsignedDeclaration()
}

func (platform *MsSqlPlatform) BuildColumnDeclarationSQL(col *Column) string {
	columnString := fmt.Sprintf("%s %s", col.Name, platform.GetTypeDeclaration(col))

	if col.NotNull {
		columnString += " " + platform.GetNotNullDeclaration()
	}

	if col.Default != "" {
		columnString += " " + platform.GetDefaultDeclaration(col.Default)
	}

	if col.Unique {
		columnString += " " + platform.GetUniqueDeclaration()
	}

	if col.Check != "" {
		columnString += " " + platform.GetColumnCheckDeclaration(col.Check)
	}

	if col.AutoIncrement {
		columnString += " " + platform.GetAutoIncrementDeclaration()
	}

	return columnString
}

func (platform *MsSqlPlatform) BuildColumnsDeclarationSQL(cols []Column) []string {
	return _buildColumnsDeclarationSQL(platform, cols)
}

func (platform *MsSqlPlatform) GetColumnCommentDeclaration(expression string) string {
	return ""
}

func (platform *MsSqlPlatform) GetColumnsCommentDeclaration(schema string, table *Table) []string {
	return make([]string, 0)
}

func (platform *MsSqlPlatform) GetColumnCheckDeclaration(expression string) string {
	return _getColumnCheckDeclaration(expression)
}

func (platform *MsSqlPlatform) BuildSchemaCreateSQL(schema *Schema) string {
	commands := make([]string, 0)
	commands = append(commands, platform.GetSchemaCreateDeclarationSQL(schema.Name))
	if schema.Comment != "" {
		commands = append(commands, platform.GetSchemaCommentDeclaration(schema.Name, schema.Comment))
	}

	return platform.ChainCommands(commands...)
}

func (platform *MsSqlPlatform) GetSchemaCreateDeclarationSQL(schema string) string {
	return fmt.Sprintf("CREATE SCHEMA %s", schema)
}

func (platform *MsSqlPlatform) GetSchemaDropDeclarationSQL(schema string) string {
	return fmt.Sprintf("DROP SCHEMA IF EXISTS %s", schema)
}

func (platform *MsSqlPlatform) GetDefaultDeclaration(expression string) string {
	return _getDefaultDeclaration(expression)
}

func (platform *MsSqlPlatform) GetSchemaAccessName(schema string, name string) string {
	return fmt.Sprintf("%s.%s", schema, name)
}

func (platform *MsSqlPlatform) GetSchemaCommentDeclaration(schema string, expression string) string {
	return ""
}

func (platform *MsSqlPlatform) GetTableChecksDeclaration(expressions []string) []string {
	return _getTableChecksDeclaration(expressions)
}

func (platform *MsSqlPlatform) GetTableCommentDeclarationSQL(name string, expression string) string {
	return ""
}

func (platform *MsSqlPlatform) BuildTableCreateSQL(schema string, table *Table) (tableString string) {
	commands := make([]string, 0)
	commands = append(commands, _buildTableCreateSQL(platform, schema, table))

	return platform.ChainCommands(commands...)
}

func (platform *MsSqlPlatform) GetTableDropSQL(schema string, table string) (tableString string) {
	return _getTableDropSQL(platform, schema, table)
}

func (platform *MsSqlPlatform) GetSequenceCreateSQL(sequence string) string {
	return _getSequenceCreateSQL(sequence)
}

func (platform *MsSqlPlatform) GetSequenceDropSQL(sequence string) string {
	return _getSequenceDropSQL(sequence)
}
