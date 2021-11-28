package utils

import (
	"os"
	"strings"
)

// FilterSliceBy filters slice of strings by predicate.
func FilterSliceBy(collection []os.FileInfo, predicate func(os.FileInfo) bool) []string {
	result := []string{}
	for _, item := range collection {
		if predicate(item) {
			result = append(result, item.Name())
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

// FileSliceContins checks if slice contains passed value.
func FileSliceContins(collection []os.FileInfo, value string) bool {
	for _, item := range collection {
		if strings.EqualFold(item.Name(), value) {
			return true
		}
	}

	return false
}
