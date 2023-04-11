package lr1

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
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
	w1.Resize(fyne.NewSize(500, 500))
	w2 := a.NewWindow("Polynom")
	w2.Resize(fyne.NewSize(500, 500))

	// Create needed labels
	titlePolynomial := widget.NewLabel("Polynomial at a:")
	titlePolynomial.Alignment = fyne.TextAlignCenter

	labelPolynomial := widget.NewLabel("0.0")
	labelPolynomial.Alignment = fyne.TextAlignCenter

	titleError := widget.NewLabel("Error of the polynomial:")
	titleError.Alignment = fyne.TextAlignCenter

	labelError := widget.NewLabel("0.0")
	labelError.Alignment = fyne.TextAlignCenter

	btnPolynom := widget.NewButton("Show polynomial", func() {
		val, err := CreatePolynomialGraph()
		
		labelPolynomial.SetText(fmt.Sprintf("%v", val))
		labelError.SetText(fmt.Sprintf("%v", err))

		img := canvas.NewImageFromFile("LR1/polynomial.png")
		
		w2.SetContent(img)

		w2.Show()
	})

	// Set a layout for the main window
	vBox := container.NewVBox(
		layout.NewSpacer(),
		titlePolynomial,
		labelPolynomial,
		layout.NewSpacer(),
		titleError,
		labelError,
		layout.NewSpacer(),
		btnPolynom,
		layout.NewSpacer(),
	)

	hBox := container.NewHBox(
		layout.NewSpacer(),
		vBox,
		layout.NewSpacer(),
	)

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
