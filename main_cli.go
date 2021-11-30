//go:build !UI
// +build !UI

package main

import (
	"os"
	"pb-dropbox-downloader/utils"
)

func main() {
	defer utils.PanicInterceptor(os.Exit, fatalExitCode)

	err := mainInternal(os.Stdout)
	if err != nil {
		panic(err)
	}
}
