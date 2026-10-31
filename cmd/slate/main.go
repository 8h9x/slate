package main

import (
	"fyne.io/fyne/v2"
	"log"
	"net/http"
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	_ = http.NewServeMux()

	log.Println("hey, slate!")

	a := app.New()
	w := a.NewWindow("Update Time")

	w.Resize(fyne.Size{
		Width:  640,
		Height: 300,
	})

	message := widget.NewLabel("Welcome")
	button := widget.NewButton("Update", func() {
		formatted := time.Now().Format("Time: 03:04:05")
		message.SetText(formatted)
	})

	w.SetContent(container.NewVBox(message, button))
	w.ShowAndRun()
}
