package gocom

import (
	"testing"
)

/*
import (
	"reflect"
	"strings"
	"testing"
)

func TestGetKeysFromMap(t *testing.T) {
	m := map[string]int{
		"a": 0,
		"b": 1,
		"c": 2,
	}
	keys := GetMapKeys(m)
	expectedKeys := []string{"a", "b", "c"}
	if !(reflect.DeepEqual(keys, expectedKeys)) {
		t.Fatalf(`[Assertion Error] Lists do not match: (%q) != (%q) ")`, strings.Join(keys, " "), strings.Join(expectedKeys, " "))
	}
}
*/

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestGetStructName(t *testing.T) {
	type item struct{}
	i := item{}
	structName := GetType(i)
	match := structName == "item"
	if !match {
		t.Fatalf(`[Assertion Error] Expecting: (%q) || Result: (%q)`, "item", i)
	}
}
