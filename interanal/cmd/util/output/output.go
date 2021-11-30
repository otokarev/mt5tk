package output

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type output struct{}

func Print(format string, data interface{}) error {
	var o output
	if format == "" {
		return errors.New("empty format")
	}
	// Parse strings like "jsonpath={[*].Group}"
	formatParts := strings.SplitN(format, "=", 2)
	formatType := formatParts[0]
	formatOptV := reflect.ValueOf("")
	if len(formatParts) == 2 {
		formatOptV = reflect.ValueOf(formatParts[1])
	}

	meth := reflect.ValueOf(&o).MethodByName("Print" + strings.Title(formatType))
	if !meth.IsValid() {
		return errors.New(fmt.Sprintf("bad output format `%s`", format))
	}

	results := meth.Call([]reflect.Value{reflect.ValueOf(data), formatOptV})
	if err := results[0].Interface(); err != nil {
		return err.(error)
	}

	return nil
}
