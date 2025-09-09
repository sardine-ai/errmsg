package errmsg_test

import (
	"encoding/json"
	"errors"
	"testing"

	"bytes"

	"github.com/daisy1754/errmsg"
	errmsg_testing "github.com/daisy1754/errmsg/testing"
	"github.com/go-playground/validator/v10"
)

type request struct {
	UserID string   `json:"userId" validate:"required"`
	Names  []string `json:"names"`
}

func TestRequiredValidationError(t *testing.T) {
	validate := validator.New()
	var r request
	input := `{}`
	json.Unmarshal([]byte(input), &r)
	err := validate.Struct(r)
	errmsg_testing.AssertEquals(t, errmsg.Message(err), `'UserID' is a required field`)
	errmsg_testing.AssertType(t, errors.Unwrap(errmsg.Wrap(err)), validator.ValidationErrors{})
}

func TestInvalidNamesTypeError(t *testing.T) {
	var r request
	input := `{"userId": "123", "names": "John Doe"}`
	err := json.NewDecoder(bytes.NewReader([]byte(input))).Decode(&r)
	errmsg_testing.AssertEquals(t, errmsg.Message(err), `'names' should be []string but received string`)
	errmsg_testing.AssertType(t, errors.Unwrap(errmsg.Wrap(err)), &json.UnmarshalTypeError{})
}

func TestInvalidNamesValidationError(t *testing.T) {
	var r request
	input := `{"userId": "123", "names": 1}`
	err := json.NewDecoder(bytes.NewReader([]byte(input))).Decode(&r)
	errmsg_testing.AssertEquals(t, errmsg.Message(err), `'names' should be []string but received number`)
	errmsg_testing.AssertType(t, errors.Unwrap(errmsg.Wrap(err)), &json.UnmarshalTypeError{})
}
