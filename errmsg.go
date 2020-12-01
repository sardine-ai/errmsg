package errmsg

import (
	"encoding/json"
	"github.com/daisy1754/errmsg/errors"
	"github.com/go-playground/validator/v10"
)

// Wrap tries to wrap known json syntax/parse/validation error into custom error type that provides better error message
func Wrap(err error) error {
	// encoding/json
	if typeErr, ok := err.(*json.UnmarshalTypeError); ok {
		return &errors.JSONTypeError{Cause: typeErr}
	}
	if syntaxErr, ok := err.(*json.SyntaxError); ok {
		return &errors.JSONSyntaxError{Cause: syntaxErr}
	}

	// go-playground/validator
	if validationErrs, ok := err.(validator.ValidationErrors); ok && len(validationErrs) > 0 {
		return &errors.GPValidatorError{Cause: validationErrs}
	}
	return err
}

// Message returns error massage for given err. This method tries to return friendlier message that makes more sense
// in the context of JSON APIs.
func Message(err error) string {
	return Wrap(err).Error()
}
