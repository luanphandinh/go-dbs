package dbs

const TABLE = "TABLE"

type Table struct {
	Name       string   `json:"name"`
	PrimaryKey []string `json:"primary_key"`
	Columns    []Column `json:"columns"`
	Check      []string `json:"check"`
	Comment    string   `json:"comment"`
}
