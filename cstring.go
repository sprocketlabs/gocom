package gocom

import (
	"strings"
)

func Normalize(str string) (string) {
	trimmedString := strings.ReplaceAll(str, " ", "")
	lowerString := strings.ToLower(trimmedString)
	return lowerString
}

func StrToBytes(str string) ([] byte) {
	return []byte(str)
}
