
package gocom

import (
    "testing"
    "strings"
    "reflect"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestStripWhitespaceFromString(t *testing.T) {
    str := " hello world "
    expectedString := "helloworld"
    normalString := normalize(str)
    match := expectedString == normalString
	if !match {
        t.Fatalf(`[Assertion Error] Expecting: (%q) || Result: (%q)`, expectedString, normalString)
    }
}

func TestLowercaseString(t *testing.T) {
    str := "ToDaY"
    expectedString := "today"
    normalString := normalize(str)
    match := expectedString == normalString
	if !match {
        t.Fatalf(`[Assertion Error] Expecting: (%q) || Result: (%q) ")`, expectedString, normalString)
    }
}

func TestItemInList(t *testing.T) {
    str := "hello"
    stringList := []string {"hello", "world"}
    strExistsInList := stringInSlice(str, stringList)
	if !strExistsInList {
        t.Fatalf(`[Assertion Error] Expecting: (%q) in List (%q) ")`, str, strings.Join(stringList, " "))
    }
}

func TestItemNotInList(t *testing.T) {
    str := "hello"
    stringList := []string {"hi", "world"}
    strExistsInList := stringInSlice(str, stringList)
	if strExistsInList {
        t.Fatalf(`[Assertion Error] Not Expecting: (%q) in List (%q) ")`, str, strings.Join(stringList, " "))
    }
}

func TestGetKeysFromMap(t *testing.T) {
    m := map[string]int{
        "a": 0,
        "b": 1,
    }
    keys := getMapKeys(m)
    expectedKeys := []string {"a", "b"}
    if !(reflect.DeepEqual(keys, expectedKeys)) {
        t.Fatalf(`[Assertion Error] Lists do not match: (%q) != (%q) ")`, strings.Join(keys, " "), strings.Join(expectedKeys, " "))
    }
}
