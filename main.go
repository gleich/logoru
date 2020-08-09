package main

import "github.com/Matt-Gleich/go_loguru/loguru"

func main() {
	loguru.Debug("Hello World!")
	loguru.Info("Hello World!")
	loguru.Success("Hello World!")
	loguru.Warning("Hello World!")
	loguru.Error("Hello World!")
	loguru.Critical("Hello World!")
}
