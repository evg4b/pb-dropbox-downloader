//go:build !UI
// +build !UI

package main

import (
	"os"
	"pb-dropbox-downloader/internal/utils"
)

func main() {
	defer utils.PanicInterceptor(os.Exit, fatalExitCode)
	if err := mainInternal(os.Stdout); err != nil {
		panic(err)
	}
}
