package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	server "task2/realization/server"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/version", server.VersionHandler)
	mux.HandleFunc("/decode", server.DecodeHandler)
	mux.HandleFunc("/hard-op", server.HardOpHandler)

	srv := &http.Server{
		Addr:    ":8082",
		Handler: mux,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()
	log.Println("Server is running on http://localhost:8082")

	<-stop
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown failed: %v", err)
	}
	log.Println("Server gracefully stopped")
}
