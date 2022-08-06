package main

import (
	"context"
	"os"
	"pb-dropbox-downloader/internal/app"
	"pb-dropbox-downloader/internal/utils"
)

const fatalExidCode = 500

func main() {
	defer utils.PanicInterceptor(os.Exit, os.Stdout, fatalExidCode)
	ctx := context.Background()
	if err := app.Run(ctx, os.Stdout); err != nil {
		panic(err)
	}
}
