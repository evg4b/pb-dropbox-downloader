package pocketbook

import (
	"path"
)

const intrenalStorage = "/mnt/ext1"
const sdCardStorage = "/mnt/ext2"

// ConfigPath returns normalized path resolved from system/config folder in intranal storage
func ConfigPath(pathPartials ...string) string {
	return normalize(path.Join(intrenalStorage, "system", "config", path.Join(pathPartials...)))
}

// Application returns normalized path resolved from application config folder in intranal storage
func Application(pathPartials ...string) string {
	return normalize(path.Join(intrenalStorage, "application", path.Join(pathPartials...)))
}

// Share returns normalized path resolved from system/share foloder in internal storage
func Share(pathPartials ...string) string {
	return normalize(path.Join(intrenalStorage, "system", "share", path.Join(pathPartials...)))
}

// SdCard returns normalized path resolved from SD card storage
func SdCard(pathPartials ...string) string {
	return normalize(path.Join(sdCardStorage, path.Join(pathPartials...)))
}

// Internal returns normalized path resolved from intranal storage
func Internal(pathPartials ...string) string {
	return normalize(path.Join(intrenalStorage, path.Join(pathPartials...)))
}

func normalize(fullpath string) string {
	return path.Clean(fullpath)
}
