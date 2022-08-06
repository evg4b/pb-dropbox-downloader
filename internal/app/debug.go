//go:build debug
// +build debug

package app

import (
	"os"
	"pb-dropbox-downloader/internal/pocketbook"
)

func init() { // nolint: gochecknoinits
	pocketbook.IntrenalStoragePath = "./testing/testdata/internal"
	pocketbook.SdCardStoragePath = "./testing/testdata/sdcard"

	const perm = 0775

	_ = os.MkdirAll(pocketbook.ConfigPath(), perm)
	_ = os.MkdirAll(pocketbook.Share(), perm)
	_ = os.MkdirAll(pocketbook.SdCard(), perm)
}
