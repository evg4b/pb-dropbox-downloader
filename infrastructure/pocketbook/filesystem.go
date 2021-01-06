package pocketbook

import (
	"path/filepath"
	"pb-dropbox-downloader/utils"
)

// IntrenalStoragePath path to the mounted internal storage folder
var IntrenalStoragePath = "/mnt/ext1"

// SdCardStoragePath path to the mounted sd card folder
var SdCardStoragePath = "/mnt/ext2"

// ConfigPath returns normalized path resolved from system/config folder in intranal storage
func ConfigPath(pathPartials ...string) string {
	return utils.JoinPath(IntrenalStoragePath, "system", "config", filepath.Join(pathPartials...))
}

// Application returns normalized path resolved from application config folder in intranal storage
func Application(pathPartials ...string) string {
	return utils.JoinPath(IntrenalStoragePath, "application", filepath.Join(pathPartials...))
}

// Share returns normalized path resolved from system/share foloder in internal storage
func Share(pathPartials ...string) string {
	return utils.JoinPath(IntrenalStoragePath, "system", "share", filepath.Join(pathPartials...))
}

// SdCard returns normalized path resolved from SD card storage
func SdCard(pathPartials ...string) string {
	return utils.JoinPath(SdCardStoragePath, filepath.Join(pathPartials...))
}

// Internal returns normalized path resolved from intranal storage
func Internal(pathPartials ...string) string {
	return utils.JoinPath(IntrenalStoragePath, filepath.Join(pathPartials...))
}
