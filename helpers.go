package dbs

import (
	"fmt"
	"testing"
)

func assertHasError(t *testing.T, err error) {
	if err == nil {
		fmt.Println("no error found.")
		t.Fail()
	}
}

func assertNotHasError(t *testing.T, err error) {
	if err != nil {
		fmt.Printf("expected[-] / actual[+] \n - %s but \n + %s \n", "nil", err.Error())
		t.Fail()
	}
}

func assertHasErrorMessage(t *testing.T, message string, err error) {
	assertHasError(t, err)

	if message != err.Error() {
		fmt.Printf("expected[-] / actual[+] \n - %s but \n + %s \n", message, err.Error())
		t.Fail()
	}
}
