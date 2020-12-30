package utils

import (
	"log"
)

// PanicInterceptor intercept panic, log it, and exit with exit code
func PanicInterceptor(exit func(code int), exitCode int) {
	if err := recover(); err != nil {
		log.Printf("Fatal error: %s", err)
		exit(exitCode)
	}
}
