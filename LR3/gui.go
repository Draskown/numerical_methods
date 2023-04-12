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

	// Set a window for the Euler's solution
	wEuler := a.NewWindow("Euler")

	wEuler.SetCloseIntercept(func() {
		wEuler.Hide()
	})

	// Set a layout for the Euler window
	vBoxEulerX := container.NewVBox()
	vBoxEulerY := container.NewVBox()

	vBoxEulerX.Add(widget.NewLabel("One equation:"))
	vBoxEulerY.Add(widget.NewLabel(""))

	vBoxEulerXS := container.NewVBox()
	vBoxEulerYS := container.NewVBox()
	vBoxEulerZS := container.NewVBox()

	vBoxEulerXS.Add(widget.NewLabel("System of equations:"))
	vBoxEulerYS.Add(widget.NewLabel(""))
	vBoxEulerZS.Add(widget.NewLabel(""))

	// Set a button to calculate and fill the Euler's window
	btnEuler := widget.NewButton("Euler", func() {
		x, y := Euler()

		var z []float64

		var xText string
		var yText string
		var zText string

		for i := -1; i < len(x); i++ {
			xText = "X"
			yText = "Y"

			if (i > -1){
				xText = fmt.Sprintf("%.2f", x[i])
				yText = fmt.Sprintf("%.3f", y[i])
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

		x, y, z = EulerSystem()

		for i := -1; i < len(x); i++ {
			xText = "X"
			yText = "Y"
			zText = "Z"

			if (i > -1){
				xText = fmt.Sprintf("%.2f", x[i])
				yText = fmt.Sprintf("%.3f", y[i])
				zText = fmt.Sprintf("%.3f", z[i])
			}
			
			vBoxEulerXS.Add(widget.NewLabelWithStyle(
				xText,
				fyne.TextAlignCenter, 
				fyne.TextStyle{}))

			vBoxEulerYS.Add(widget.NewLabelWithStyle(
				yText,
				fyne.TextAlignCenter, 
				fyne.TextStyle{}))

			vBoxEulerZS.Add(widget.NewLabelWithStyle(
				zText,
				fyne.TextAlignCenter, 
				fyne.TextStyle{}))
		}

		hBox := container.NewHBox(
			vBoxEulerX,
			vBoxEulerY,
			widget.NewSeparator(),
			vBoxEulerXS,
			vBoxEulerYS,
			vBoxEulerZS,
		)

		wEuler.SetContent(hBox)

		wEuler.Show()
	})

	// Set a window for the Runge-Kutte's solution
	wRungeKutte := a.NewWindow("Runge-Kutte")

	wRungeKutte.SetCloseIntercept(func() {
		wRungeKutte.Hide()
	})

	// Set a layout for the Runge-Kutte window
	vBoxRKX := container.NewVBox()
	vBoxRKY := container.NewVBox()

	vBoxRKX.Add(widget.NewLabel("One equation:"))
	vBoxRKY.Add(widget.NewLabel(""))

	vBoxRKXS := container.NewVBox()
	vBoxRKYS := container.NewVBox()
	vBoxRKZS := container.NewVBox()

	vBoxRKXS.Add(widget.NewLabel("System of equations:"))
	vBoxRKYS.Add(widget.NewLabel(""))
	vBoxRKZS.Add(widget.NewLabel(""))

	// Set a button to calculate and fill the Euler's window
	btnRungeKutta := widget.NewButton("Runge-Kutte", func() {
		x, y := RungeKutte()

		var z []float64

		var xText string
		var yText string
		var zText string

		for i := -1; i < len(x); i++ {
			xText = "X"
			yText = "Y"

			if (i > -1){
				xText = fmt.Sprintf("%.2f", x[i])
				yText = fmt.Sprintf("%.3f", y[i])
			}
			
			vBoxRKX.Add(widget.NewLabelWithStyle(
				xText,
				fyne.TextAlignCenter, 
				fyne.TextStyle{}))

			vBoxRKY.Add(widget.NewLabelWithStyle(
				yText,
				fyne.TextAlignCenter, 
				fyne.TextStyle{}))
		}

		x, y, z = RungeKutteSystem()

		for i := -1; i < len(x); i++ {
			xText = "X"
			yText = "Y"
			zText = "Z"

			if (i > -1){
				xText = fmt.Sprintf("%.2f", x[i])
				yText = fmt.Sprintf("%.3f", y[i])
				zText = fmt.Sprintf("%.3f", z[i])
			}
			
			vBoxRKXS.Add(widget.NewLabelWithStyle(
				xText,
				fyne.TextAlignCenter, 
				fyne.TextStyle{}))

			vBoxRKYS.Add(widget.NewLabelWithStyle(
				yText,
				fyne.TextAlignCenter, 
				fyne.TextStyle{}))

			vBoxRKZS.Add(widget.NewLabelWithStyle(
				zText,
				fyne.TextAlignCenter, 
				fyne.TextStyle{}))
		}

		hBox := container.NewHBox(
			vBoxRKX,
			vBoxRKY,
			widget.NewSeparator(),
			vBoxRKXS,
			vBoxRKYS,
			vBoxRKZS,
		)

		wRungeKutte.SetContent(hBox)

		wRungeKutte.Show()
	})

	// Set a layout for the main window
	vBoxMain := container.NewVBox(
		layout.NewSpacer(),
		btnEuler,
		layout.NewSpacer(),
		btnRungeKutta,
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
