package main

import (
	two "LRs/LR2"

	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()

	two.SetUi(a)

	a.Run()
}
