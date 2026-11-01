package server

import (
	"net/http"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Backend struct {
	server *http.Server
}

func NewBackend() *Backend {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /fortnite/api/cloudstorage/system", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("heya!"))
	})

	return &Backend{
		server: &http.Server{
			Addr:    "localhost:8080",
			Handler: mux,
		},
	}
}

func (b *Backend) Widget() fyne.CanvasObject {
	return container.NewVBox(
		widget.NewLabel("Running on port :8080"),
		widget.NewButton("Start Server", nil),
	)
}

func (b *Backend) ListenAndServe() error {
	if err := b.server.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

func (b *Backend) Shutdown() error {
	return nil
}
