package utils

import "strings"

// FilterSliceBy filters slice of strings by predicate.
func FilterSliceBy(collection []string, predicate func(string) bool) []string {
	result := []string{}
	for _, item := range collection {
		if predicate(item) {
			result = append(result, item)
		}
	}

	return result
}

// SliceContins checks if slice contains passed value.
func SliceContins(collection []string, value string) bool {
	for _, item := range collection {
		if strings.EqualFold(item, value) {
			return true
		}
	}

	return false
}
