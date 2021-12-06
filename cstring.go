package gocom

import (
	"bytes"
	"strconv"
	"strings"
)

func StrToBytes(str string) []byte {
	return []byte(str)
}

func Normalize(str string) string {
	trimmedString := strings.ReplaceAll(str, " ", "")
	lowerString := strings.ToLower(trimmedString)
	return lowerString
}

func StringInSlice(str string, list []string) bool {
	for _, item := range list {
		if item == str {
			return true
		}
	}
	return false
}

func StripAll(str string, charactersToStrip []string) string {
	sanitizedString := str
	for _, char := range charactersToStrip {
		sanitizedString = strings.ReplaceAll(str, char, "")
	}
	return sanitizedString
}

func StrAt(slice []string, str string) int {
	for pos, s := range slice {
		if s == str {
			return pos
		}
	}
	return -1
}

func IntArrayToString(A []int, delim string) string {

	var buffer bytes.Buffer
	for i := 0; i < len(A); i++ {
		buffer.WriteString(strconv.Itoa(A[i]))
		if i != len(A)-1 {
			buffer.WriteString(delim)
		}
	}

	return buffer.String()
}
