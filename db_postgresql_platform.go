package dbs

import "fmt"

type PostgresPlatform struct {
}

func (platform *PostgresPlatform) GetTypeDeclaration(col *Column) string {
	if col.Length > 0 {
		return fmt.Sprintf("%s(%d)", col.Type, col.Length)
	}

	return col.Type
}

func (platform *PostgresPlatform) GetUniqueDeclaration() string {
	return _getUniqueDeclaration()
}

func (platform *PostgresPlatform) GetNotNullDeclaration() string {
	return _getNotNullDeclaration()
}

func (platform *PostgresPlatform) GetPrimaryDeclaration(table *Table) string {
	return _getPrimaryDeclaration(table)
}

func (platform *PostgresPlatform) GetAutoIncrementDeclaration() string {
	return _getAutoIncrementDeclaration()
}

func (platform *PostgresPlatform) GetUnsignedDeclaration() string {
	return _getUnsignedDeclaration()
}

func (platform *PostgresPlatform) GetColumnDeclarationSQL(col *Column) string {
	columnString := fmt.Sprintf("%s %s", col.Name, platform.GetTypeDeclaration(col))

	if col.Unsigned {
		columnString += " " + platform.GetUnsignedDeclaration()
	}

	if col.NotNull {
		columnString += " " + platform.GetNotNullDeclaration()
	}

	if col.AutoIncrement {
		columnString += " " + platform.GetAutoIncrementDeclaration()
	}

	if col.Unique {
		columnString += " " + platform.GetUniqueDeclaration()
	}

	return columnString
}

func (platform *PostgresPlatform) GetTableCreateSQL(table *Table) (tableString string) {
	return _getTableCreateSQL(platform, table)
}

func (platform *PostgresPlatform) GetTableDropSQL(table *Table) (tableString string) {
	return _getTableDropSQL(platform, table)
}


