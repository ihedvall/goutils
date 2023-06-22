package logchain

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type TypeOfLogger int

const (
	LogToConsole TypeOfLogger = iota
	LogToFile
	LogToSyslog
)

type Logger struct {
	TypeOfLogger TypeOfLogger
	Name         string
	ConnectInfo  string
	ShowLocation bool
	LogLevel     Severity

	messages chan *LogMessage
}

func (l *Logger) fixEmptyPath() {
	// Empty string. Temp Dir / app name/ appName
	var appName string
	if len(os.Args) > 0 {
		filename := filepath.Base(os.Args[0])
		appName = strings.TrimSuffix(filename, filepath.Ext(filename))
	} else {
		appName = "default"
	}
	tempDir := os.TempDir()
	tempDir = filepath.Join(tempDir, appName)

	if errDir := os.MkdirAll(tempDir, os.ModePerm); errDir != nil {
		log.Println(errDir)
		return
	}
	tempDir = filepath.Join(tempDir, appName)
	tempDir += ".log"
	l.ConnectInfo = tempDir
}

func (l *Logger) fixBasePath() {
	// Empty string. Temp Dir / base name/ appName
	filename := filepath.Base(l.ConnectInfo)
	appName := strings.TrimSuffix(filename, filepath.Ext(filename))

	tempDir := os.TempDir()
	tempDir = filepath.Join(tempDir, appName)
	if errDir := os.MkdirAll(tempDir, os.ModePerm); errDir != nil {
		log.Println(errDir)
		return
	}
	tempDir = filepath.Join(tempDir, appName)
	tempDir += ".log"
	l.ConnectInfo = tempDir
}

func (l *Logger) fixLogPath() {
	filename := l.ConnectInfo
	if filename == "" {
		l.fixEmptyPath()
	} else if filename == path.Base(filename) {
		l.fixBasePath()
	} else {
		tempDir := path.Dir(l.ConnectInfo)
		if errDir := os.MkdirAll(tempDir, os.ModePerm); errDir != nil {
			log.Println(errDir)
		}
	}

}

func NewLogger(name string, loggerType TypeOfLogger, connectInfo string) *Logger {
	temp := Logger{TypeOfLogger: loggerType, Name: name, ConnectInfo: connectInfo}
	switch loggerType {
	case LogToFile:
		temp.ShowLocation = true
		temp.LogLevel = Debug
		temp.fixLogPath() // Fix any log file name failure

	case LogToSyslog:
		temp.ShowLocation = false
		temp.LogLevel = Info

	default:
		temp.ShowLocation = false
		temp.LogLevel = Trace
	}
	temp.messages = make(chan *LogMessage, 1000)
	return &temp
}

func (l *Logger) AddMessage(msg *LogMessage) {
	if msg.Severity <= l.LogLevel {
		l.messages <- msg
	}
}

func (l *Logger) NofMessages() int {
	return len(l.messages)
}

func (l *Logger) DoLogToConsole() {
	// The channels are blocking until the log chain object closes the
	// channel, returning msgOK = false
	for msg, msgOK := <-l.messages; msg != nil && msgOK; msg, msgOK = <-l.messages {
		// Output should be:
		// [Date/Time with ms] [Severity] Message  Source location
		fmt.Printf("[%s] ", msg.LocalTimeWithMs())
		fmt.Printf("[%s] ", msg.SeverityAsString())
		message := strings.TrimSpace(msg.Message)
		fmt.Printf("%s ", message)
		if l.ShowLocation && msg.File != "" {
			fmt.Printf("[%s:%d]", msg.File, msg.Line)
		}
		fmt.Println()
	}
}

func (l *Logger) SendMessages() {
	switch l.TypeOfLogger {
	case LogToFile:
		l.DoLogToConsole()
	default:
		l.DoLogToConsole()
	}

}
