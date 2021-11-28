package utils_test

import (
	"pb-dropbox-downloader/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterSliceBy(t *testing.T) {
	tests := []struct {
		name       string
		collection []string
		predicate  func(string) bool
		expected   []string
	}{
		{
			name:       "empty slice",
			collection: []string{},
			predicate:  func(s string) bool { return true },
			expected:   []string{},
		},
		{
			name:       "predicate return always true",
			collection: []string{"1", "2", "3", "4"},
			predicate:  func(s string) bool { return true },
			expected:   []string{"1", "2", "3", "4"},
		},
		{
			name:       "predicate return always false",
			collection: []string{"1", "2", "3", "4"},
			predicate:  func(s string) bool { return false },
			expected:   []string{},
		},
		{
			name:       "predicate filter '2'",
			collection: []string{"1", "2", "3", "4"},
			predicate:  func(s string) bool { return s == "2" },
			expected:   []string{"2"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// filtered := utils.FilterSliceBy(tt.collection, tt.predicate)

			// assert.EqualValues(t, tt.expected, filtered)
		})
	}
}

func TestSliceContins(t *testing.T) {
	tests := []struct {
		name       string
		collection []string
		value      string
		expected   bool
	}{
		{
			name:       "empty slice",
			collection: []string{},
			value:      "11",
			expected:   false,
		},
		{
			name:       "contains value",
			collection: []string{"1", "2", "3", "4"},
			value:      "3",
			expected:   true,
		},
		{
			name:       "not contains value",
			collection: []string{"1", "2", "3", "4"},
			value:      "11",
			expected:   false,
		},
		{
			name:       "multiple contains",
			collection: []string{"1", "11", "11", "2", "3", "4"},
			value:      "11",
			expected:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isContains := utils.SliceContins(tt.collection, tt.value)

			assert.Equal(t, tt.expected, isContains)
		})
	}
}
