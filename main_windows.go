package main

import (
	"os"
	"pb-dropbox-downloader/infrastructure/pocketbook"
)

func init() {
	pocketbook.IntrenalStoragePath = "./testdata/internal"
	pocketbook.SdCardStoragePath = "./testdata/sdcard"

	os.MkdirAll(pocketbook.ConfigPath("/"), 0775)
	os.MkdirAll(pocketbook.Share("/"), 0775)
	os.MkdirAll(pocketbook.SdCard("/"), 0775)
}
