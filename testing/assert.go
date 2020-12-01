package testing

import (
	"reflect"
	"strings"
	"testing"
)

func AssertEquals(t *testing.T, exp, got interface{}) {
	if !reflect.DeepEqual(exp, got) {
		t.Fatalf("Expected '%v' but received '%v'", exp, got)
	}
}

func AssertHasPrefix(t *testing.T, data string, prefix string) {
	if !strings.HasPrefix(data, prefix) {
		t.Fatalf("Expected %s to start with %s", data, prefix)
	}
}
