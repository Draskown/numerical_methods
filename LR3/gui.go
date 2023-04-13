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
	vBoxEX := container.NewVBox()
	vBoxEY := container.NewVBox()

	vBoxEX.Add(widget.NewLabel("One equation:"))
	vBoxEY.Add(widget.NewLabel(""))

	vBoxEXS := container.NewVBox()
	vBoxEYS := container.NewVBox()
	vBoxEZS := container.NewVBox()

	vBoxEXS.Add(widget.NewLabel("System of equations:"))
	vBoxEYS.Add(widget.NewLabel(""))
	vBoxEZS.Add(widget.NewLabel(""))

	// Calculate the Euler's solution
	x, y := Euler()

	var z []float64

	var xText string
	var yText string
	var zText string

	// Fill in the answers
	for i := -1; i < len(x); i++ {
		xText = "X"
		yText = "Y"

		if (i > -1){
			xText = fmt.Sprintf("%.2f", x[i])
			yText = fmt.Sprintf("%.3f", y[i])
		}
		
		vBoxEX.Add(widget.NewLabelWithStyle(
			xText,
			fyne.TextAlignCenter, 
			fyne.TextStyle{}))

		vBoxEY.Add(widget.NewLabelWithStyle(
			yText,
			fyne.TextAlignCenter, 
			fyne.TextStyle{}))
	}

	// Calculate the Euler's solution for the system
	x, y, z = EulerSystem()

	// Fill in the answers
	for i := -1; i < len(x); i++ {
		xText = "X"
		yText = "Y"
		zText = "Z"

		if (i > -1){
			xText = fmt.Sprintf("%.2f", x[i])
			yText = fmt.Sprintf("%.3f", y[i])
			zText = fmt.Sprintf("%.3f", z[i])
		}
		
		vBoxEXS.Add(widget.NewLabelWithStyle(
			xText,
			fyne.TextAlignCenter, 
			fyne.TextStyle{}))

		vBoxEYS.Add(widget.NewLabelWithStyle(
			yText,
			fyne.TextAlignCenter, 
			fyne.TextStyle{}))

		vBoxEZS.Add(widget.NewLabelWithStyle(
			zText,
			fyne.TextAlignCenter, 
			fyne.TextStyle{}))
	}

	hBoxE := container.NewHBox(
		vBoxEX,
		vBoxEY,
		widget.NewSeparator(),
		vBoxEXS,
		vBoxEYS,
		vBoxEZS,
	)

	wEuler.SetContent(hBoxE)

	// Set a button to show Euler's window
	btnEuler := widget.NewButton("Euler", func() {
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

	// Calculate the Runge-Kutte's solution
	x, y = RungeKutte()

	// Fill in the answers
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

	// Calculate the Runge-Kutte's solution for the system
	x, y, z = RungeKutteSystem()

	// Fill in the answers
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

	hBoxRK := container.NewHBox(
		vBoxRKX,
		vBoxRKY,
		widget.NewSeparator(),
		vBoxRKXS,
		vBoxRKYS,
		vBoxRKZS,
	)

	wRungeKutte.SetContent(hBoxRK)
	
	// Set a button to calculate and fill the Runge-Kutte window
	btnRungeKutta := widget.NewButton("Runge-Kutte", func() {
		wRungeKutte.Show()
	})

	// Set a window for the Milne's solution
	wMilne := a.NewWindow("Milne")

	wMilne.SetCloseIntercept(func() {
		wMilne.Hide()
	})

	// Set a layout for the Milne window
	vBoxMX := container.NewVBox()
	vBoxMY := container.NewVBox()

	vBoxMX.Add(widget.NewLabel("One equation:"))
	vBoxMY.Add(widget.NewLabel(""))

	vBoxMXS := container.NewVBox()
	vBoxMYS := container.NewVBox()
	vBoxMZS := container.NewVBox()

	vBoxMXS.Add(widget.NewLabel("System of equations:"))
	vBoxMYS.Add(widget.NewLabel(""))
	vBoxMZS.Add(widget.NewLabel(""))

	// Calculate the Milne's solution
	x, y = Milne()

	// Fill in the answers
	for i := -1; i < len(x); i++ {
		xText = "X"
		yText = "Y"

		if (i > -1){
			xText = fmt.Sprintf("%.2f", x[i])
			yText = fmt.Sprintf("%.3f", y[i])
		}
		
		vBoxMX.Add(widget.NewLabelWithStyle(
			xText,
			fyne.TextAlignCenter, 
			fyne.TextStyle{}))

		vBoxMY.Add(widget.NewLabelWithStyle(
			yText,
			fyne.TextAlignCenter, 
			fyne.TextStyle{}))
	}

	// Calculate the Milne's solution for the system
	x, y, z = MilneSystem()

	// Fill in the answers
	for i := -1; i < len(x); i++ {
		xText = "X"
		yText = "Y"
		zText = "Z"

		if (i > -1){
			xText = fmt.Sprintf("%.2f", x[i])
			yText = fmt.Sprintf("%.3f", y[i])
			zText = fmt.Sprintf("%.3f", z[i])
		}
		
		vBoxMXS.Add(widget.NewLabelWithStyle(
			xText,
			fyne.TextAlignCenter, 
			fyne.TextStyle{}))

		vBoxMYS.Add(widget.NewLabelWithStyle(
			yText,
			fyne.TextAlignCenter, 
			fyne.TextStyle{}))

		vBoxMZS.Add(widget.NewLabelWithStyle(
			zText,
			fyne.TextAlignCenter, 
			fyne.TextStyle{}))
	}

	hBoxM := container.NewHBox(
		vBoxMX,
		vBoxMY,
		widget.NewSeparator(),
		vBoxMXS,
		vBoxMYS,
		vBoxMZS,
	)

	wMilne.SetContent(hBoxM)

	// Set a button to calculate and fill the Miline window
	btnMilne := widget.NewButton("Milne", func() {
		wMilne.Show()
	})

	// Set a layout for the main window
	vBoxMain := container.NewVBox(
		layout.NewSpacer(),
		btnEuler,
		layout.NewSpacer(),
		btnRungeKutta,
		layout.NewSpacer(),
		// Adams
		btnMilne,
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
