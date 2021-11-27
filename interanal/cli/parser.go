package cli

import (
	"strconv"
	"strings"
)

func toIntegers(value string, delimeter string) []int64 {
	items := strings.Split(value, delimeter)
	var result []int64
	for _, val := range items {
		parsed, err := strconv.ParseInt(val, 10, 64)
		if err == nil {
			result = append(result, parsed)
		}
	}

	return result
}

func toStrings(value string, delimeter string) []string {
	items := strings.Split(value, delimeter)
	var result []string
	for _, val := range items {
		parsed := strings.TrimSpace(val)
		result = append(result, parsed)
	}

	return result
}
