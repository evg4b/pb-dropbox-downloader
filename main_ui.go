//go:build UI
// +build UI

package main

import (
	"context"
	"fmt"
	"io"

	ink "github.com/dennwc/inkview"
)

func main() {
	ink.DefaultFontHeight = 20
	ink.RunCLI(func(ctx context.Context, w io.Writer) error {
		_, err := ink.KeepNetwork()
		if err != nil {
			fmt.Fprintf(w, "Error: %v", err)
			return nil
		}

		mainInternal(w)

		return nil
	}, nil)
}
