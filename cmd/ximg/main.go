package main

import (
	"fmt"
	"os"

	"github.com/Kshitij-2412/ximg/internal/cli"
)

func main() {
	app := cli.New()

	if err := app.Run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
