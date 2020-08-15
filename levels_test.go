package logoru

import "testing"

const line1 = "Testing testing!"

func TestDebug(t *testing.T) {
	Debug(line1)
	Debug("Hello", "World!")
}

func TestInfo(t *testing.T) {
	Info(line1)
	Info("Hello", "World!")
}

func TestSuccess(t *testing.T) {
	Success(line1)
	Success("Hello", "World!")
}

func TestWarning(t *testing.T) {
	Warning(line1)
	Warning("Hello", "World!")
}

func TestError(t *testing.T) {
	Error(line1)
	Error("Hello", "World!")
}

func TestCritical(t *testing.T) {
	Critical(line1)
	Critical("Hello", "World!")
}
