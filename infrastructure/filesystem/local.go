package filesystem

import (
	"io"
	"os"
	"path/filepath"
)

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

func (*Local) CopyDataToFile(filePath string, source io.Reader) error {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = io.Copy(file, source)

	return err
}

func (*Local) DeleteFile(file string) error {
	return os.Remove(file)
}
