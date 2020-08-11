package logoru

import (
	"log"

	"github.com/fatih/color"
)

var levels = []string{"DEBUG", "INFO", "SUCCESS", "WARNING", "ERROR", "CRITICAL"}

// Output a debugging message
func Debug(msg interface{}) {
	err := output(levels[0], color.New(color.FgBlue, color.Bold), msg)
	if err != nil {
		log.Fatal(err)
	}
}

// Output an info message
func Info(msg interface{}) {
	err := output(levels[1], color.New(color.FgWhite, color.Bold), msg)
	if err != nil {
		log.Fatal(err)
	}
}

// Output a success message
func Success(msg interface{}) {
	err := output(levels[2], color.New(color.FgGreen, color.Bold), msg)
	if err != nil {
		log.Fatal(err)
	}
}

// Output a warning message
func Warning(msg interface{}) {
	err := output(levels[3], color.New(color.FgYellow, color.Bold), msg)
	if err != nil {
		log.Fatal(err)
	}
}

// Output an error message
func Error(msg interface{}) {
	err := output(levels[4], color.New(color.FgRed, color.Bold), msg)
	if err != nil {
		log.Fatal(err)
	}
}

// Output a critical message
func Critical(msg interface{}) {
	err := output(levels[5], color.New(color.FgWhite, color.BgRed, color.Bold), msg)
	if err != nil {
		log.Fatal(err)
	}
}
