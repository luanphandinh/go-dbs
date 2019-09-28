package dbs

import (
	"fmt"
	"runtime/debug"
	"strconv"
	"testing"
)

func assertNotNil(t *testing.T, actual interface{}) {
	if actual == nil {
		failLog("nil", actual)
		t.Fail()
	}
}

func assertTrue(t *testing.T, actual bool) {
	if actual != true {
		failLog("true", actual)
		t.Fail()
	}
}

func assertFalse(t *testing.T, actual bool) {
	if actual != false {
		failLog("false", actual)
		t.Fail()
	}
}

func assertStringEquals(t *testing.T, expected string, actual string) {
	if expected != actual {
		failLog(expected, actual)
		t.Fail()
	}
}

func assertFloatEquals(t *testing.T, expected float32, actual float32) {
	if expected != actual {
		failLog(expected, actual)
		t.Fail()
	}
}

func assertIntEquals(t *testing.T, expected int, actual int) {
	if expected != actual {
		failLog(expected, actual)
		t.Fail()
	}
}

func assertInt64Equals(t *testing.T, expected int64, actual int64) {
	if expected != actual {
		failLog(expected, actual)
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
		failLog("nil", err.Error())
		t.Fail()
	}
}

func assertHasErrorMessage(t *testing.T, message string, err error) {
	assertHasError(t, err)

	if message != err.Error() {
		failLog(message, err.Error())
		t.Fail()
	}
}

func assertArrayStringEquals(t *testing.T, expected []string, actual []string) {
	if len(expected) != len(actual) {
		failLog(
			fmt.Sprintf("Expected subarray strings: %#v", expected),
			fmt.Sprintf("Found: %#v", actual),
		)
		t.Fail()
		return
	}

	for index, value := range expected {
		acVal := actual[index]
		if value != acVal {
			failLog(
				"Expected \"" + value + "\" at index: " + strconv.Itoa(index),
				"Found \"" +acVal + "\" at index: " + strconv.Itoa(index),
			)
			t.Fail()
		}
	}
}

func failLog(expected interface{}, actual interface{})  {
	fmt.Println("expected[-] / actual[+]")
	fmt.Println("- ", expected)
	fmt.Println("+ ", actual)
	fmt.Println(string(debug.Stack()))
}
