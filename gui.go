package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/sqweek/dialog"
)

var a = app.New()

func BadPath(Message string) {
	b := a.NewWindow("Error")
	b.SetContent(container.NewVBox(widget.NewLabel(Message), widget.NewButton("OK", func() { b.Close() })))
	b.Show()
}

func Dialog() string {

	filename, _ := dialog.File().Filter("Select CSV-File", "csv").Load()

	return filename

}

func SaveDialog() string {
	filename, _ := dialog.File().Filter("Save JSON-Fiel", "json").Title("Export to JSON").Save()
	return filename
}

func GUI() {

	w := a.NewWindow("CSV to JSON-Converter")
	w.Resize(fyne.Size{800, 300})
	hello := widget.NewLabel("Path to File:")
	path := widget.NewEntry()
	w.SetContent(container.NewVBox(hello, path, widget.NewButton("Select CSV-File", func() { path.SetText(Dialog()) }), widget.NewButton("Convert to JSON", func() { Conversion(path.Text) })))
	w.ShowAndRun()
}
