package lr2

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
	w := a.NewWindow("LR2")
	w.Resize(fyne.NewSize(300, 100))
	w.SetMaster()

	// Create additional windows
	wDer := a.NewWindow("Derivative")
	wDer.Resize(fyne.NewSize(220, 300))
	wInt := a.NewWindow("Integral")
	wInt.Resize(fyne.NewSize(220, 300))

	// Handle hiding the additional windows
	// Instead of closing them
	wDer.SetCloseIntercept(func() {
		wDer.Hide()
	})

	wInt.SetCloseIntercept(func() {
		wInt.Hide()
	})

	// Create needed labels for derivative window
	titleXDer := widget.NewLabel("X")
	titleXDer.Alignment = fyne.TextAlignCenter

	labelXDer := widget.NewLabel(fmt.Sprintf("%.2f, %.2f, %.2f, %.2f", X[0], X[1], X[2], X[3]))
	labelXDer.Alignment = fyne.TextAlignCenter

	titleDX := widget.NewLabel("dx")
	titleDX.Alignment = fyne.TextAlignCenter

	labelDX := widget.NewLabel("0.000, 0.000, 0.000, 0.000")
	labelDX.Alignment = fyne.TextAlignCenter

	titleDDX := widget.NewLabel("d2x")
	titleDDX.Alignment = fyne.TextAlignCenter

	labelDDX := widget.NewLabel("0.000, 0.000, 0.000, 0.000")
	labelDDX.Alignment = fyne.TextAlignCenter

	// Set a layout for the derivative window
	vBoxDer := container.NewVBox(
		layout.NewSpacer(),
		titleXDer,
		labelXDer,
		layout.NewSpacer(),
		titleDX,
		labelDX,
		layout.NewSpacer(),
		titleDDX,
		labelDDX,
		layout.NewSpacer(),
	)

	hBoxDer := container.NewHBox(
		layout.NewSpacer(),
		vBoxDer,
		layout.NewSpacer(),
	)

	wDer.SetContent(hBoxDer)

	// Create needed labels for integral window
	titleXInt := widget.NewLabel("X")
	titleXInt.Alignment = fyne.TextAlignCenter

	labelXInt := widget.NewLabel(fmt.Sprintf("%.2f, %.2f, %.2f, %.2f, %.2f", X[0], X[1], X[2], X[3], X[4]))
	labelXInt.Alignment = fyne.TextAlignCenter

	titleNK := widget.NewLabel("Newton-Kotes:")
	titleNK.Alignment = fyne.TextAlignCenter

	labelNK := widget.NewLabel("0.000")
	labelNK.Alignment = fyne.TextAlignCenter

	titleS := widget.NewLabel("Squares:")
	titleS.Alignment = fyne.TextAlignCenter

	labelS := widget.NewLabel("0.000")
	labelS.Alignment = fyne.TextAlignCenter

	titleT := widget.NewLabel("Trapeze:")
	titleT.Alignment = fyne.TextAlignCenter

	labelT := widget.NewLabel("0.000")
	labelT.Alignment = fyne.TextAlignCenter

	titleSim := widget.NewLabel("Simpson:")
	titleSim.Alignment = fyne.TextAlignCenter

	labelSim := widget.NewLabel("0.000")
	labelSim.Alignment = fyne.TextAlignCenter


	// Set a layout for the integral window
	vBoxInt := container.NewVBox(
		layout.NewSpacer(),
		titleXInt,
		labelXInt,
		layout.NewSpacer(),
		titleNK,
		labelNK,
		layout.NewSpacer(),
		titleS,
		labelS,
		layout.NewSpacer(),
		titleT,
		labelT,
		layout.NewSpacer(),
		titleSim,
		labelSim,
		layout.NewSpacer(),
	)

	hBoxInt := container.NewHBox(
		layout.NewSpacer(),
		vBoxInt,
		layout.NewSpacer(),
	)

	wInt.SetContent(hBoxInt)

	btnDer := widget.NewButton("Calculate Derivative", func() {
		dx, ddx := CalculateDerivatives()

		labelDX.SetText(fmt.Sprintf("%.3f, %.3f, %.3f, %.3f", dx[0], dx[1], dx[2], dx[3]))
		labelDDX.SetText(fmt.Sprintf("%.3f, %.3f, %.3f, %.3f", ddx[0], ddx[1], ddx[2], ddx[3]))
	
		wDer.Show()
	})

	btnInt := widget.NewButton("Calculate Integral", func() {
		// Calculate integrals

		wInt.Show()
	})

	// Set a layout for the main window
	vBoxMain := container.NewVBox(
		layout.NewSpacer(),
		btnDer,
		layout.NewSpacer(),
		btnInt,
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
