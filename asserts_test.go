package dbs

import "testing"

func TestAssertArrayStringEquals(t *testing.T) {
	assertArrayStringEquals(
		t,
		[]string{"employee", "department", "storage"},
		[]string{"employee", "department", "storage"},
	)

	assertArrayStringEquals(
		t,
		[]string{"employee", "department", "storage", "yolo"},
		[]string{"employee", "department", "storage", "yolo"},
	)
}
