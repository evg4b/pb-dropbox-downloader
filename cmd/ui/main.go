package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"pb-dropbox-downloader/internal/app"
	"pb-dropbox-downloader/internal/utils"

	ink "github.com/dennwc/inkview"
)

const fatalExidCode = 500

func main() {
	ink.DefaultFontHeight = 18
	ink.RunCLI(func(ctx context.Context, w io.Writer) error {
		defer utils.PanicInterceptor(os.Exit, w, fatalExidCode)

		_, err := ink.KeepNetwork()
		if err != nil {
			fmt.Fprintf(w, "Error: %v", err)
			return nil
		}

		return app.Run(ctx, w)
	}, nil)
}
