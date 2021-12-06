package gocom

import "reflect"

func GetMapKeys(m map[string]interface{}) []string {
	keys := []string{}
	for key, _ := range m {
		keys = append(keys, key)
	}
	return keys
}

func GetType(myvar interface{}) string {
	return reflect.TypeOf(myvar).Name()
}
