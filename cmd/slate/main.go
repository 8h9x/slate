package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/8h9x/slate/internal/backend"
	"github.com/8h9x/slate/internal/egl"
)

type Application struct {
	frontend fyne.App
	backend  *backend.Backend
}

func NewApplication() *Application {
	return &Application{
		frontend: app.New(),
		backend:  backend.NewBackend(),
	}
}

func (c *Application) Run() {
	w := c.frontend.NewWindow("Slate Launcher")

	epicInstallPath, err := egl.LocateInstallation("Fortnite")
	if err != nil {
		panic(err)
	}

	w.Resize(fyne.NewSize(640, 300))

	combo := widget.NewSelectEntry([]string{"Red", "Green", epicInstallPath})
	combo.PlaceHolder = "Game Path"
	combo.OnChanged = func(value string) {
		log.Println("Colour is", value)
	}

	backendBtn := widget.NewButton("Start Backend", func() {})
	backendBtn.OnTapped = func() {
		go func() {
			fyne.Do(func() {
				backendBtn.SetText("Stop Backend")
				backendBtn.Refresh()
			})

			err = c.backend.ListenAndServe()
			if err != nil {
				fyne.Do(func() {
					backendBtn.SetText("Start Backend")
					backendBtn.Refresh()
				})
				log.Fatal(err)
			}
		}()
	}

	content := container.NewVBox(combo, widget.NewButton("Launch", func() {
		log.Println("Content was:", combo.Text)
	}), backendBtn)

	w.SetContent(content)
	w.Show()

	c.frontend.Run()
}

func main() {
	a := NewApplication()

	log.Println("hey, slate!")

	a.Run()
}
