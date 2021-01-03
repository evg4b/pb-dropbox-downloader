package utils

import (
	"path/filepath"
	"strings"
)

// JoinPath return joined and normalized path (all '\\' will be replaced '/')
func JoinPath(elem ...string) string {
	return strings.ReplaceAll(filepath.Join(elem...), "\\", "/")
}
