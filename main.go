package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
)

type Personnage struct {
	Bonjour int
}

func makeUI() (*widget.Label, *widget.Entry) {
	return widget.NewLabel("Hello world!"),
		widget.NewEntry()
}

func main() {
	a := app.New()
	w := a.NewWindow("SPOTIFI")
	w.Resize(fyne.NewSize(200, 200))
	w.SetContent(widget.NewLabel("Hello World"))
	w.SetContent(container.NewVBox(makeUI()))

	w.ShowAndRun()

}
