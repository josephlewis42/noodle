package main

import (
	"fmt"
	"os"

	"github.com/josephlewis42/noodle/internal/cmd"
)

func main() {
	if err := cmd.Noodle().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
