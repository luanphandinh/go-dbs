package dbs

func (schema *Schema) Validate() error  {
	for _, table := range schema.Tables {
		if err := table.Validate(); err != nil {
			return err
		}
	}

	return nil
}
