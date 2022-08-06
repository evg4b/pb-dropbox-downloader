package utils

import (
	"fmt"
	"io"
	"log"
)

// PanicInterceptor intercept panic, log it, and exit with exit code.
func PanicInterceptor(exit func(code int), w io.Writer, exitCode int) {
	if err := recover(); err != nil {
		fmt.Fprintf(w, "Critical error: %s\n", err)
		log.Printf("Fatal error: %s", err)

		exit(exitCode)
	}
}
