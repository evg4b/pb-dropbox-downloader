package testutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// AssertError error by string, where expected string is empty then error should be null.
// Otherwise error message should be matched to string.
func AssertError(t *testing.T, expected string, actual error) {
	t.Helper()

	if len(expected) > 0 {
		assert.EqualError(t, actual, expected)
	} else {
		assert.NoError(t, actual)
	}
}
