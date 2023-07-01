package eventlog

import (
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"goutils/pkg/logchain"
	"goutils/pkg/syslog"
)

func CreateMainMenu() *fyne.MainMenu {
	connectItem := fyne.MenuItem{
		Label:  "Connect",
		Action: onConnect,
		Icon:   theme.LoginIcon(),
	}

	quitItem := fyne.MenuItem{
		Label:  "Quit",
		Action: nil,
		Icon:   theme.CancelIcon(),
	}
	fileMenu := fyne.NewMenu("File", &connectItem, &quitItem)

	openLogFileItem := fyne.MenuItem{
		Label:  "Open Log File",
		Action: onShowLogFile,
		Icon:   theme.DocumentIcon(),
	}
	aboutItem := fyne.MenuItem{
		Label:  "About",
		Action: onAbout,
		Icon:   theme.InfoIcon(),
	}
	aboutMenu := fyne.NewMenu("Help", &openLogFileItem, &aboutItem)

	mainMenuBar := fyne.NewMainMenu(fileMenu, aboutMenu)
	return mainMenuBar
}

func onConnect() {
	m := GetMainWindow()
	dialog.ShowCustomConfirm("Connect to gRPC Server", "Connect", "Cancel",
		CreateConnectionDialog(), onConnectReturn, m.MainWin)
}

func onConnectReturn(connectOK bool) {
	m := GetMainWindow()
	p := m.Application.Preferences()

	if connectOK {
		client := syslog.NewClient()
		client.ConnectInfo = fmt.Sprintf("%s:%s", m.HostName, m.Port)
		isConnected, err := client.TryConnect()
		p.SetString("HostName", m.HostName)
		p.SetInt("Port", int(m.Port))
		if isConnected {
			dialog.ShowInformation("Connection Successful", "Connection OK", m.MainWin)
		} else {
			e := errors.New("Connection error.\nMore information in the log file.")
			logchain.ErrorLog("Connection error. Error:", err.Error())
			dialog.ShowError(e, m.MainWin)
		}

	}
}

func onShowLogFile() {
	lc := logchain.GetLogChain()
	m := GetMainWindow()
	err := lc.ShowLogFile()
	if err != nil {
		// Show Message
		msg := fmt.Sprintf("Cannot open the log file. Info: %s", err.Error())
		dialog.ShowInformation("Open Log File", msg, m.MainWin)
	}
}

func onAbout() {
	m := GetMainWindow()
	dialog.ShowCustom("About Event Viewer", "OK",
		CreateAboutDialog(), m.MainWin)
}
