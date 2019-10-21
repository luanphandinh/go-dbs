package dbs

import (
	"fmt"
	"reflect"
	"regexp"
	"unsafe"
)

func concatStrings(values []string, glue string) (s string) {
	if glue == "" {
		for _, value := range values {
			s += value
		}

		return s
	}

	firstEl := true
	for _, value := range values {
		if value == "" {
			continue
		}

		if firstEl {
			s += value
			firstEl = false
		} else {
			s += glue + value
		}
	}

	return s
}

func inStringArray(needle string, values []string) bool  {
	for _, value := range values {
		if value == needle {
			return true
		}
	}

	return false
}

func bytesToString(b []byte) string {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := reflect.StringHeader{Data: bh.Data, Len: bh.Len}

	return *(*string)(unsafe.Pointer(&sh))
}

// fmt.Sprintf("%#v", arg) will return a Go-syntax representation of the value
// eg: []string{"1", "2"}
// This func will get content inside {} and return as string
func getContentOutOfArraySyntax(arg interface{}) string {
	splitter := regexp.MustCompile("[{}]")
	parsedArg := splitter.Split(fmt.Sprintf("%#v", arg), -1)
	re := regexp.MustCompile(`"`)
	return re.ReplaceAllString(parsedArg[1], `'`)
}
