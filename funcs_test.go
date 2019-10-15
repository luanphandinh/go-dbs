package dbs

import "testing"

func TestConcatStrings(t *testing.T) {
	assertStringEquals(t, "a,b", concatStrings([]string{"a", "b"},","))
	assertStringEquals(t, "a, b", concatStrings([]string{"a", "b"},", "))
	assertStringEquals(t, "", concatStrings([]string{},", "))
	assertStringEquals(t, "", concatStrings([]string{"", ""},", "))
	assertStringEquals(t, "a", concatStrings([]string{"a", "", ""},", "))
	assertStringEquals(t, "a, b", concatStrings([]string{"a", "", "", "b"},", "))
	assertStringEquals(t, "abc", concatStrings([]string{"a", "b", "", "c"},""))
	assertStringEquals(t, "a", concatStrings([]string{"", "a", ""}," "))
}

func TestGetContentOutOfArraySyntax(t *testing.T) {
	assertStringEquals(t, "'a', 'b'", getContentOutOfArraySyntax([]string{"a", "b"}))
	assertStringEquals(t, "1, 2", getContentOutOfArraySyntax([]int{1, 2}))
}
