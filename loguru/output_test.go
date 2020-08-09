package loguru

import (
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/fatih/color"
	"github.com/stretchr/testify/assert"
)

func TestOutput(t *testing.T) {
	err := output("DEBUG", color.New(color.FgGreen, color.Bold), "Testing testing")
	if err != nil {
		t.Error(err)
	}
}

func TestAddSpacing(t *testing.T) {
	sameInstance, err := addSpacing([]string{"test", "test"}, "test")
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, sameInstance, "TEST")
	differentInstance, err := addSpacing([]string{"testtest", "test"}, "test")
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, differentInstance, "TEST    ")
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
