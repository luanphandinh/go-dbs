package dbs

import "strings"

// Column defined the db column struct
type Column struct {
	name          string
	dbType        string
	notNull       bool
	autoIncrement bool
	unsigned      bool
	unique        bool
	length        int
	defaultValue  string
	check         string
	comment       string
}

// WithName set name for column.
func (col *Column) WithName(name string) *Column {
	col.name = name

	return col
}

// WithComment set comment for column.
func (col *Column) WithComment(comment string) *Column {
	col.comment = comment

	return col
}

// WithType define column type.
func (col *Column) WithType(dbType string) *Column {
	col.dbType = dbType

	return col
}

// IsNotNull mark column as NOT NULL.
func (col *Column) IsNotNull() *Column {
	col.notNull = true

	return col
}

// IsAutoIncrement mark column:
// 		mysql: 		AUTO_INCREMENT
// 		postgresql: GENERATE A SEQUENCE FOR THAT COLUMN
// 		msssql: 	IDENTITY(1,1)
func (col *Column) IsAutoIncrement() *Column {
	col.autoIncrement = true

	return col
}

// IsUnsigned mark column as UNSIGNED in mysql.
func (col *Column) IsUnsigned() *Column {
	col.unsigned = true

	return col
}

// IsUnique mark column as UNIQUE.
func (col *Column) IsUnique() *Column {
	col.unique = true

	return col
}

// WithLength set length of column's type.
// eg: NVARCHAR(length)
func (col *Column) WithLength(length int) *Column {
	col.length = length

	return col
}

// WithDefault set "DEFAULT" value for column.
func (col *Column) WithDefault(val string) *Column {
	col.defaultValue = val

	return col
}

// AddCheck for column.
// eg: "age > 10"
func (col *Column) AddCheck(check string) *Column {
	col.check = check

	return col
}

// @TODO: This is experiment method and have no actual value for now.
func (col *Column) diff(col2 *Column) bool {
	if _platform().getDriverName() == mysql {
		return col.diffAll(col2)
	}

	if col.name != col2.name {
		return true
	}

	return false
}

func (col *Column) diffAll(col2 *Column) bool {
	if col.name != col2.name {
		return true
	}

	// @TODO: enhance type mapping
	if ! strings.Contains(col.dbType, col2.dbType) && ! strings.Contains(col2.dbType, col.dbType) {
		return true
	}

	if col.unsigned != col2.unsigned {
		return true
	}

	if col.notNull != col.notNull {
		return true
	}

	if col.autoIncrement != col2.autoIncrement {
		return true
	}

	if col.unsigned != col2.unsigned {
		return true
	}

	// @TODO: primary and unique in mysql ???
	if col.unique != col2.unique {
		return true
	}

	if col.defaultValue != col2.defaultValue {
		return true
	}

	// @TODO: check compare

	return false
}
