package filesystem

import (
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

const perm = 0755

type Local struct {
}

// GetFilesInFolder return relative file paths in folder (include subfolder to)
func (*Local) GetFilesInFolder(folder string) []string {
	if _, err := os.Stat(folder); os.IsNotExist(err) {
		return []string{}
	}

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
func (*Local) CopyDataToFile(filename string, source io.Reader) error {
	mkdirIfNotExistDir(filename)
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, perm)
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
	mkdirIfNotExistDir(filename)
	return ioutil.WriteFile(filename, data, perm)
}

func mkdirIfNotExistDir(filename string) {
	dir := path.Dir(filename)
	if len(dir) > 0 {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			os.Mkdir(dir, perm)
		}
	}
}
