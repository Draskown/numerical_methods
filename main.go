package main

import (
	one "LRs/LR1"

	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()

	one.SetUi(a)

	a.Run()
}