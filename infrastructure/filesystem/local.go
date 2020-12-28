package filesystem

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

const perm = 0755

type Local struct {
}

// GetFilesInFolder return relative file paths in folder (include subfolder to)
func (*Local) GetFilesInFolder(folder string) []string {
	files := []string{}
	err := filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		relativePath, err := filepath.Rel(folder, path)
		if err != nil {
			return err
		}

		files = append(files, relativePath)

		return nil
	})

	if err != nil {
		panic(err)
	}

	return files
}

// CopyDataToFile copy data from reader to file
func (*Local) CopyDataToFile(filePath string, source io.Reader) error {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, perm)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = io.Copy(file, source)

	return err
}

// DeleteFile delete file from target filesystem
func (*Local) DeleteFile(file string) error {
	return os.Remove(file)
}

// ReadFile read file content from target filesystem
func (*Local) ReadFile(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}

// WriteFile writes content to file in target filesystem
func (*Local) WriteFile(filename string, data []byte) error {
	return ioutil.WriteFile(filename, data, perm)
}
