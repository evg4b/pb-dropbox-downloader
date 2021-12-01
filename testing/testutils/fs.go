package testutils

import (
	"os"
	"pb-dropbox-downloader/testing/mocks"
	"testing"

	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-billy/v5/util"
	"github.com/stretchr/testify/assert"
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

func AssertFiles(t *testing.T, fs billy.Filesystem, folder string, files map[string]string) {
	t.Helper()

	if files == nil || fs == nil {
		return
	}

	if _, ok := fs.(*mocks.FilesystemMock); ok {
		return
	}

	existingFiles, err := fs.ReadDir(folder)
	assert.NoError(t, err)
	if len(files) == len(existingFiles) {
		for file, content := range files {
			fileContent, err := util.ReadFile(fs, file)
			assert.NoError(t, err)
			assert.Equal(t, content, string(fileContent))
		}
	} else {
		t.Error("files not matched")
	}
}
