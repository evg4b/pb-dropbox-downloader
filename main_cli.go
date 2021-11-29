//go:build !UI
// +build !UI

package main

import "os"

func main() {
	mainInternal(os.Stdout)
}
