package dbs

import (
	"fmt"
	"testing"
)

func assertStringEquals(t *testing.T, expected string, actual string) {
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

func compareLog(expected string, actual string)  {
	fmt.Printf("expected[-] / actual[+] \n - %s \n + %s \n", expected, actual)
}