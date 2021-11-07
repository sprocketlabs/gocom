
package gocom

import (
    "testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestStripWhitespaceFromString(t *testing.T) {
    str := " hello world "
    expectedString := "helloworld"
    normalString := normalize(str)
    match := expectedString == normalString
	if !match {
        t.Fatalf(`[Assertion Error] Expected: (%q) || Result: (%q)`, expectedString, normalString)
    }
}

func TestLowercaseString(t *testing.T) {
    str := "ToDaY"
    expectedString := "today"
    normalString := normalize(str)
    match := expectedString == normalString
	if !match {
        t.Fatalf(`[Assertion Error] Expected: (%q) || Result: (%q) ")`, expectedString, normalString)
    }
}
