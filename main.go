package main

import (
	"fmt"

	"github.com/Matt-Gleich/go_loguru/loguru"
)

func main() {
	fmt.Println("")
	loguru.Debug("Here is a debug message")
	fmt.Println("")
	loguru.Info("Here is an info message")
	fmt.Println("")
	loguru.Success("Here is a success message")
	fmt.Println("")
	loguru.Warning("Here is a warning message")
	fmt.Println("")
	loguru.Error("Here is an error message")
	fmt.Println("")
	loguru.Critical("Here is a critical message")
	fmt.Println("")
}
