package output

import (
	"encoding/json"
	"fmt"
)

func (o *output) PrintJson(results interface{}, pattern string) error {
	j, err := json.MarshalIndent(results, "", "\t")

	if err != nil {
		return err
	}

	fmt.Println(string(j))

	return nil
}
