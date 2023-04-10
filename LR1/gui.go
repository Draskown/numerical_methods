package lr1

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// Sets a GUI for passed application
func SetUi(a fyne.App) {
	// Create the main window
	w := a.NewWindow("LR1")
	w.SetMaster()

	// Create additional windows
	w1 := a.NewWindow("Approximation")
	w2 := a.NewWindow("Polynom")

	titleApproximation := widget.NewLabel("Approximation at a:")
	titleApproximation.Alignment = fyne.TextAlignCenter

	labelAppoximation := widget.NewLabel("0.0")
	labelAppoximation.Alignment = fyne.TextAlignCenter

	titlePolynom := widget.NewLabel("Polynom at a:")
	titlePolynom.Alignment = fyne.TextAlignCenter

	labelPolynom := widget.NewLabel("0.0")
	labelPolynom.Alignment = fyne.TextAlignCenter

	// Create a button for calculation
	btnCalculate := widget.NewButton("Calculate", func() {
		// Calculate
	})

	// Create a button to show Approximation window
	btnApproximation := widget.NewButton("Show approximation", func() {
		CreateGraph()
		w1.Show()
	})

	btnPolynom := widget.NewButton("Show polynom", func() {
		// Create graph
		w2.Show()
	})

	// Set a layout for the main window
	vBox := container.New(layout.NewVBoxLayout(),
		layout.NewSpacer(),
		btnCalculate,
		layout.NewSpacer(),
		titleApproximation,
		labelAppoximation,
		layout.NewSpacer(),
		titlePolynom,
		labelPolynom,
		layout.NewSpacer(),
		btnApproximation,
		layout.NewSpacer(),
		btnPolynom,
		layout.NewSpacer(),
	)

	hBox := container.New(layout.NewHBoxLayout(),
		layout.NewSpacer(),
		vBox,
		layout.NewSpacer())

	// Handle hiding the additional windows
	// Instead of closing them
	w1.SetCloseIntercept(func() {
		w1.Hide()
	})

	w2.SetCloseIntercept(func() {
		w2.Hide()
	})

	w.SetContent(hBox)
	w.Resize(fyne.NewSize(300, 400))
	w.Show()
}
