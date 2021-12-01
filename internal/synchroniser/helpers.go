package synchroniser

import (
	"os"
	"strings"
)

func fileSliceContins(collection []os.FileInfo, value string) bool {
	for _, item := range collection {
		if strings.EqualFold(item.Name(), value) {
			return true
		}
	}

	return false
}

func filterMapBy(data map[string]string, predicate func(string, string) bool) map[string]string {
	result := make(map[string]string)
	for key, value := range data {
		if predicate(key, value) {
			result[key] = value
		}
	}

	return result
}
