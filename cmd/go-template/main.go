package main

import (
	"fmt"
	"os"

	"github.com/grantbirki/go-template/internal/cmd"
)

func main() {
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
