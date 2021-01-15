package logoru

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
)

// Core output function for loguru
func output(level string, levelColor *color.Color, message ...interface{}) error {
	log.SetFlags(0)

	var msg string
	msg = color.New(color.FgGreen).Sprint(genTime()) + " | "
	msg = msg + levelColor.Sprint(level)
	msg = msg + " | "
	for _, messageChunk := range message {
		msg = strings.TrimLeft(fmt.Sprintf(msg+" %v", messageChunk), " ")
	}
	log.Println(msg)
	return nil
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
