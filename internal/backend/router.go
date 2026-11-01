package backend

import "net/http"

type Backend struct {
	server *http.Server
}

func NewBackend() *Backend {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /fortnite/api/cloudstorage/system", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("heya!"))
	})

	r := &Backend{
		server: &http.Server{
			Addr:    "localhost:8080",
			Handler: mux,
		},
	}

	return r
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
