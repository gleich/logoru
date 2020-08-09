package loguru

import (
	"fmt"
	"strconv"
	"time"

	"github.com/fatih/color"
)

// Core output function for loguru
func output(level string, c *color.Attribute, message string) {
	green := color.New(color.FgGreen)
	boldColor := color.New(*c, color.Bold)
	green.Print(genTime())
	fmt.Print(" | ")
	boldColor.Print(level)
	fmt.Print(" | ")
	boldColor.Println(message)
}

// Generate the time output for loguru
func genTime() string {
	now := time.Now()
	year, month, day := now.Date()
	return fmt.Sprintf("%v-%v-%v %v:%v:%v.%v", year, addZero(int(month)), addZero(day), addZero(now.Hour()), addZero(now.Minute()), addZero(now.Second()), addZero(now.Nanosecond())[:3])
}

// Add a zero before a number if it is under 10
func addZero(num int) string {
	strNum := strconv.Itoa(num)
	if num < 10 {
		return "0" + strNum
	}
	return strNum
}
