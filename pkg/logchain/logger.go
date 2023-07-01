package logchain

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"
	"strings"
	"time"
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
		if temp.ConnectInfo == "" {
			temp.ConnectInfo = "127.0.0.1:50600"
		}

	default:
		temp.ShowLocation = false
		temp.LogLevel = Trace
	}
	temp.messages = make(chan *LogMessage, 1000)
	return &temp
}

func FileExist(file string) bool {
	if _, err := os.Stat(file); err == nil {
		return true
	}
	return false
}

func CopyFile(src string, dst string) (int64, error) {
	sourceFileStat, errSrc := os.Stat(src)
	if errSrc != nil {
		return 0, errSrc
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, errSrcOpen := os.Open(src)
	if errSrcOpen != nil {
		return 0, errSrcOpen
	}
	defer func(source *os.File) {
		errSrcClose := source.Close()
		if errSrcClose != nil {
		}
	}(source)

	destination, errCreate := os.Create(dst)
	if errCreate != nil {
		return 0, errCreate
	}
	defer func(destination *os.File) {
		errDstClose := destination.Close()
		if errDstClose != nil {

		}
	}(destination)

	nBytes, errCopy := io.Copy(destination, source)
	return nBytes, errCopy
}

func FileSize(filename string) int64 {
	fileInfo, err := os.Stat(filename)
	if err != nil {
		return 0
	}
	return fileInfo.Size()
}

func BackupFiles(filename string, removeFile bool) (bool, error) {
	if filename == "" {
		return false, errors.New("file name cannot be empty")
	}
	full, errAbs := filepath.Abs(filename)
	if errAbs != nil {
		return false, nil
	}
	parentDir := filepath.Dir(full)
	base := filepath.Base(full)
	ext := filepath.Ext(base)
	stem := strings.TrimSuffix(base, ext) // File name without extension

	// If base file doesn't exist, continue as nothing happened
	if !FileExist(full) {
		return true, nil
	}

	// shift all file xxx_N -> xxx_N-1 and last xxx -> xxx_0
	for i := 9; i >= 0; i-- {
		temp1 := fmt.Sprintf("%s_%d%s", stem, i, ext)
		file1 := filepath.Join(parentDir, temp1)
		if i == 9 && FileExist(file1) {
			errRemove := os.Remove(file1)
			if errRemove != nil {
				return false, errRemove
			}
		}

		if i == 0 {
			if removeFile {
				errRename := os.Rename(full, file1)
				if errRename != nil {
					return false, errRename
				}
			} else {
				_, errCopy := CopyFile(full, file1)
				if errCopy != nil {
					return false, errCopy
				}
			}
		} else {
			temp2 := fmt.Sprintf("%s_%d%s", stem, i-1, ext)
			file2 := filepath.Join(parentDir, temp2)
			if FileExist(file1) {
				errRemove := os.Remove(file1)
				if errRemove != nil {
					return false, errRemove
				}
			}
			if FileExist(file2) {
				errRename2 := os.Rename(file2, file1)
				if errRename2 != nil {
					return false, errRename2
				}
			}
		}
	}

	return true, nil
}

func (l *Logger) AddMessage(msg *LogMessage) {
	if msg.Severity <= l.LogLevel {
		l.messages <- msg
	}
}

func (l *Logger) NofMessages() int {
	return len(l.messages)
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
	} else if filename == filepath.Base(filename) {
		l.fixBasePath()
	} else {
		tempDir := filepath.Dir(filename)
		if errDir := os.MkdirAll(tempDir, os.ModePerm); errDir != nil {
			log.Println(errDir)
		}
	}

}

func (l *Logger) SendMessages() {
	switch l.TypeOfLogger {
	case LogToFile:
		l.doLogToFile()

	case LogToSyslog:
		l.doLogToSyslog()

	default:
		l.doLogToConsole()
	}

}

func (l *Logger) doLogToConsole() {
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

func (l *Logger) doLogToFile() {
	filename := l.ConnectInfo
	_, errBackup := BackupFiles(filename, true)
	if errBackup != nil {
		// Not a good sign but just try to continue
	}
	// The channels are blocking until the log chain object closes the
	// channel, returning msgOK = false
	var fileOpen = false
	var file *os.File = nil
	var errOpen error = nil

	for msg, msgOK := <-l.messages; msg != nil && msgOK; msg, msgOK = <-l.messages {

		if !fileOpen {
			file, errOpen = os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
			if errOpen == nil && file != nil {
				fileOpen = true
			}
		}
		if fileOpen {
			// Write to file
			// Output should be:
			// [Date/Time with ms] [Severity] Message  Source location
			_, errWrite := fmt.Fprintf(file, "[%s] ", msg.LocalTimeWithMs())
			if errWrite == nil {
				_, errWrite = fmt.Fprintf(file, "[%s] ", msg.SeverityAsString())
			}

			message := strings.TrimSpace(msg.Message)
			if errWrite == nil {
				_, errWrite = fmt.Fprintf(file, "%s ", message)
			}
			if l.ShowLocation && msg.File != "" && errWrite == nil {
				_, errWrite = fmt.Fprintf(file, "[%s:%d]", msg.File, msg.Line)
			}
			if errWrite == nil {
				_, errWrite = fmt.Fprintln(file)
			}
			if errWrite != nil {
			}
		}
		if len(l.messages) <= 0 && fileOpen {
			// Close file
			errClose := file.Close()
			if errClose == nil {
				file = nil
			}
			fileOpen = false
			if FileExist(filename) && FileSize(filename) > 10000000 {
				_, errBackup = BackupFiles(filename, true)
			}
			time.Sleep(time.Second) // Do not save  more than each second
		}
	}
}

func (l *Logger) doLogToSyslog() {

	// The channels are blocking until the log chain object closes the
	// channel, returning msgOK = false
	var connectionOpen = false
	var connection net.Conn
	var errOpen error = nil

	for msg, msgOK := <-l.messages; msg != nil && msgOK; msg, msgOK = <-l.messages {
		if !connectionOpen {
			connection, errOpen = net.Dial("udp", l.ConnectInfo)
			if errOpen == nil && connection != nil {
				connectionOpen = true
			}
		}
		if connectionOpen {
			sm := NewSyslogMessage()
			sm.Populate(msg)
			_, errWrite := fmt.Fprint(connection, sm.MakeString())
			if errWrite != nil {
			}
		}
		if len(l.messages) <= 0 && connectionOpen {
			// Close file
			errClose := connection.Close()
			if errClose == nil {
				connection = nil
			}
			connectionOpen = false
		}
	}
}
