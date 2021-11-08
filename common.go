package gocom
import (
	"strings"
)
func Normalize(str string) (string) {
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

func GetMapKeys(m map[string]int) []string {
    keys := []string{}
    for key, _ := range m {
        keys = append(keys, key )
}
    return keys
}