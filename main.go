package main

import (
	three "LRs/LR3"

	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()

	three.SetUi(a)

	a.Run()
}
