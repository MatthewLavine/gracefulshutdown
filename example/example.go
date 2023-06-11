package main

import (
	"context"
	"flag"
	"log"
	"net/http"

	"github.com/MatthewLavine/gracefulshutdown"
)

var (
	port = flag.String("port", ":8080", "Server port")
)

func main() {
	ctx := context.Background()

	httpServer := &http.Server{
		Addr: *port,
	}

	gracefulshutdown.AddShutdownHandler(func() error {
		log.Println("Shutting down HTTP server")
		return httpServer.Shutdown(ctx)
	})

	log.Printf("Server listening on port %s", *port)

	if err := httpServer.ListenAndServe(); err != nil {
		if err != http.ErrServerClosed {
			log.Fatalf("Could not listen on %s: %v\n", *port, err)
		}
	}

	gracefulshutdown.WaitForShutdown()
}
