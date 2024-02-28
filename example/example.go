package main

import (
	"context"
	"flag"
	"fmt"
	"html"
	"log"
	"net/http"
	"time"

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

	// Gracefully shut down the HTTP server.
	gracefulshutdown.AddShutdownHandler(func() error {
		log.Println("Shutting down HTTP server...")
		httpServer.Shutdown(ctx)
		log.Println("HTTP server shut down.")
		return nil
	})

	// Perform application specific cleanup before exiting.
	gracefulshutdown.AddShutdownHandler(func() error {
		log.Println("Running some cleanup routine...")
		time.Sleep(5 * time.Second)
		log.Println("Cleanup routine complete.")
		return nil
	})

	log.Printf("Server listening on port %s", *port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	if err := httpServer.ListenAndServe(); err != nil {
		if err != http.ErrServerClosed {
			log.Fatalf("Could not listen on %s: %v\n", *port, err)
		}
	}

	gracefulshutdown.WaitForShutdown()
}
