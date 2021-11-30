package synchroniser

import (
	"os"
	"strings"
)

// FileSliceContins checks if slice contains passed value.
func FileSliceContins(collection []os.FileInfo, value string) bool {
	for _, item := range collection {
		if strings.EqualFold(item.Name(), value) {
			return true
		}
	}

	return false
}

// FilterMapBy filters map of strings by predicate.
func FilterMapBy(data map[string]string, predicate func(string, string) bool) map[string]string {
	result := make(map[string]string)
	for key, value := range data {
		if predicate(key, value) {
			result[key] = value
		}
	}

	return result
}
