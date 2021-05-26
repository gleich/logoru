package logoru

import (
	"log"

	"github.com/fatih/color"
)

var levels = []string{
	" DEBUG  ",
	" INFO   ",
	"SUCCESS ",
	"WARNING ",
	" ERROR  ",
	"CRITICAL",
}

// Output a debugging message
func Debug(msg ...interface{}) {
	out := output(levels[0], color.New(color.FgBlue, color.Bold), msg...)
	log.Println(out)
}

// Output an info message
func Info(msg ...interface{}) {
	out := output(levels[1], color.New(color.FgWhite, color.Bold), msg...)
	log.Println(out)
}

// Output a success message
func Success(msg ...interface{}) {
	out := output(levels[2], color.New(color.FgGreen, color.Bold), msg...)
	log.Println(out)
}

// Output a warning message
func Warning(msg ...interface{}) {
	out := output(levels[3], color.New(color.FgYellow, color.Bold), msg...)
	log.Println(out)

}

// Output an error message
func Error(msg ...interface{}) {
	out := output(levels[4], color.New(color.FgRed, color.Bold), msg...)
	log.Println(out)
}

// Output a critical message
func Critical(msg ...interface{}) {
	out := output(levels[5], color.New(color.FgWhite, color.BgRed, color.Bold), msg...)
	log.Panicln(out)
}
