package main

import (
	"os"
	"pb-dropbox-downloader/infrastructure/pocketbook"
)

// nolint: gochecknoinits
func init() {
	pocketbook.IntrenalStoragePath = "./testing/testdata/internal"
	pocketbook.SdCardStoragePath = "./testing/testdata/sdcard"

	const perm = 0775

	_ = os.MkdirAll(pocketbook.ConfigPath("/"), perm)
	_ = os.MkdirAll(pocketbook.Share("/"), perm)
	_ = os.MkdirAll(pocketbook.SdCard("/"), perm)
}
