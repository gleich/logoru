package main

import (
	"github.com/Matt-Gleich/go_loguru/loguru"
	"github.com/fatih/color"
)

func main() {
	loguru.Output("DEBUG    ", color.New(color.FgBlue), "Hello!")
}
