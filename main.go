package main

import (
	"fmt"
	"os"

	"github.com/Matt-Gleich/go_loguru/loguru"
)

func main() {
	err := loguru.Debug("Hello World!")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
