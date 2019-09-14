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
