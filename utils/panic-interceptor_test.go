package utils_test

import (
	"pb-dropbox-downloader/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPanicInterceptor(t *testing.T) {
	tests := []struct {
		name     string
		exitCode int
	}{
		{name: "Intercepts panic and return with exit code 3", exitCode: 3},
		{name: "Intercepts panic and return with exit code 0", exitCode: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			called := false

			assert.NotPanics(t, func() {
				defer utils.PanicInterceptor(func(code int) {
					assert.Equal(t, tt.exitCode, code)
					called = true
				}, tt.exitCode)
				panic("test panic")
			})

			assert.True(t, called)
		})
	}
}
