package eventlog

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

type MainWindow struct {
	Application fyne.App
	MainWin     fyne.Window
	HostName    string
	Port        uint16 // Should be a numeric value but fyne only handle strings
}

var instanceMain *MainWindow

func NewMainWindow(application fyne.App) *MainWindow {
	mainWin := application.NewWindow("Event Viewer")
	p := application.Preferences()
	main := MainWindow{
		Application: application,
		MainWin:     mainWin,
		HostName:    p.StringWithFallback("HostName", "127.0.0.1"),
		Port:        uint16(p.IntWithFallback("Port", 50600)),
	}
	main.drawMainWin()
	instanceMain = &main
	return &main
}

func GetMainWindow() *MainWindow {
	return instanceMain
}

func (m *MainWindow) drawMainWin() {
	w := m.MainWin
	w.SetMainMenu(CreateMainMenu())
	w.SetIcon(resourceAppiconPng)
	filterView := NewFilterView()
	mainContainer := container.NewBorder(nil, filterView.container, nil, nil)
	w.SetContent(mainContainer)

	w.Resize(fyne.NewSize(1200, 800))
}
