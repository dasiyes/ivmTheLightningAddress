package main

import (
	"fmt"
	"log"
	"os"

	"github/dasiyes/ivmtla/internal/server"
	"github/dasiyes/ivmtla/internal/server/router"
)

func main() {
	fmt.Println("Hello, Lightning!")

	// Init a new HTTP server instance
	httpServer := server.NewInstance()
	hdlr := router.NewHandler(log.New(os.Stderr, "[ivmtla] ", log.LstdFlags))
	errs := make(chan error, 2)
	go func() {
		addr := ":8181"
		log.Printf("[main] ...starting ivmtla instance at %s...", addr)
		errs <- httpServer.Start(addr, hdlr)
	}()

	log.Printf("ivmtla http server terminated! %v", <-errs)
}
