package eventlogexplorer

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"goutils/pkg/logchain"
	"image/color"
)

type MainWindow struct {
	Application *fyne.App
	MainWin     *fyne.Window
}

var mainWindowInstance *MainWindow

func GetMainWindow() *MainWindow {
	return mainWindowInstance
}

func makeSign() fyne.CanvasObject {
	bg := canvas.NewCircle(color.NRGBA{255, 0, 0, 255})
	bg.StrokeColor = color.White
	bg.StrokeWidth = 5
	bg.Resize(fyne.NewSize(100, 100))
	bg.Move(fyne.NewPos(10, 10))

	bar := canvas.NewRectangle(color.White)
	bar.Resize(fyne.NewSize(80, 20))
	bar.Move(fyne.NewPos(20, 50))

	c := container.NewWithoutLayout(bg, bar)
	return c
}

func makeText() fyne.CanvasObject {
	text1 := canvas.NewText("Hello Mommy", color.NRGBA{R: 255, A: 255})
	text2 := canvas.NewText("Hello Daddy", color.NRGBA{R: 255, A: 100})
	//	text.Resize(fyne.NewSize(100, 100))
	//	text.Move(fyne.NewPos(1, 1))
	c := container.NewHBox(text1, text2)
	return c
}

func onTableSize() (rows int, cols int) {
	return 1000, 4
}
func onCreateCell() fyne.CanvasObject {
	item := widget.NewLabel("XXX")
	item.Wrapping = fyne.TextWrapBreak
	return item
}

func onUpdateCell(id widget.TableCellID, obj fyne.CanvasObject) {
	if id.Row < 0 {
		return
	}
	item := obj.(*widget.Label)
	switch id.Col {
	case 0:
		item.SetText("2023-06-20 HH::MM::SS.sss")

	case 1:
		item.SetText("A")

	case 2:
		item.SetText("Message")

	case 3:
		item.SetText("Information")

	default:
		return
	}
}
func onCreateHeader() fyne.CanvasObject {
	header := widget.NewButton("XXX", nil)
	return header
}
func onUpdateHeader(id widget.TableCellID, obj fyne.CanvasObject) {
	header := obj.(*widget.Button)
	if id.Col < 0 { //
		header.SetText("")
	} else {
		switch id.Col {
		case 0:
			header.SetText("Time")
		case 1:
			header.SetText("Type")
		case 2:
			header.SetText("Message")
		default:
			header.SetText("Information")
		}
	}
}

func makeList() fyne.CanvasObject {

	list := widget.NewTableWithHeaders(onTableSize, onCreateCell, onUpdateCell)
	list.ShowHeaderColumn = false
	list.ShowHeaderRow = true
	list.CreateHeader = onCreateHeader
	list.UpdateHeader = onUpdateHeader

	text := canvas.NewText("YYYY-MM-DD HH:MM:SS.sss ", color.Opaque)
	width := text.MinSize().Width
	list.SetColumnWidth(0, width)

	text.Text = "WAEX "
	width = text.MinSize().Width
	list.SetColumnWidth(1, width)

	text.Text = "XXXX XXXX XXXXX XXXXX XXXXX "
	width = text.MinSize().Width
	list.SetColumnWidth(2, width)

	text.Text = "XXXX XXXX XXXXX XXXXX XXXXX "
	width = text.MinSize().Width
	list.SetColumnWidth(3, width)

	return list
}

func NewMainWindow(application fyne.App) fyne.Window {
	main := application.NewWindow("Event Log Explorer")
	mainWindowInstance = &MainWindow{Application: &application, MainWin: &main}
	main.SetMainMenu(CreateMainMenu())
	listContainer := container.NewStack(makeList())

	mainContainer := container.NewBorder(nil, makeText(), nil, nil)
	mainContainer.Add(listContainer)

	// c.Resize(fyne.NewSize(200, 200))

	//b := canvas.NewText("Hello", color.White)
	//b.Resize(fyne.NewSize(100, 10))
	//c := fyne.Container{
	//	Hidden:  false,
	//	Layout:  layout.NewBorderLayout(nil, b, nil, nil),
	//	Objects: b,
	//}

	main.SetContent(mainContainer)

	// mainWindow.SetContent(makeSign())
	//mainWindow.SetPadded(false)
	main.Resize(fyne.NewSize(800, 600))

	return main
}

func OnShowLogFile() {
	lc := logchain.GetLogChain()
	_ = lc.ShowLogFile()
}

func OnDisableLogFile() bool {
	lc := logchain.GetLogChain()
	filename, fileOK := lc.GetLogFilePath()
	return lc.EditorPath == "" || filename == "" || !fileOK
}

func CreateMainMenu() *fyne.MainMenu {

	connectItem := fyne.MenuItem{
		Label:  "Connect",
		Action: nil,
		Icon:   theme.LoginIcon(),
	}

	quitItem := fyne.MenuItem{
		Label:  "Quit",
		Action: nil,
		Icon:   theme.CancelIcon(),
	}
	fileMenu := fyne.NewMenu("File", &connectItem, &quitItem)

	openLogFileItem := fyne.MenuItem{
		Label:    "Open Log File",
		Action:   OnShowLogFile,
		Disabled: OnDisableLogFile(),
		Icon:     theme.DocumentIcon(),
	}
	aboutItem := fyne.MenuItem{
		Label:  "About",
		Action: nil,
		Icon:   theme.InfoIcon(),
	}
	aboutMenu := fyne.NewMenu("About", &openLogFileItem, &aboutItem)

	mainMenuBar := fyne.NewMainMenu(fileMenu, aboutMenu)
	return mainMenuBar
}
