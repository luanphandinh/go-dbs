package dbs

// Column defined the db column struct
type Column struct {
	Name          string `json:"name"`
	Type          string `json:"type"`
	NotNull       bool   `json:"not_null"`
	AutoIncrement bool   `json:"auto_increment"`
	Unsigned      bool   `json:"unsigned"`
	Unique        bool   `json:"unique"`
	Length        int    `json:"length"`
	Default       string `json:"default"`
	Check         string `json:"check"`
	Comment       string `json:"comment"`
}

// WithName set name for column.
func (col *Column) WithName(name string) *Column {
	col.Name = name

	return col
}

// WithComment set comment for column.
func (col *Column) WithComment(comment string) *Column {
	col.Comment = comment

	return col
}

// WithType define column type.
func (col *Column) WithType(dbType string) *Column {
	col.Type = dbType

	return col
}

// IsNotNull mark column as NOT NULL.
func (col *Column) IsNotNull() *Column {
	col.NotNull = true

	return col
}

// IsAutoIncrement mark column:
// 		mysql: 		AUTO_INCREMENT
// 		postgresql: GENERATE A SEQUENCE FOR THAT COLUMN
// 		msssql: 	IDENTITY(1,1)
func (col *Column) IsAutoIncrement() *Column {
	col.AutoIncrement = true

	return col
}

// IsUnsigned mark column as UNSIGNED in mysql.
func (col *Column) IsUnsigned() *Column {
	col.Unsigned = true

	return col
}

// IsUnique mark column as UNIQUE.
func (col *Column) IsUnique() *Column {
	col.Unique = true

	return col
}

// WithLength set length of column's type.
// eg: NVARCHAR(length)
func (col *Column) WithLength(length int) *Column {
	col.Length = length

	return col
}

// WithDefault set "DEFAULT" value for column.
func (col *Column) WithDefault(val string) *Column {
	col.Default = val

	return col
}

// AddCheck for column.
// eg: "age > 10"
func (col *Column) AddCheck(check string) *Column {
	col.Check = check

	return col
}
