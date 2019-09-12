package dbs

const COLUMN = "COLUMN"

type Column struct {
	Name          string `json:"name"`
	Type          string `json:"type"`
	NotNull       bool   `json:"not_null"`
	AutoIncrement bool   `json:"auto_increment"`
	Unsigned      bool   `json:"unsigned"`
	Unique        bool   `json:"unique"`
	Length        int    `json:"length"`
	Default       string `json:"default"`
	Check         string `json:"check"`
	Comment       string `json:"comment"`
}
