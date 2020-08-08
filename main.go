package main

import (
	"github.com/fatih/color"
	"github.com/github_username/go_loguru/loguru"
)

func main() {
	loguru.Output("DEBUG    ", color.New(color.FgBlue), "Hello!")
}
