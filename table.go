package dbs

type Table struct {
	Name       string   `json:"name"`
	PrimaryKey []string `json:"primary_key"`
	Columns    []Column `json:"columns"`
}
