package main

import (
	"os"

  "github.com/one-click-platform/deployer/internal/cli"

)

func main() {
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}
