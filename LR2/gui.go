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
	w := a.NewWindow("LR1")
	w.SetMaster()

	// Create needed labels
	titleX := widget.NewLabel("X")
	titleX.Alignment = fyne.TextAlignCenter

	labelX := widget.NewLabel(fmt.Sprintf("%.2f, %.2f, %.2f, %.2f", X[0], X[1], X[2], X[3]))
	labelX.Alignment = fyne.TextAlignCenter

	titleDX := widget.NewLabel("dx")
	titleDX.Alignment = fyne.TextAlignCenter

	labelDX := widget.NewLabel("0.000, 0.000, 0.000, 0.000")
	labelDX.Alignment = fyne.TextAlignCenter

	titleDDX := widget.NewLabel("d2x")
	titleDDX.Alignment = fyne.TextAlignCenter

	labelDDX := widget.NewLabel("0.000, 0.000, 0.000, 0.000")
	labelDDX.Alignment = fyne.TextAlignCenter

	btn := widget.NewButton("Calculate", func() {
		dx, ddx := CalculateDerivatives()

		labelDX.SetText(fmt.Sprintf("%.3f, %.3f, %.3f, %.3f", dx[0], dx[1], dx[2], dx[3]))
		labelDDX.SetText(fmt.Sprintf("%.3f, %.3f, %.3f, %.3f", ddx[0], ddx[1], ddx[2], ddx[3]))
	})

	// Set a layout for the main window
	vBox := container.NewVBox(
		layout.NewSpacer(),
		titleX,
		labelX,
		layout.NewSpacer(),
		titleDX,
		labelDX,
		layout.NewSpacer(),
		titleDDX,
		labelDDX,
		layout.NewSpacer(),
		btn,
	)

	hBox := container.NewHBox(
		layout.NewSpacer(),
		vBox,
		layout.NewSpacer(),
	)

	w.SetContent(hBox)
	w.Resize(fyne.NewSize(300, 400))
	w.Show()
}
