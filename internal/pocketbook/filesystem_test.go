package pocketbook_test

import (
	"pb-dropbox-downloader/internal/pocketbook"
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
		{
			name:         "resolve folder path",
			pathPartials: []string{},
			expected:     "/mnt/ext1/system/config",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := pocketbook.ConfigPath(tt.pathPartials...)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestApplication(t *testing.T) {
	tests := []struct {
		name         string
		pathPartials []string
		expected     string
	}{
		{
			name:         "resolve file",
			pathPartials: []string{"pb.app"},
			expected:     "/mnt/ext1/application/pb.app",
		},
		{
			name:         "resolve folder with array",
			pathPartials: []string{"demo", "pb.app"},
			expected:     "/mnt/ext1/application/demo/pb.app",
		},
		{
			name:         "resolve folder and file from string",
			pathPartials: []string{"demo/pb.app"},
			expected:     "/mnt/ext1/application/demo/pb.app",
		},
		{
			name:         "resolve folder and file from string with backslash",
			pathPartials: []string{"demo\\pb.app"},
			expected:     "/mnt/ext1/application/demo/pb.app",
		},
		{
			name:         "resolve folder path",
			pathPartials: []string{},
			expected:     "/mnt/ext1/application",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := pocketbook.Application(tt.pathPartials...)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestShare(t *testing.T) {
	tests := []struct {
		name         string
		pathPartials []string
		expected     string
	}{
		{
			name:         "resolve file",
			pathPartials: []string{"data.bin"},
			expected:     "/mnt/ext1/system/share/data.bin",
		},
		{
			name:         "resolve folder with array",
			pathPartials: []string{"demo", "data.bin"},
			expected:     "/mnt/ext1/system/share/demo/data.bin",
		},
		{
			name:         "resolve folder and file from string",
			pathPartials: []string{"demo/data.bin"},
			expected:     "/mnt/ext1/system/share/demo/data.bin",
		},
		{
			name:         "resolve folder and file from string with backslash",
			pathPartials: []string{"demo\\data.bin"},
			expected:     "/mnt/ext1/system/share/demo/data.bin",
		},
		{
			name:         "resolve folder path",
			pathPartials: []string{},
			expected:     "/mnt/ext1/system/share",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := pocketbook.Share(tt.pathPartials...)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestSdCard(t *testing.T) {
	tests := []struct {
		name         string
		pathPartials []string
		expected     string
	}{
		{
			name:         "resolve file",
			pathPartials: []string{"book.epub"},
			expected:     "/mnt/ext2/book.epub",
		},
		{
			name:         "resolve folder with array",
			pathPartials: []string{"demo", "book.epub"},
			expected:     "/mnt/ext2/demo/book.epub",
		},
		{
			name:         "resolve folder and file from string",
			pathPartials: []string{"demo/book.epub"},
			expected:     "/mnt/ext2/demo/book.epub",
		},
		{
			name:         "resolve folder and file from string with backslash",
			pathPartials: []string{"demo\\book.epub"},
			expected:     "/mnt/ext2/demo/book.epub",
		},
		{
			name:         "resolve folder path",
			pathPartials: []string{},
			expected:     "/mnt/ext2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := pocketbook.SdCard(tt.pathPartials...)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestInternal(t *testing.T) {
	tests := []struct {
		name         string
		pathPartials []string
		expected     string
	}{
		{
			name:         "resolve file",
			pathPartials: []string{"book.epub"},
			expected:     "/mnt/ext1/book.epub",
		},
		{
			name:         "resolve folder with array",
			pathPartials: []string{"demo", "book.epub"},
			expected:     "/mnt/ext1/demo/book.epub",
		},
		{
			name:         "resolve folder and file from string",
			pathPartials: []string{"demo/book.epub"},
			expected:     "/mnt/ext1/demo/book.epub",
		},
		{
			name:         "resolve folder and file from string with backslash",
			pathPartials: []string{"demo\\book.epub"},
			expected:     "/mnt/ext1/demo/book.epub",
		},
		{
			name:         "resolve folder path",
			pathPartials: []string{},
			expected:     "/mnt/ext1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := pocketbook.Internal(tt.pathPartials...)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
