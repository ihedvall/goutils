package logchain

import (
	"errors"
	"fmt"
	"os/exec"
	"sync"
)

var instance *LogChain

type LogChain struct {
	EditorPath  string // Path to an editor type notepad++ or notepad
	DefaultData []StructuredData

	loggers     map[string]*Logger
	isNotepadPP bool
	lockData    sync.Mutex
}

func TraceLog(a ...any) {
	lcLog(Trace, a)
}

func TraceLogf(format string, a ...any) {
	lcLogf(Trace, format, a)
}

func DebugLog(a ...any) {
	lcLog(Debug, a)
}

func DebugLogf(format string, a ...any) {
	lcLogf(Debug, format, a)
}

func InfoLog(a ...any) {
	lcLog(Info, a)
}

func InfoLogf(format string, a ...any) {
	lcLogf(Info, format, a)
}

func ErrorLog(a ...any) {
	lcLog(Error, a)
}

func ErrorLogf(format string, a ...any) {
	lcLogf(Error, format, a)
}

func newLogChain() *LogChain {

	lc := LogChain{loggers: make(map[string]*Logger)}
	lc.findEditor() // Search for a proper editor to open the log file in
	return &lc
}

func GetLogChain() *LogChain {
	if instance == nil {
		instance = newLogChain()
	}
	return instance
}

func (lc *LogChain) NewLogger(name string, loggerType TypeOfLogger, connectInfo string) (*Logger, error) {
	if name == "" {
		return nil, errors.New("invalid empty name")
	}
	lc.loggers[name] = NewLogger(name, loggerType, connectInfo)
	return lc.loggers[name], nil
}

func (lc *LogChain) Clear() {
	lc.loggers = make(map[string]*Logger)
}

func (lc *LogChain) AddLogMessage(message *LogMessage) {
	if message == nil {
		return
	}
	for _, logger := range lc.loggers {
		if logger == nil {
			continue
		}
		logger.AddMessage(message)
	}
}

func (lc *LogChain) IsEmpty() bool {
	for _, logger := range lc.loggers {
		if logger == nil {
			continue
		}
		if logger.NofMessages() != 0 {
			return false
		}
	}
	return true
}

func (lc *LogChain) Start() (bool, string) {
	for _, logger := range lc.loggers {
		if logger == nil {
			continue
		}
		go logger.SendMessages()
	}
	return true, ""
}

func (lc *LogChain) Stop() (bool, string) {
	for _, logger := range lc.loggers {
		if logger == nil {
			continue
		}
		close(logger.messages) // Stop all running threads aka goroutines
	}

	return true, ""
}

func (lc *LogChain) GetLogger(name string) *Logger {
	return lc.loggers[name]
}

func (lc *LogChain) GetLogFilePath() (filename string, fileOK bool) {
	for _, logger := range lc.loggers {
		if logger.TypeOfLogger == LogToFile && logger.ConnectInfo != "" {
			return logger.ConnectInfo, true
		}
	}
	return "", false
}

func (lc *LogChain) ShowLogFile() error {
	file, _ := lc.GetLogFilePath()
	if file == "" {
		return errors.New("no log file defined")
	}
	if lc.EditorPath == "" {
		return errors.New("no default editor found")
	}
	if !FileExist(file) {
		return errors.New("no log file created yet")
	}
	var cmd *exec.Cmd
	var startErr error
	if lc.isNotepadPP {
		cmd = exec.Command(lc.EditorPath, "-monitor", file)
	} else {
		cmd = exec.Command(lc.EditorPath, file)
	}
	startErr = cmd.Start()
	return startErr
}

func (lc *LogChain) GetDefaultData(key string) (string, error) {
	lc.lockData.Lock()
	defer lc.lockData.Unlock()

	for _, sd := range lc.DefaultData {
		for k := range sd.ParameterList {
			if k == key {
				return sd.ParameterList[key], nil
			}
		}
	}
	return "", errors.New("no data found")
}

func (lc *LogChain) SetDefaultData(key string, value string) {
	lc.lockData.Lock()
	defer lc.lockData.Unlock()

	for _, sd := range lc.DefaultData {
		for k := range sd.ParameterList {
			if k == key {
				sd.ParameterList[key] = value
			}
		}
	}
}

func (lc *LogChain) findEditor() {
	// Try Notepad++
	if lc.EditorPath == "" {
		path, err := exec.LookPath("notepad++")
		if path != "" && err == nil {
			lc.EditorPath = path
			lc.isNotepadPP = true
		}
	}
	if lc.EditorPath == "" {
		path, err := exec.LookPath("c:/program files/notepad++/notepad++")
		if path != "" && err == nil {
			lc.EditorPath = path
			lc.isNotepadPP = true
		}
	}
	// Try Notepad
	if lc.EditorPath == "" {
		path, err := exec.LookPath("notepad")
		if path != "" && err == nil {
			lc.EditorPath = path
		}
	}

	// Try gedit
	if lc.EditorPath == "" {
		path, err := exec.LookPath("gedit")
		if path != "" && err == nil {
			lc.EditorPath = path
		}
	}

	// Try kate
	if lc.EditorPath == "" {
		path, err := exec.LookPath("kate")
		if path != "" && err == nil {
			lc.EditorPath = path
		}
	}
	// Todo: Better finding on macOS and Linux
}

func lcLog(severity Severity, a ...any) {
	text := fmt.Sprint(a)
	lc := GetLogChain()
	if text != "" && lc != nil {
		msg := NewLogMessage(severity, text)
		lc.AddLogMessage(msg)
	}
}

func lcLogf(severity Severity, format string, a ...any) {
	text := fmt.Sprintf(format, a)
	lc := GetLogChain()
	if text != "" && lc != nil {
		msg := NewLogMessage(severity, text)
		lc.AddLogMessage(msg)
	}
}
