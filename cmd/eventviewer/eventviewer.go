package main

import (
	"fyne.io/fyne/v2/app"
	"goutils/internal/pkg/eventlog"
	"goutils/pkg/logchain"
)

func main() {
	lc := logchain.GetLogChain()
	fileLogger, _ := lc.NewLogger("File", logchain.LogToFile, "eventviewer.log")
	fileLogger.LogLevel = logchain.Trace
	fileLogger.ShowLocation = true
	lc.Start()
	logchain.DebugLog("Starting application")

	application := app.NewWithID("eventviewer")
	mainWindow := eventlog.NewMainWindow(application)
	mainWindow.MainWin.CenterOnScreen()
	mainWindow.MainWin.ShowAndRun()

	logchain.DebugLog("Stopping application")
	lc.Stop()
	lc.Clear()
}
