package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"k8s.io/client-go/util/jsonpath"
	"sort"
	"strings"
)

func PrintResult(jsonPath string, results interface{}) error {
	if jsonPath != "" {
		out, err := ApplyJsonPath(jsonPath, results)
		if err != nil {
			return err
		}
		fmt.Println(out)
	} else {
		json, err := json.MarshalIndent(results, "", "\t")
		if err != nil {
			return err
		}
		fmt.Println(string(json))
	}
	return nil
}

func ApplyJsonPath(pattern string, results interface{}) ([]string, error) {
	j := jsonpath.New("A")
	err := j.Parse(pattern)
	if err != nil {
		return nil, err
	}
	buf := new(bytes.Buffer)
	err = j.Execute(buf, results)
	if err != nil {
		return nil, err
	}
	sortedOut := strings.Fields(buf.String())
	sort.Strings(sortedOut)

	return sortedOut, nil
}
