package logoru

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/fatih/color"
	"github.com/stretchr/testify/assert"
)

func TestOutput(t *testing.T) {
	debug_value := output("DEBUG", color.New(color.FgGreen, color.Bold), "Testing testing")
	assert.True(
		t, strings.HasSuffix(
			debug_value,
			"| DEBUG | Testing testing",
		),
		fmt.Sprintf("is value of %#v", debug_value),
	)

}

func TestGenTime(t *testing.T) {
	instance := genTime()
	year := strconv.Itoa(time.Now().Year()) + "-"
	if len(strings.Split(instance, " ")) != 2 && strings.Contains(instance, year) {
		t.Error("Wrong format", instance)
	}
}

func TestAddZero(t *testing.T) {
	under10Instance := addZero(9)
	assert.Equal(t, under10Instance, "09")
	over10Instance := addZero(10)
	assert.Equal(t, over10Instance, "10")
}
