package gocom
import (
	"strings"
)
func normalize(str string) (string) {
	trimmedString := strings.ReplaceAll(str, " ", "")
	lowerString := strings.ToLower(trimmedString)
	return lowerString
}

func stringInSlice(str string, list []string) bool {
    for _, item := range list {
        if item == str {
            return true
        }
    }
    return false
}
