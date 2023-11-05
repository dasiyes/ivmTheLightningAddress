package router

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
)

// Constructing web application depenedencies in the format of handler
type srvHandler struct {
	l *log.Logger
	// ... add other dependencies here
}

func (h *srvHandler) router() chi.Router {

	rtr := chi.NewRouter()

	// Building middleware chain
	rtr.Use(accessControl)

	// Route the home calls
	rtr.Route("/", func(r chi.Router) {
		lgr := log.New(os.Stdout, "[http-srv-root-hdlr] ", log.LstdFlags)
		rh := rootHandler{lgr}
		r.Mount("/", rh.router())
	})

	return rtr
}

// Handler to manage endpoints
func NewHandler(l *log.Logger) http.Handler {

	log.Printf("...initializing router (new Handler) ...")
	e := srvHandler{
		l: l,
	}

	return e.router()
}
