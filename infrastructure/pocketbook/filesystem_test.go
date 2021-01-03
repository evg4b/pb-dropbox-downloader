package pocketbook_test

import (
	"pb-dropbox-downloader/infrastructure/pocketbook"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigPath(t *testing.T) {
	tests := []struct {
		name         string
		pathPartials []string
		expected     string
	}{
		{
			name:         "resolve file",
			pathPartials: []string{"config.json"},
			expected:     "/mnt/ext1/system/config/config.json",
		},
		{
			name:         "resolve folder with array",
			pathPartials: []string{"demo", "config.json"},
			expected:     "/mnt/ext1/system/config/demo/config.json",
		},
		{
			name:         "resolve folder and file from string",
			pathPartials: []string{"demo/config.json"},
			expected:     "/mnt/ext1/system/config/demo/config.json",
		},
		{
			name:         "resolve folder and file from string with backslash",
			pathPartials: []string{"demo\\config.json"},
			expected:     "/mnt/ext1/system/config/demo/config.json",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := pocketbook.ConfigPath(tt.pathPartials...)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
