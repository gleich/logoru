package logoru

import "testing"

const line = "Testing testing!"

var words = []interface{}{"Hello", "World!"}

func TestDebug(t *testing.T) {
	Debug(line)
	Debug(words...)
}

func TestInfo(t *testing.T) {
	Info(line)
	Info(words...)
}

func TestSuccess(t *testing.T) {
	Success(line)
	Success(words...)
}

func TestWarning(t *testing.T) {
	Warning(line)
	Warning(words...)
}

func TestError(t *testing.T) {
	Error(line)
	Error(words...)
}

func TestCritical(t *testing.T) {
	defer func() { recover() }()
	Critical(line)
	t.Error("Critical didn't panic")
}
