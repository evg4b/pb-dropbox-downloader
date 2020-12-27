package utils

import "strings"

// FilterBy filters slice of strings  by predicate
func FilterBy(collection []string, predicate func(string) bool) []string {
	result := []string{}
	for _, item := range collection {
		if predicate(item) {
			result = append(result, item)
		}
	}

	return result
}

func Contins(collection []string, value string) bool {
	for _, item := range collection {
		if strings.EqualFold(item, value) {
			return true
		}
	}

	return false
}
