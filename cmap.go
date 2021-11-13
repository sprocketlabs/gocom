package gocom

func GetMapKeys(m map[string]interface{}) []string {
	keys := []string{}
	for key, _ := range m {
		keys = append(keys, key)
	}
	return keys
}
