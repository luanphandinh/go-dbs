package dbs

// ForeignKey of table
type ForeignKey struct {
	referer   string
	reference string
}

// TableIndex (es)
type TableIndex struct {
	cols []string
	name string
}

// Table defined db table structure
type Table struct {
	name        string
	primaryKey  []string
	columns     []*Column
	checks      []string
	comment     string
	foreignKeys []*ForeignKey
	indexes     []*TableIndex
}

// WithName set name for table
func (table *Table) WithName(name string) *Table {
	table.name = name

	return table
}

// WithComment set comment for table
func (table *Table) WithComment(comment string) *Table {
	table.comment = comment

	return table
}

// AddPrimaryKey defined primary of table
func (table *Table) AddPrimaryKey(key ...string) *Table {
	table.primaryKey = key

	return table
}

// AddColumns add defined column into table
func (table *Table) AddColumns(col ...*Column) *Table {
	table.columns = append(table.columns, col...)

	return table
}

// AddChecks to table
// eg: table.AddChecks("age > 10")
func (table *Table) AddChecks(check ...string) *Table {
	table.checks = append(table.checks, check...)

	return table
}

// AddForeignKey create a ForeignKey object and add to table declaration
func (table *Table) AddForeignKey(referer string, reference string) *Table {
	table.foreignKeys = append(table.foreignKeys, &ForeignKey{referer: referer, reference: reference})

	return table
}

// AddIndex create a TableIndex object and add to table declaration
// eg:
//		table.AddIndex("last_name", "first_name")
//		table.AddIndex("country_code", "phone_number")
func (table *Table) AddIndex(cols ...string) *Table {
	indexName := concatStrings(cols, "_")
	table.indexes = append(table.indexes, &TableIndex{name: "idx_" + table.name + "_" + indexName, cols: cols})

	return table
}
