package loguru

import (
	"fmt"
	"strconv"
	"time"

	"github.com/fatih/color"
)

func Output(level string, c *color.Color, message string) {
	green := color.New(color.FgGreen)
	green.Print(genTime())
	fmt.Print(" | ")
	c.Print(level)
	fmt.Print(" | ")
	c.Println(message)
}

func genTime() string {
	now := time.Now()
	year, month, day := now.Date()
	return fmt.Sprintf("%v-%v-%v %v:%v:%v.%v", year, addZero(int(month)), addZero(day), addZero(now.Hour()), addZero(now.Minute()), addZero(now.Second()), addZero(now.Nanosecond())[:3])
}

func addZero(num int) string {
	strNum := strconv.Itoa(num)
	if num < 10 {
		return "0" + strNum
	}
	return strNum
}
