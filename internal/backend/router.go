package backend

import "net/http"

type Router struct {
	*http.ServeMux
}
