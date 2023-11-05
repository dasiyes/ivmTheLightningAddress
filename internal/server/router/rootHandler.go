package router

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

type rootHandler struct {
	l *log.Logger
	// place any dependencies ...
}

func (rh *rootHandler) router() chi.Router {
	rtr := chi.NewRouter()

	// Route the static files
	// fileServer := http.FileServer(http.FS(ui.Files))
	// rtr.Handle("/static/*", fileServer)

	rtr.Route("/", func(r chi.Router) {
		r.Get("/", rh.welcome)
		// r.Post("/send", rh.home)
		// r.Get("/nostr", rh.nostr)
	})

	return rtr
}

func (rh *rootHandler) welcome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	_, _ = w.Write([]byte("HI! Welcome to Ivmanto's The Lightning Address service!"))
}
