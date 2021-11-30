package utils_test

import (
	"pb-dropbox-downloader/internal/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJoinPath(t *testing.T) {
	tests := []struct {
		name     string
		elem     []string
		expected string
	}{
		{
			name:     "",
			elem:     []string{"/mnt", "ext1", "base.bin"},
			expected: "/mnt/ext1/base.bin",
		},
		{
			name:     "",
			elem:     []string{"/mnt/ext1/base.bin"},
			expected: "/mnt/ext1/base.bin",
		},
		{
			name:     "",
			elem:     []string{"/mnt/ext1", "/base.bin"},
			expected: "/mnt/ext1/base.bin",
		},
		{
			name:     "",
			elem:     []string{"\\mnt\\ext1", "/base.bin"},
			expected: "/mnt/ext1/base.bin",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			joined := utils.JoinPath(tt.elem...)

			assert.Equal(t, tt.expected, joined)
		})
	}
}
