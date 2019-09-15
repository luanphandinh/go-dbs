package dbs

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

func (col *Column) WithName(name string) *Column {
	col.Name = name

	return col
}

func (col *Column) WithComment(comment string) *Column {
	col.Comment = comment

	return col
}

func (col *Column) WithType(dbType string) *Column {
	col.Type = dbType

	return col
}

func (col *Column) IsNotNull() *Column {
	col.NotNull = true

	return col
}

func (col *Column) IsAutoIncrement() *Column {
	col.AutoIncrement = true

	return col
}

func (col *Column) IsUnsigned() *Column {
	col.Unsigned = true

	return col
}

func (col *Column) IsUnique() *Column {
	col.Unique = true

	return col
}

func (col *Column) WithLength(length int) *Column {
	col.Length = length

	return col
}

func (col *Column) WithDefault(val string) *Column {
	col.Default = val

	return col
}

func (col *Column) AddCheck(check string) *Column {
	col.Check = check

	return col
}
