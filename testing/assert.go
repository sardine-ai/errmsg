package testing

import (
	"reflect"
	"testing"
)

func AssertEquals(t *testing.T, exp, got interface{}) {
	if !reflect.DeepEqual(exp, got) {
		t.Fatalf("Expected '%v' but received '%v'", exp, got)
	}
}
