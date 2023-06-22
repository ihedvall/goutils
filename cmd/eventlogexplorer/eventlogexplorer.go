package main

import (
	"fyne.io/fyne/v2/app"
	"goutils/internal/pkg/eventlogexplorer"
)

func main() {
	application := app.New()
	mainWindow := eventlogexplorer.NewMainWindow(application)
	mainWindow.CenterOnScreen()
	mainWindow.ShowAndRun()
}
