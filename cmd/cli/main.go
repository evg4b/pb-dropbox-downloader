package main

import (
	"os"
	"pb-dropbox-downloader/internal/app"
	"pb-dropbox-downloader/internal/utils"
)

func main() {
	defer utils.PanicInterceptor(os.Exit, os.Stdout, 500)
	if err := app.Run(os.Stdout); err != nil {
		panic(err)
	}
}
