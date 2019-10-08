package dbs

// Schema defined the db schema structure
type Schema struct {
	name    string
	tables  []*Table
	comment string
}

// WithName set the schema name
func (schema *Schema) WithName(name string) *Schema {
	schema.name = name

	return schema
}

// WithComment Set comment for schema
// This only works on postgresql
func (schema *Schema) WithComment(comment string) *Schema {
	schema.comment = comment

	return schema
}

// AddTables add defined table to schema
func (schema *Schema) AddTables(table ...*Table) *Schema {
	schema.tables = append(schema.tables, table...)

	return schema
}

// GetTableAt return a table in schema.tables at `index`
func (schema *Schema) GetTableAt(index int) *Table {
	return schema.tables[index]
}

// GetTable return a table in schema.tables with name
func (schema *Schema) GetTable(name string) *Table {
	for _, table := range schema.tables {
		if table.name == name {
			return table
		}
	}

	return nil
}
