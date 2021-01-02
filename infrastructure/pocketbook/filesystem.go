package pocketbook

import (
	"path/filepath"
)

const intrenalStorage = "/mnt/ext1"
const sdCardStorage = "/mnt/ext2"

// ConfigPath returns normalized path resolved from system/config folder in intranal storage
func ConfigPath(pathPartials ...string) string {
	return filepath.FromSlash(filepath.Join(intrenalStorage, "system", "config", filepath.Join(pathPartials...)))
}

// Application returns normalized path resolved from application config folder in intranal storage
func Application(pathPartials ...string) string {
	return filepath.FromSlash(filepath.Join(intrenalStorage, "application", filepath.Join(pathPartials...)))
}

// Share returns normalized path resolved from system/share foloder in internal storage
func Share(pathPartials ...string) string {
	return filepath.FromSlash(filepath.Join(intrenalStorage, "system", "share", filepath.Join(pathPartials...)))
}

// SdCard returns normalized path resolved from SD card storage
func SdCard(pathPartials ...string) string {
	return filepath.FromSlash(filepath.Join(sdCardStorage, filepath.Join(pathPartials...)))
}

// Internal returns normalized path resolved from intranal storage
func Internal(pathPartials ...string) string {
	return filepath.FromSlash(filepath.Join(intrenalStorage, filepath.Join(pathPartials...)))
}
