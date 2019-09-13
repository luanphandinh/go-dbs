package dbs

import "fmt"

func concatStrings(values []string, glue string) (s string) {
	for index, value := range values {
		if value == "" {
			continue
		}

		if index == 0 {
			s += fmt.Sprintf("%s", value)
		} else {
			s += fmt.Sprintf("%s%s", glue, value)
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