package eventlog

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	//	widgetx "fyne.io/x/fyne/widget"
	"goutils/pkg/logchain"
	"time"
)

type FilterView struct {
	ClearButton *widget.Button

	container *fyne.Container
}

var severityList = []string{
	"Emergency",
	"Alert",
	"Critical",
	"Error",
	"Warning",
	"Notice",
	"Info",
	"Debug",
}

func NewFilterView() *FilterView {

	severityLabel := widget.NewLabel("Severity Level:")
	severitySelect := widget.NewSelect(severityList, nil)
	severitySelect.SetSelectedIndex(int(logchain.Debug))
	severityC := container.NewVBox(severityLabel, severitySelect)

	msgLabel := widget.NewLabel("Message:           ")
	msgEntry := widget.NewEntry()
	msgEntry.SetText("*")
	msgC := container.NewVBox(msgLabel, msgEntry)

	dateLabel := widget.NewLabel("From Date:   ")
	now := time.Now()
	date := widget.NewButton(now.Format(time.DateOnly), nil)
	dateC := container.NewVBox(dateLabel, date)

	limitLabel := widget.NewLabel("Limit:          ")
	limit := widget.NewEntry()
	limit.SetText(fmt.Sprintf("%d", 1000))
	limitC := container.NewVBox(limitLabel, limit)

	offsetLabel := widget.NewLabel("Offset:        ")
	offset := widget.NewEntry()
	offset.SetText(fmt.Sprintf("%d", 0))
	offsetC := container.NewVBox(offsetLabel, offset)

	clearLabel := widget.NewLabel("Clear Filter:")
	clearButton := widget.NewButtonWithIcon("", resourceFilterlistoffSvg, nil)
	clearC := container.NewVBox(clearLabel, clearButton)

	c := container.NewHBox(
		severityC,
		msgC,
		limitC,
		offsetC,
		dateC,
		clearC)

	fw := FilterView{
		ClearButton: clearButton,
		container:   c,
	}
	return &fw
}
