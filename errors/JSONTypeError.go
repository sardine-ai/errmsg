package errors

import (
	"encoding/json"
	"fmt"
)

// JSONTypeError wraps json.UnmarshalTypeError
type JSONTypeError struct {
	Cause *json.UnmarshalTypeError
}

func (e *JSONTypeError) Error() string {
	typeErr := e.Cause
	if typeErr.Field != "" {
		return fmt.Sprintf(`"%v" should be %s but received %s`, typeErr.Field, typeErr.Type.Name(), typeErr.Value)
	}
	return fmt.Sprintf("expected %s but received %s", typeErr.Type.Name(), typeErr.Value)
}

func (e *JSONTypeError) Unwrap() error {
	return e.Cause
}