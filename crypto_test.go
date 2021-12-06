package gocom

//import (
//    "testing"
//    "strings"
//    "reflect"
//)
import (
	"testing"
)

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
