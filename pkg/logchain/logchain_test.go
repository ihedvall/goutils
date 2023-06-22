package logchain

import (
	"fmt"
	"testing"
	"time"
)

func TestGetInstance(t *testing.T) {
	var lc = GetLogChain()
	if lc == nil {
		t.Error("Instance is nil")
	}
	lc.Clear()
}

func TestNewLogger(t *testing.T) {
	var lc = GetLogChain()
	defLogger, _ := lc.NewLogger("Default", LogToConsole, "")
	defLogger.ShowLocation = true
	if len(lc.loggers) != 1 {
		t.Error("Invalid length of logger list")
	}
	lc.Clear()
}

func TestAddLogger(t *testing.T) {
	var lc = GetLogChain()
	if len(lc.loggers) == 0 {
		defLogger, _ := lc.NewLogger("Default", LogToConsole, "")
		defLogger.ShowLocation = true
	}
	for i := 0; i < 100; i++ {
		temp := fmt.Sprintf("Hello: %02d", i)
		msg := NewLogMessage(Info, temp)
		lc.AddLogMessage(msg)
	}
	for _, logger := range lc.loggers {
		count := logger.NofMessages()
		if count != 100 {
			t.Errorf("Nof messages is %d", count)
		}
	}
	t.Log("Starting")
	lc.Start()
	t.Log("Started")
	for timeout := 0; timeout < 1000 && !lc.IsEmpty(); timeout++ {
		time.Sleep(time.Millisecond)
		t.Log("Looping")
	}

	time.Sleep(time.Second)
	msg1 := NewLogMessage(Info, "Pelle before stop")
	lc.AddLogMessage(msg1)

	t.Log("Stopping")
	lc.Stop()
	time.Sleep(time.Second)
	t.Log("Stopped")
	lc.Clear()
}

func TestFileLogger(t *testing.T) {
	var lc = GetLogChain()

	// Test that GetLogger returns nil i logger doesn't exist
	xxxLogger := lc.GetLogger("XXX")
	if xxxLogger != nil {
		t.Error("XXX logger seems to exists")
	}

	fileLogger := lc.GetLogger("File")
	if fileLogger == nil {
		fileLogger, _ = lc.NewLogger("File", LogToFile, "")
		fileLogger.ShowLocation = true
		t.Log("File location:", fileLogger.ConnectInfo)
	}
	for i := 0; i < 100; i++ {
		temp := fmt.Sprintf("File: %02d", i)
		msg := NewLogMessage(Info, temp)
		lc.AddLogMessage(msg)
	}
	showLogFile := lc.ShowLogFile()
	if !showLogFile {
		t.Error("Failed to open the log file")
	}
	lc.Start()
	for timeout := 0; timeout < 1000 && !lc.IsEmpty(); timeout++ {
		time.Sleep(time.Millisecond)
	}
	lc.Stop()
	lc.Clear()
}
