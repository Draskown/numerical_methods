package lr3

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// Sets a GUI for passed application
func SetUi(a fyne.App) {
	// Create the main window
	w := a.NewWindow("LR3")
	w.Resize(fyne.NewSize(300, 100))
	w.SetMaster()

	wEuler := a.NewWindow("Euler")

	wEuler.SetCloseIntercept(func() {
		wEuler.Hide()
	})

	// Set a layout for the Euler window
	vBoxEulerX := container.NewVBox()
	vBoxEulerY := container.NewVBox()

	vBoxEulerX.Add(widget.NewLabel("One equation:"))
	vBoxEulerY.Add(widget.NewLabel(""))

	// Set a button to calculate and fill the Euler's window
	btnEuler := widget.NewButton("Euler", func() {
		x, y := Euler()

		for i := -1; i < len(x); i++ {
			xText := "X"
			yText := "Y"

			if (i > -1){
				xText = fmt.Sprintf("%.2f", x[i])
				yText = fmt.Sprintf("%.2f", y[i])
			}
			
			vBoxEulerX.Add(widget.NewLabelWithStyle(
				xText,
				fyne.TextAlignCenter, 
				fyne.TextStyle{}))

			vBoxEulerY.Add(widget.NewLabelWithStyle(
				yText,
				fyne.TextAlignCenter, 
				fyne.TextStyle{}))
		}

		hBoxEuler := container.NewHBox(
			vBoxEulerX,
			vBoxEulerY,
			widget.NewSeparator(),
		)

		wEuler.SetContent(hBoxEuler)

		wEuler.Show()
	})

	// Set a layout for the main window
	vBoxMain := container.NewVBox(
		layout.NewSpacer(),
		btnEuler,
		layout.NewSpacer(),
	)

	hBoxMain := container.NewHBox(
		layout.NewSpacer(),
		vBoxMain,
		layout.NewSpacer(),
	)

	w.SetContent(hBoxMain)
	w.Show()
}
