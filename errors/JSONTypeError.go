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
	expectedType := typeErr.Type.Name()
	actualValue := typeErr.Value
	if isIntType(expectedType) && actualValue == "string" {
		// instead of saying 'int64' etc, let's just say 'integer'
		expectedType = "integer"
	}
	if typeErr.Field != "" {
		return fmt.Sprintf(`'%v' should be %s but received %s`, typeErr.Field, expectedType, actualValue)
	}
	return fmt.Sprintf("expected %s but received %s", expectedType, actualValue)
}

func (e *JSONTypeError) Unwrap() error {
	return e.Cause
}

func isIntType(typeName string) bool {
	return typeName == "int" || typeName == "unit" ||
		typeName == "int8" || typeName == "unit8" ||
		typeName == "int16" || typeName == "unit16"||
		typeName == "int32" || typeName == "unit32"||
		typeName == "int64" || typeName == "unit64"
}