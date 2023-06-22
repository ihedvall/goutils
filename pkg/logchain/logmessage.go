package logchain

import (
	"fmt"
	"path/filepath"
	"runtime"
	"time"
)

type Severity int

const (
	Emergency Severity = iota
	Alert
	Critical
	Error
	Warning
	Notice
	Info
	Debug
	Trace
)

type LogMessage struct {
	Severity  Severity
	Timestamp time.Time
	Message   string
	MessageID string
	File      string
	Line      int
	Values    map[string]string
}

func NewLogMessage(severity Severity, message string) *LogMessage {
	msg := LogMessage{Severity: severity, Timestamp: time.Now(), Message: message}
	_, file, line, callOK := runtime.Caller(1)
	if callOK && file != "" {
		msg.File = filepath.Base(file)
		msg.Line = line
	}

	return &msg
}

func (m *LogMessage) LocalTimeWithMs() string {
	t := m.Timestamp
	return fmt.Sprintf("%s.%03d", t.Format(time.DateTime), t.UnixMilli()%1000)
}

func (m *LogMessage) SeverityAsString() string {
	switch m.Severity {
	case Emergency:
		return "Emergency"
	case Alert:
		return "Alert"
	case Critical:
		return "Critical"
	case Error:
		return "Error"
	case Warning:
		return "Warning"
	case Notice:
		return "Notice"
	case Info:
		return "Info"
	case Debug:
		return "Debug"
		// default is Trace
	}
	return "Trace"
}
