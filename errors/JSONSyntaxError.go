package errors

import (
	"encoding/json"
	"fmt"
	"strings"
)

// JSONSyntaxError wraps json.SyntaxError
type JSONSyntaxError struct {
	Cause *json.SyntaxError
}

func (e *JSONSyntaxError) Error() string {
	original := e.Cause.Error()
	if strings.Contains(strings.ToLower(original), "json") {
		return original
	}
	return fmt.Sprintf("JSON syntax error: %s", original)
}

func (e *JSONSyntaxError) Unwrap() error {
	return e.Cause
}
