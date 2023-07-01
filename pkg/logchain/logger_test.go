package logchain

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestEmptyLogPath(t *testing.T) {
	fileLogger := NewLogger("File", LogToFile, "")
	if fileLogger == nil {
		t.Fatal("NewLogger() return nil.")
	}
	logPath := fileLogger.ConnectInfo
	if logPath == "" {
		t.Error("Getting default log path failed.")
	}

	t.Logf("Log Path: %s", logPath)
}

func TestBaseLogPath(t *testing.T) {
	baseFile := "basetest.log"
	fileLogger := NewLogger("File", LogToFile, baseFile)
	if fileLogger == nil {
		t.Fatal("NewLogger() return nil.")
	}
	logPath := fileLogger.ConnectInfo
	if logPath == "" {
		t.Error("Getting base log path failed.")
	}
	baseTest := filepath.Base(logPath)
	if baseFile != baseTest {
		t.Errorf("Failed to get the right base file. %s/%s", baseFile, baseTest)
	}
	t.Logf("Log Path: %s", logPath)
}

func TestFullLogPath(t *testing.T) {
	file := filepath.Join(t.TempDir(), "fulltest.log")
	fileLogger := NewLogger("File", LogToFile, file)
	if fileLogger == nil {
		t.Fatal("NewLogger() return nil.")
	}
	logPath := fileLogger.ConnectInfo
	if logPath == "" {
		t.Error("Getting base log path failed.")
	}
	if file != logPath {
		t.Errorf("Failed to get the right base file. %s/%s", file, logPath)
	}
	t.Logf("Log Path: %s", logPath)
}

func TestFileFunctions(t *testing.T) {
	rootDir := t.TempDir()
	tempFile := filepath.Join(rootDir, "existtest.txt")
	tempFile1 := filepath.Join(rootDir, "existtest_0.txt")
	copyFile := filepath.Join(rootDir, "copytest.txt")
	errRemove := os.Remove(tempFile)
	if errRemove != nil {
		t.Log(errRemove.Error())
	}

	if FileExist(tempFile) {
		t.Error("Reported exist of file not existing.")
	}
	file, errCreate := os.Create(tempFile)
	if errCreate != nil {
		t.Errorf("Failed to create temp file. Error: %s", errCreate.Error())
	}

	if !FileExist(tempFile) {
		t.Error("Reported not exist on opened file.")
	}

	_, errWrite := file.WriteString("Kilroy was here")
	if errWrite != nil {
		t.Error(errWrite.Error())
	}

	errClose := file.Close()
	if errClose != nil {
		t.Error(errClose)
	}

	if !FileExist(tempFile) {
		t.Error("Reported not exist on closed file.")
	}

	nofBytes, errCopy := CopyFile(tempFile, copyFile)
	if errCopy != nil {
		t.Errorf("Copy error. Error: %s", errCopy.Error())
	}
	if nofBytes <= 0 {
		t.Error("Invalid number of bytes.")
	}
	if !FileExist(copyFile) {
		t.Error("Reported not exist on copied file.")
	}
	// Check overwrite
	nofBytes, errCopy = CopyFile(tempFile, copyFile)
	if errCopy != nil {
		t.Errorf("Overwrite test failed. Error: %s", errCopy.Error())
	}

	backup, errBackup := BackupFiles(tempFile, true)
	if errBackup != nil {
		t.Errorf("Backup failed. Error: %s", errBackup.Error())
	}
	if !backup {
		t.Error("Backup failed.")
	}
	if FileExist(tempFile) {
		t.Error("Original file should have been deleted")
	}
	if !FileExist(tempFile1) {
		t.Error("Backup file doesn't exist")
	}

	if FileSize(tempFile1) <= 0 {
		t.Error("Backup file is empty")
	}
}

func TestTime(t *testing.T) {
	now := time.Now()
	t.Log(now.UTC())
	t.Log(now.UTC().Year())

	now = time.Time{}
	t.Log(now.UTC())
	t.Log(now.UTC().Year())
}
