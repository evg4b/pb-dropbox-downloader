package pocketbook

import (
	"os"
	"path"
	"path/filepath"
)

const intrenalStorage = "/mnt/ext1"
const sdCardStorage = "/mnt/ext2"

// GetFiles return relative file paths in folder (include subfolder to)
func GetFiles(folder string) []string {
	files := []string{}
	err := filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
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

// ConfigPath return normalized path resolved from application config folder
func ConfigPath(pathPartials ...string) string {
	return normalize(path.Join(intrenalStorage, "system", "config", path.Join(pathPartials...)))
}

// // Config return normalized to application config folder
// func Share(pathPartials ...string) string {
// 	return path.Join(filepath.Clean(path.Join(pathPartials...)))
// }

func normalize(fullpath string) string {
	return path.Clean(fullpath)
}
