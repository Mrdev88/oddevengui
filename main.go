package main

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Oddevengui")
	w.Resize(fyne.NewSize(300, 200))

	label := widget.NewLabel("Oddevengui")
	label.TextStyle = fyne.TextStyle{Bold: true}
	label.Alignment = fyne.TextAlignCenter

	label2 := widget.NewLabel("By Mahdi Ruiz")
	label2.Alignment = fyne.TextAlignCenter

	input := widget.NewEntry()
	input.SetPlaceHolder("Number to check if even or odd")
	// Center alignment for Entry is not supported directly; wrap in a centered container if needed

	result := widget.NewLabel("")

	checkbutton := widget.NewButton("Check", func() {
		number := input.Text
		if number == "" {
			result.SetText("Please enter a number")
			return
		}
		num, err := strconv.Atoi(number)
		if err != nil {
			result.SetText("Please enter a vaild number")
			return
		}
		if num%2 != 0 {
			result.SetText(number + " is odd")
		} else {
			result.SetText(number + " is even")
		}
	})

	content := container.NewVBox(
		label,
		label2,
		input,
		checkbutton,
		result,
	)

	w.SetContent(content)
	w.ShowAndRun()
}
