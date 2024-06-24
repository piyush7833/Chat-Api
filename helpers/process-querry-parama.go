package helpers

import (
	"strconv"
	"strings"
)

func ProcessQuerryParams(selectedValues []string, defaultValues []string, req_type string) any {
	if req_type == "number" {
		var i int
		if len(selectedValues) <= 0 {
			i, _ = strconv.Atoi(defaultValues[0])
		} else {
			i, _ = strconv.Atoi(selectedValues[0])
		}
		return i
	} else if req_type == "string" {
		if len(selectedValues) <= 0 {
			return defaultValues[0]
		} else {
			return selectedValues[0]
		}
	} else if req_type == "bool" {
		if len(selectedValues) <= 0 {
			return defaultValues[0] == "true"
		} else {
			return selectedValues[0] == "true"
		}
	} else { //req_type is array
		if len(selectedValues) <= 0 {
			selectedValues = append(selectedValues, defaultValues...)
		} else {
			selectedValues = append(selectedValues, strings.Split(selectedValues[0], ",")...)
		}
		return selectedValues
	}
}
