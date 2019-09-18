package dbs

import (
	"fmt"
	"strconv"
	"testing"
)

func assertNotNil(t *testing.T, actual interface{}) {
	if actual == nil {
		compareLog("nil", actual)
		t.Fail()
	}
}

func assertTrue(t *testing.T, actual bool) {
	if actual != true {
		compareLog("true", actual)
		t.Fail()
	}
}

func assertFalse(t *testing.T, actual bool) {
	if actual != false {
		compareLog("false", actual)
		t.Fail()
	}
}

func assertStringEquals(t *testing.T, expected string, actual string) {
	if expected != actual {
		compareLog(expected, actual)
		t.Fail()
	}
}

func assertFloatEquals(t *testing.T, expected float32, actual float32) {
	if expected != actual {
		compareLog(expected, actual)
		t.Fail()
	}
}

func assertIntEquals(t *testing.T, expected int, actual int) {
	if expected != actual {
		compareLog(expected, actual)
		t.Fail()
	}
}

func assertInt64Equals(t *testing.T, expected int64, actual int64) {
	if expected != actual {
		compareLog(expected, actual)
		t.Fail()
	}
}

func assertHasError(t *testing.T, err error) {
	if err == nil {
		fmt.Println("no error found.")
		t.Fail()
	}
}

func assertNotHasError(t *testing.T, err error) {
	if err != nil {
		compareLog("nil", err.Error())
		t.Fail()
	}
}

func assertHasErrorMessage(t *testing.T, message string, err error) {
	assertHasError(t, err)

	if message != err.Error() {
		compareLog(message, err.Error())
		t.Fail()
	}
}

func assertArrayStringEquals(t *testing.T, expected []string, actual []string) {
	if len(expected) != len(actual) {
		compareLog("Expected subarray strings", "Length can not be larger than actual")
		t.Fail()
		return
	}

	for index, value := range expected {
		acVal := actual[index]
		if value != acVal {
			compareLog(
				"Expected \"" + value + "\" at index: " + strconv.Itoa(index),
				"Found \"" +acVal + "\" at index: " + strconv.Itoa(index),
			)
			t.Fail()
		}
	}
}

func compareLog(expected interface{}, actual interface{})  {
	fmt.Println("expected[-] / actual[+]")
	fmt.Println("- ", expected)
	fmt.Println("+ ", actual)
}
