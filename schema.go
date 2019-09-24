package dbs

// Schema defined the db schema structure
type Schema struct {
	Name    string   `json:"name"`
	Tables  []*Table `json:"tables"`
	Comment string   `json:"comment"`
}

// WithName set the schema name
func (schema *Schema) WithName(name string) *Schema {
	schema.Name = name

	return schema
}

// WithComment Set comment for schema
// This only works on postgresql
func (schema *Schema) WithComment(comment string) *Schema {
	schema.Comment = comment

	return schema
}

// AddTable add defined table to schema
func (schema *Schema) AddTable(table *Table) *Schema {
	schema.Tables = append(schema.Tables, table)

	return schema
}

// AddTables add a list of defined tables to schema
func (schema *Schema) AddTables(tables []*Table) *Schema {
	schema.Tables = append(schema.Tables, tables...)

	return schema
}
