package errors

import (
"fmt"
	"github.com/go-playground/validator/v10"
)

// GPValidatorError wraps go-playground's validator.ValidationErrors
type GPValidatorError struct {
	Cause validator.ValidationErrors
}

func (e *GPValidatorError) Error() string {
	// only takes the first error
	switch e.Cause[0].Tag() {
	case "required":
		// TODO get json tag not struct field name
		return fmt.Sprintf(`"%s" is a required field`, e.Cause[0].Field())
	default:
		return e.Cause.Error()
	}
}

func (e *GPValidatorError) Unwrap() error {
	return e.Cause
}
