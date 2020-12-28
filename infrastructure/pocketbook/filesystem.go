package pocketbook

import (
	"path"
)

const intrenalStorage = "/mnt/ext1"
const sdCardStorage = "/mnt/ext2"

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
