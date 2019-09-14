package dbs

const TABLE = "TABLE"

type ForeignKey struct {
	Referer   string `json:"referer"`
	Reference string `json:"reference"`
}

type Table struct {
	Name        string       `json:"name"`
	PrimaryKey  []string     `json:"primary_key"`
	Columns     []Column     `json:"columns"`
	Checks      []string     `json:"checks"`
	Comment     string       `json:"comment"`
	ForeignKeys []ForeignKey `json:"foreign_keys"`
}

func (table *Table) WithName(name string) *Table {
	table.Name = name

	return table
}

func (table *Table) WithComment(comment string) *Table {
	table.Comment = comment

	return table
}

func (table *Table) AddPrimaryKey(key []string) *Table {
	table.PrimaryKey = key

	return table
}

func (table *Table) AddColumn(col Column) *Table {
	table.Columns = append(table.Columns, col)

	return table
}

func (table *Table) AddColumns(cols []Column) *Table {
	table.Columns = append(table.Columns, cols...)

	return table
}

func (table *Table) AddCheck(check string) *Table {
	table.Checks = append(table.Checks, check)

	return table
}

func (table *Table) AddChecks(checks []string) *Table {
	table.Checks = append(table.Checks, checks...)

	return table
}

func (table *Table) AddForeignKey(referer string, reference string) *Table {
	table.ForeignKeys = append(table.ForeignKeys, ForeignKey{Referer: referer, Reference: reference})

	return table
}
