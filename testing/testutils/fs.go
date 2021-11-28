package testutils

import (
	"os"
	"testing"

	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-billy/v5/util"
)

// FsFromMap creates billy.Filesystem in memory from map.
// Where key is a filename and value is file context.
func FsFromMap(t *testing.T, files map[string]string) billy.Filesystem {
	t.Helper()

	fs := memfs.New()
	for path, content := range files {
		err := util.WriteFile(fs, path, []byte(content), os.ModePerm)
		if err != nil {
			t.Error(err)
		}
	}

	return fs
}

// FsFromMap creates billy.Filesystem in memory from slice.
// Each element of slice is filename, content always "test".
func FsFromSlice(t *testing.T, files []string) billy.Filesystem {
	t.Helper()

	fs := memfs.New()
	for _, path := range files {
		err := util.WriteFile(fs, path, []byte("test"), os.ModePerm)
		if err != nil {
			t.Error(err)
		}
	}

	return fs
}

// MakeFiles creates in billy.Filesystem files from map.
// Where key is a filename and value is file context.
func MakeFiles(t *testing.T, fs billy.Basic, files map[string]string) {
	t.Helper()

	for filemane, content := range files {
		err := util.WriteFile(fs, filemane, []byte(content), os.ModePerm)
		if err != nil {
			t.Error(err)
		}
	}
}
