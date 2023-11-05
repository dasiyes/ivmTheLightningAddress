package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"
)

type Instance struct {
	lg         *log.Logger
	httpServer *http.Server
}

func NewInstance() *Instance {

	log.Println("...initiating new Instance of HTTP server...")
	s := &Instance{
		lg:         log.New(os.Stdout, "ivmtla: ", log.LstdFlags),
		httpServer: &http.Server{},
	}

	return s
}

func (s *Instance) Start(addr string, endp http.Handler) error {

	s.httpServer = &http.Server{
		Addr:    addr,
		Handler: endp,
	}

	err := s.httpServer.ListenAndServe() // Blocks!
	if err != http.ErrServerClosed {
		log.Printf("Http Server stopped unexpected: %v", err)
		s.Shutdown()
	} else {
		log.Printf("Critical: Http Server stopped: %v", err)
	}
	return err
}

func (s *Instance) Shutdown() {
	if s.httpServer != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		err := s.httpServer.Shutdown(ctx)
		if err != nil {
			s.lg.Printf("Failed to shutdown http server gracefully: %v", err)
		} else {
			s.httpServer = nil
		}
	}
}
