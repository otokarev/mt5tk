package output

import (
	"bytes"
	"fmt"
	"k8s.io/client-go/util/jsonpath"
	"sort"
	"strings"
)

func (o *output) PrintJsonpath(results interface{}, pattern string) error {
	j := jsonpath.New("A")
	err := j.Parse(pattern)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	err = j.Execute(buf, results)
	if err != nil {
		return err
	}
	sortedOut := strings.Fields(buf.String())
	sort.Strings(sortedOut)
	if err != nil {
		return err
	}

	fmt.Println(sortedOut)

	return nil
}
