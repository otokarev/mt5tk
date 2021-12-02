package connection

import "fmt"

type Error struct {
	Retcode string
}

func NewError(retcode string) *Error {
	return &Error{Retcode: retcode}
}

func (e *Error) Error() string {
	return fmt.Sprintf("request failed, retcode: %s", e.Retcode)
}

func (e *Error) IsNotFound() bool {
	return e.Retcode == "13 Not found"
}
