package dbs

const TABLE = "TABLE"

type Table struct {
	Name        string   `json:"name"`
	PrimaryKey  []string `json:"primary_key"`
	Columns     []Column `json:"columns"`
	Checks      []string `json:"checks"`
	Comment     string   `json:"comment"`
	ForeignKeys []string `json:"foreign_keys"`
}
