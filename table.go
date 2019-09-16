package dbs

// ForeignKey of table
type ForeignKey struct {
	Referer   string `json:"referer"`
	Reference string `json:"reference"`
}

// Table defined db table structure
type Table struct {
	Name        string       `json:"name"`
	PrimaryKey  []string     `json:"primary_key"`
	Columns     []*Column    `json:"columns"`
	Checks      []string     `json:"checks"`
	Comment     string       `json:"comment"`
	ForeignKeys []ForeignKey `json:"foreign_keys"`
}

// WithName set name for table
func (table *Table) WithName(name string) *Table {
	table.Name = name

	return table
}

// WithComment set comment for table
func (table *Table) WithComment(comment string) *Table {
	table.Comment = comment

	return table
}

// AddPrimaryKey defined primary of table
func (table *Table) AddPrimaryKey(key []string) *Table {
	table.PrimaryKey = key

	return table
}

// AddColumn add defined column into table
func (table *Table) AddColumn(col *Column) *Table {
	table.Columns = append(table.Columns, col)

	return table
}

// AddColumns add a list of defined columns into table
func (table *Table) AddColumns(cols []*Column) *Table {
	table.Columns = append(table.Columns, cols...)

	return table
}

// AddCheck to table
// eg: table.AddCheck("age > 10")
func (table *Table) AddCheck(check string) *Table {
	table.Checks = append(table.Checks, check)

	return table
}

// AddChecks to table
// eg: table.AddCheck([]string{"age > 10", "age < 100"})
func (table *Table) AddChecks(checks []string) *Table {
	table.Checks = append(table.Checks, checks...)

	return table
}

// AddForeignKey create a ForeignKey object and add to table
func (table *Table) AddForeignKey(referer string, reference string) *Table {
	table.ForeignKeys = append(table.ForeignKeys, ForeignKey{Referer: referer, Reference: reference})

	return table
}
