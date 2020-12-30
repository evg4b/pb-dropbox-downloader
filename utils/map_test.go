package utils_test

import (
	"pb-dropbox-downloader/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterMapBy(t *testing.T) {
	tests := []struct {
		name      string
		data      map[string]string
		predicate func(string, string) bool
		expected  map[string]string
	}{
		{
			name: "predicate return always true",
			data: map[string]string{
				"1": "11",
				"2": "22",
				"3": "33",
			},
			predicate: func(s1, s2 string) bool { return true },
			expected: map[string]string{
				"1": "11",
				"2": "22",
				"3": "33",
			},
		},
		{
			name: "predicate return always false",
			data: map[string]string{
				"1": "11",
				"2": "22",
				"3": "33",
			},
			predicate: func(s1, s2 string) bool { return false },
			expected:  map[string]string{},
		},
		{
			name: "predicate filter by key",
			data: map[string]string{
				"1": "11",
				"2": "22",
				"3": "33",
			},
			predicate: func(s1, s2 string) bool { return s1 != "1" },
			expected: map[string]string{
				"2": "22",
				"3": "33",
			},
		},
		{
			name: "predicate filter by value",
			data: map[string]string{
				"1": "11",
				"2": "22",
				"3": "33",
			},
			predicate: func(s1, s2 string) bool { return s2 == "22" },
			expected: map[string]string{
				"2": "22",
			},
		},
		{
			name:      "predicate filter by value",
			data:      map[string]string{},
			predicate: func(s1, s2 string) bool { return s2 == "22" },
			expected:  map[string]string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filtered := utils.FilterMapBy(tt.data, tt.predicate)

			assert.EqualValues(t, tt.expected, filtered)
		})
	}
}
