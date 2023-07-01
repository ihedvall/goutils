package eventlog

import (
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	widgetx "fyne.io/x/fyne/widget"
	"strconv"
)

func CreateConnectionDialog() *fyne.Container {
	m := GetMainWindow()
	h := binding.BindString(&m.HostName)

	hostLabel := widget.NewLabel("Host Name:")
	hostName := widget.NewEntryWithData(h)

	portLabel := widget.NewLabel("Host Name:")
	port := widgetx.NewNumericalEntry()
	port.AllowFloat = false
	port.Validator = PortValidator
	port.SetText(fmt.Sprintf("%d", m.Port))

	c := container.New(layout.NewFormLayout(), hostLabel, hostName, portLabel, port)
	return c
}

func PortValidator(value string) error {
	v, err := strconv.Atoi(value)
	if err != nil {
		return err
	}
	if v <= 0 || v > 65535 {
		return errors.New("valid range is 1..65535")
	}
	m := GetMainWindow()
	m.Port = uint16(v)
	return nil
}
