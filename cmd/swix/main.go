package main

import (
	"os"

	"github.com/tulin404/SwiX/internal/cli"
)

// Cobra's entry point
func main() {
	if err := cli.Execute(); err != nil {
		os.Exit(1)
	}
}
