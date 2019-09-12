package dbs

import "testing"

func TestConcatString(t *testing.T) {
	assertStringEquals(t, "a,b", concatString([]string{"a", "b"},","))
	assertStringEquals(t, "a, b", concatString([]string{"a", "b"},", "))
	assertStringEquals(t, "", concatString([]string{},", "))
	assertStringEquals(t, "", concatString([]string{"", ""},", "))
	assertStringEquals(t, "a", concatString([]string{"a", "", ""},", "))
}
