package errmsg_test

import (
	"encoding/json"
	"github.com/daisy1754/errmsg"
	errmsg_testing "github.com/daisy1754/errmsg/testing"
	"testing"
)

type birthday struct {
	Month int8 `json:"month"`
	Day int8 `json:"day"`
}

type user struct {
	ID string `json:"id"`
	BirthDay birthday `json:"birthday"`
}

func TestJsonTypeError(t *testing.T) {
	var u user
	input := `{"id": 1}`
	err := json.Unmarshal([]byte(input), &u)
	errmsg_testing.AssertEquals(t, errmsg.Message(err), `"id" should be string but received number`)
}


func TestJsonTypeErrorForNestedField(t *testing.T) {
	var u user
	input := `{"birthday": {"month": ""}}`
	err := json.Unmarshal([]byte(input), &u)
	errmsg_testing.AssertEquals(t, errmsg.Message(err), `"birthday.month" should be int8 but received string`)
}

func TestJsonTypeErrorForNestedFieldOverflow(t *testing.T) {
	var u user
	input := `{"birthday": {"month": 4294967296}}`
	err := json.Unmarshal([]byte(input), &u)
	errmsg_testing.AssertEquals(t, errmsg.Message(err), `"birthday.month" should be int8 but received number 4294967296`)
}
