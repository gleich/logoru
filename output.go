package logoru

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
)

// Core output function for loguru
func output(level string, levelColor *color.Color, message ...interface{}) error {
	green := color.New(color.FgGreen)
	green.Print(genTime())
	fmt.Print(" | ")
	fixedSpacing, err := addSpacing(levels, level)
	if err != nil {
		return err
	}
	levelColor.Print(fixedSpacing)
	fmt.Print(" | ")
	joinedMessage := ""
	for _, messageChunk := range message {
		joinedMessage = strings.TrimLeft(fmt.Sprintf(joinedMessage+" %v", messageChunk), " ")
	}
	levelColor.Print(joinedMessage)
	fmt.Println()
	return nil
}

// Add spacing to align the columns
func addSpacing(levels []string, level string) (string, error) {
	var found bool
	for _, tempLevel := range levels {
		if strings.EqualFold(tempLevel, level) {
			found = true
		}
	}
	if !found {
		return "", errors.New("Level " + level + " not found in levels")
	}

	var maxLen int
	for _, tempLevel := range levels {
		levelLen := len(tempLevel)
		if levelLen > maxLen {
			maxLen = levelLen
		}
	}

	fixedLevel := level
	for i := len(level); i < maxLen; i++ {
		fixedLevel = fixedLevel + " "
	}
	return strings.ToUpper(fixedLevel), nil
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
