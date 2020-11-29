package errmsg

import (
	"encoding/json"
	"github.com/daisy1754/errmsg/errors"
)

// Wrap tries to wrap known json syntax/parse/validation error into custom error type that provides better error message
func Wrap(err error) error {
	if typeErr, ok := err.(*json.UnmarshalTypeError); ok {
		return &errors.JSONTypeError{Cause: typeErr}
	}
	return err
}

// Message returns error massage for given err. This method tries to return friendlier message that makes more sense
// in the context of JSON APIs.
func Message(err error) string {
	return Wrap(err).Error()
}
