package utils_test

import (
	"bytes"
	"errors"
	"pb-dropbox-downloader/internal/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPanicInterceptor(t *testing.T) {
	tests := []struct {
		name            string
		err             error
		expectedMesasge string
		exitCode        int
	}{
		{
			name:            "Intercepts panic and return with exit code 3",
			err:             errors.New("Test error"),
			expectedMesasge: "Critical error: Test error\n",
			exitCode:        3,
		},
		{
			name:            "Intercepts panic and return with exit code 0",
			err:             errors.New("Other error"),
			expectedMesasge: "Critical error: Other error\n",
			exitCode:        0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			called := false

			buffer := bytes.NewBufferString("")

			assert.NotPanics(t, func() {
				defer utils.PanicInterceptor(func(code int) {
					assert.Equal(t, tt.exitCode, code)
					called = true
				}, buffer, tt.exitCode)
				panic(tt.err)
			})

			assert.True(t, called)
			assert.Equal(t, tt.expectedMesasge, buffer.String())
		})
	}
}
