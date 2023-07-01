package logchain

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
	"unicode/utf8"
)

type Facility int

const RFC339micro string = "2006-01-02T15:04:05.000000Z07:00"
const (
	Kernel Facility = iota
	UserLevel
	Mail
	SystemDaemon
	Security
	Internal
	Printer
	Network
	UUCP
	Clock
	Authorization
	FTP
	NTP
	LogAudit
	LogAlert
	ClockDaemon
	Local0
	Local1
	Local2
	Local3
	Local4
	Local5
	Local6
	Local7
)

type SyslogMessage struct {
	Index           int64    // Unique index of the message
	Severity        Severity // Same as log message severity but Trace doesn't exist in syslog
	Facility        Facility
	Version         uint8
	Timestamp       time.Time
	Hostname        string
	ApplicationName string
	ProcessID       string
	MessageID       string
	Message         string
	Data            []StructuredData
}

func NewSyslogMessage() *SyslogMessage {
	sm := SyslogMessage{
		Severity:  Info,
		Facility:  Local0,
		Version:   1,
		Timestamp: time.Now(),
		ProcessID: fmt.Sprintf("%d", os.Getpid()),
	}
	sm.Hostname, _ = os.Hostname()
	if len(os.Args) > 0 {
		filename := filepath.Base(os.Args[0])
		sm.ApplicationName = strings.TrimSuffix(filename, filepath.Ext(filename))
	}
	return &sm
}

func (sm *SyslogMessage) Populate(lm *LogMessage) {
	if lm == nil {
		return
	}

	// Note that the log level for event/syslog is normally set to Info i.e. Debug and Trace are not sent
	sm.Severity = lm.Severity
	if sm.Severity >= Trace {
		sm.Severity = Debug
	}
	sm.Timestamp = lm.Timestamp
	sm.Message = lm.Message
	sm.MessageID = lm.MessageID

	// Add default data defined in the log chain
	lc := GetLogChain()
	if lc == nil {
		return
	}
	lc.lockData.Lock()
	defer lc.lockData.Unlock()

	for _, sd := range lc.DefaultData {
		copyData := NewStructuredData(sd.Identity)

		for key, value := range sd.ParameterList {
			if key == "File" && lm.File != "" {
				copyData.ParameterList["File"] = lm.File
			} else if key == "Line" && lm.File != "" { // Cannot test on line number
				copyData.ParameterList["Line"] = fmt.Sprintf("%d", lm.Line)
			} else if value != "" {
				copyData.ParameterList[key] = value
			}
		}
		// Add if any parameters to send
		if len(copyData.ParameterList) > 0 {
			sm.Data = append(sm.Data, *copyData)
		}
	}

}

func (sm *SyslogMessage) MakeString() string {
	severity := sm.Severity
	if severity > Debug {
		severity = Debug
	}
	pri := (int(sm.Facility) * 8) + int(severity)

	temp := strings.Builder{}

	// PRIORITY
	temp.WriteString(fmt.Sprintf("<%d>%d ", pri, sm.Version))

	// TIME
	temp.WriteString(sm.Timestamp.UTC().Format(RFC339micro))
	temp.WriteString(" ")

	// HOSTNAME
	if sm.Hostname == "" {
		temp.WriteString("-")
	} else {
		temp.WriteString(sm.Hostname)
	}
	temp.WriteString(" ")

	// APPLICATION NAME
	if sm.ApplicationName == "" {
		temp.WriteString("-")
	} else {
		temp.WriteString(sm.ApplicationName)
	}
	temp.WriteString(" ")

	// PROCESS ID
	if sm.ProcessID == "" {
		temp.WriteString("-")
	} else {
		temp.WriteString(sm.ProcessID)
	}
	temp.WriteString(" ")

	// MESSAGE ID
	if sm.MessageID == "" {
		temp.WriteString("-")
	} else {
		temp.WriteString(sm.MessageID)
	}
	temp.WriteString(" ")

	// STRUCTURED DATA
	if len(sm.Data) <= 0 {
		temp.WriteString("-")
	} else {
		for _, sd := range sm.Data {
			temp.WriteString(sd.MakeString())
		}
	}
	temp.WriteString(" ")
	isUTF8 := utf8.ValidString(sm.Message)
	if isUTF8 {
		temp.WriteString("\xEF\xBB\xBF")
	}
	temp.WriteString(sm.Message)
	return temp.String()
}
