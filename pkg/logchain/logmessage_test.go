package logchain

import (
	"testing"
)

func TestLogMessage(test *testing.T) {
	msg := NewLogMessage(Info, "Olle was here")

	test.Log(msg)
}
