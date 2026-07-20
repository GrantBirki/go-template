package main

import (
	"fmt"
	"os"

	"github.com/grantbirki/go-template/internal/cmd"
)

func main() {
	err := cmd.Run(os.Args[1:], os.Stdout, os.Stderr)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
