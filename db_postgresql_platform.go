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
	return _getUnsignedDeclaration()
}

func (platform *PostgresPlatform) GetColumnDeclarationSQL(col *Column) string {
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

	return columnString
}

func (platform *PostgresPlatform) GetColumnsDeclarationSQL(cols []Column) string {
	return _getColumnsDeclarationSQL(platform, cols)
}

func (platform *PostgresPlatform) GetColumnCheckDeclaration(expression string) string {
	return _getColumnCheckDeclaration(expression)
}

func (platform *PostgresPlatform) GetSchemaCreateDeclarationSQL(schema string) string {
	return _getSchemaCreateDeclarationSQL(schema)
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

func (platform *PostgresPlatform) GetTableCheckDeclaration(expressions []string) string {
	return _getTableCheckDeclaration(expressions)
}

func (platform *PostgresPlatform) GetTableCreateSQL(schema string, table *Table) (tableString string) {
	sequences := ""
	for _, col := range table.Columns {
		if col.AutoIncrement {
			seqName := fmt.Sprintf("%s_%s_seq", table.Name, col.Name)
			sequences += fmt.Sprintf(
				"; %s; ALTER TABLE %s ALTER %s SET DEFAULT NEXTVAL('%s');",
				platform.GetSequenceCreateSQL(schema, seqName),
				platform.GetSchemaAccessName(schema, table.Name),
				col.Name,
				platform.GetSchemaAccessName(schema, seqName),
			)
		}
	}

	return _getTableCreateSQL(platform, schema, table) + sequences
}

func (platform *PostgresPlatform) GetTableDropSQL(schema string, table string) (tableString string) {
	return _getTableDropSQL(platform, schema, table)
}

func (platform *PostgresPlatform) GetSequenceCreateSQL(schema string, sequence string) string {
	return _getSequenceCreateSQL(schema, sequence)
}

func (platform *PostgresPlatform) GetSequenceDropSQL(schema string, sequence string) string {
	return _getSequenceDropSQL(schema, sequence)
}
