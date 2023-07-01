package eventlog

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"strings"
)

func CreateAboutDialog() *fyne.Container {
	image := canvas.NewImageFromResource(resourceAppiconPng)
	image.FillMode = canvas.ImageFillOriginal
	iconC := container.NewCenter(image)

	appText := widget.NewLabel("Event Viewer 1.0")
	appText.TextStyle.Bold = true

	descText := widget.NewLabel("Syslog client that shows syslog as event messages")

	infoC := container.NewVBox(appText, descText)

	top := container.NewHBox(iconC, infoC)

	licenseLabel := widget.NewLabel("License")
	licenseLabel.TextStyle.Bold = true

	temp := strings.Builder{}
	temp.WriteString("MIT License (https://opensource.org/licenses/MIT)\n")
	temp.WriteString("Copyright (C) 2023 Ingemar Hedvall\n")
	temp.WriteString("\n")

	temp.WriteString("Permission is hereby granted, free of charge, to any person obtaining a copy of this\n")
	temp.WriteString("software and associated documentation files (the \"Software\"),\n")
	temp.WriteString("to deal in the Software without restriction, including without limitation the rights to use, copy,\n")
	temp.WriteString("modify, merge, publish, distribute, sublicense, and/or sell copies of the Software,\n")
	temp.WriteString("and to permit persons to whom the Software is furnished to do so, subject to the following conditions:\n")
	temp.WriteString("\n")
	temp.WriteString("The above copyright notice and this permission notice shall be included in all copies or substantial\n")
	temp.WriteString("portions of the Software.\n")
	temp.WriteString("\n")
	temp.WriteString("THE SOFTWARE IS PROVIDED \"AS IS\", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED,\n")
	temp.WriteString("INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR\n")
	temp.WriteString("PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,\n")
	temp.WriteString("DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR\n")
	temp.WriteString("IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.")
	licenseText := widget.NewLabel(temp.String())

	devLabel := widget.NewLabel("Developers")
	devLabel.TextStyle.Bold = true

	devText1 := widget.NewLabel("Ingemar Hedvall")

	c := container.NewVBox(top, licenseLabel, licenseText, devLabel, devText1)
	return c
}
