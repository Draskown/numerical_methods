package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()

	setUi(a)

	a.Run()
}

func setUi(a fyne.App) {
	w := a.NewWindow("Fuck")
	w.SetMaster()

	w1 := a.NewWindow("Click")

	w1.SetContent(widget.NewLabel("Fucking hell mate"))

	w1.SetCloseIntercept(func() {
		w1.Hide()
	})

	b := widget.NewButton("click me", func() {
		w1.Show()
		log.Println("tapped")
	})

	vBox := container.New(layout.NewVBoxLayout(),
		widget.NewLabel("Hello there"),
		layout.NewSpacer(),
		b,
		layout.NewSpacer())

	hBox := container.New(layout.NewHBoxLayout(),
		layout.NewSpacer(),
		vBox,
		layout.NewSpacer())

	w.SetContent(hBox)
	w.Resize(fyne.NewSize(300, 400))
	w.Show()
}
