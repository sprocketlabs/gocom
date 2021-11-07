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

func getMapKeys(m map[string]int) []string {
    keys := []string{}
    for key, _ := range m {
        keys = append(keys, key)
    }
    return keys
}