package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"minta/internal/handler"
)

func main() {
	addr := os.Getenv("HTTP_ADDR")
	if addr == "" {
		addr = ":8080"
	}

	server := &http.Server{
		Addr:              addr,
		Handler:           handler.New(),
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Printf("Minta API listening on %s", addr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("run HTTP server: %v", err)
	}
}
