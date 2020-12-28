package filesystem_test

import (
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"pb-dropbox-downloader/infrastructure/filesystem"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLocal_GetFilesInFolder(t *testing.T) {
	fs := filesystem.Local{}
	files := fs.GetFilesInFolder("../../mocks/test_directory")

	assert.EqualValues(t, []string{
		"book1.epub",
		"book2.epub",
		filepath.Join("test", "book3.epub"),
		filepath.Join("test", "book4.epub"),
	}, files)
}

func TestLocal_CopyDataToFile(t *testing.T) {
	tmpfile, err := ioutil.TempFile("", "example")
	if err != nil {
		log.Fatal(err)
	}

	defer os.Remove(tmpfile.Name())

	fs := filesystem.Local{}
	reader := strings.NewReader("This is test content")
	err = fs.CopyDataToFile(tmpfile.Name(), reader)

	assert.NoError(t, err)
	assert.FileExists(t, tmpfile.Name())
	data, _ := ioutil.ReadFile(tmpfile.Name())
	assert.Equal(t, "This is test content", string(data))
}

func TestLocal_DeleteFile(t *testing.T) {
	tmpfile, err := ioutil.TempFile("", "example")
	if err != nil {
		t.Error(err)
	}

	filePath := tmpfile.Name()

	_, err = tmpfile.WriteString("This is test content")
	if err != nil {
		t.Error(err)
	}

	tmpfile.Close()

	defer os.Remove(filePath)

	fs := filesystem.Local{}
	err = fs.DeleteFile(filePath)

	_, statError := os.Stat(filePath)

	assert.NoError(t, err)
	assert.True(t, os.IsNotExist(statError))
}

func TestLocal_DeleteFile_NotExist(t *testing.T) {
	fs := filesystem.Local{}
	err := fs.DeleteFile(path.Join(os.TempDir(), "not-exist-file"))

	assert.True(t, os.IsNotExist(err))
}
