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

// GetTablesAt return a table in schema.Tables at `index`
func (schema *Schema) GetTablesAt(index int) *Table {
	return schema.Tables[index]
}

// GetTables return a table in schema.Tables with name
func (schema *Schema) GetTables(name string) *Table {
	for _, table := range schema.Tables {
		if table.Name == name {
			return table
		}
	}

	return nil
}
