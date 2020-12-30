package utils

import (
	"log"
)

func PanicInterceptor(exit func(code int), exitCode int) {
	if err := recover(); err != nil {
		log.Printf("Fatal error: %s", err)
		exit(exitCode)
	}
}
