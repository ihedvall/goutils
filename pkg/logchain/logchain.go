package logchain

import (
	"errors"
	"os/exec"
)

var instance *LogChain

type LogChain struct {
	EditorPath string // Path to an editor type notepad++ or notepad

	loggers map[string]*Logger
}

func newLogChain() *LogChain {

	lc := LogChain{loggers: make(map[string]*Logger)}

	// Try Notepad++
	if lc.EditorPath == "" {
		path, err := exec.LookPath("notepad++")
		if path != "" && err == nil {
			lc.EditorPath = path
		}
	}
	if lc.EditorPath == "" {
		path, err := exec.LookPath("c:/program files/notepad++/notepad++")
		if path != "" && err == nil {
			lc.EditorPath = path
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
		close(logger.messages)
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

func (lc *LogChain) ShowLogFile() bool {

	file, _ := lc.GetLogFilePath()
	if file == "" || lc.EditorPath == "" {
		return false
	}

	cmd := exec.Command(lc.EditorPath, file)
	startErr := cmd.Start()
	return startErr == nil
}
