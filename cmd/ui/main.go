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

func main() {
	ink.DefaultFontHeight = 18
	ink.RunCLI(func(ctx context.Context, w io.Writer) error {
		defer utils.PanicInterceptor(os.Exit, w, 500)

		_, err := ink.KeepNetwork()
		if err != nil {
			fmt.Fprintf(w, "Error: %v", err)
			return nil
		}

		return app.Run(w)
	}, nil)
}
