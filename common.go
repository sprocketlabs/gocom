package gocom
import (
	"strings"
)
func normalize(str string) (string) {
	trimmedString := strings.ReplaceAll(str, " ", "")
	lowerString := strings.ToLower(trimmedString)
	return lowerString
}
