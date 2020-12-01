package errmsg_test

import (
	"encoding/json"
	"errors"
	"github.com/daisy1754/errmsg"
	errmsg_testing "github.com/daisy1754/errmsg/testing"
	"github.com/go-playground/validator/v10"
	"testing"
)

type request struct {
	UserID string    `json:"userId" validate:"required"`
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
