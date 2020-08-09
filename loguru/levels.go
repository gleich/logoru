package loguru

import "github.com/fatih/color"

var levels = []string{"DEBUG", "SUCCESS", "ERROR", "INFO"}

// Output a debugging message
func Debug(msg string) error {
	c := color.FgGreen
	err := output(levels[0], &c, msg)
	return err
}
