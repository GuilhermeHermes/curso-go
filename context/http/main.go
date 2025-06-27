package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("handler started")
	defer log.Println("handler ended")

	select {
	case <-time.After(5 * time.Second):
		fmt.Fprintln(w, "Request processed")
	case <-ctx.Done():
		err := ctx.Err()
		http.Error(w, "Request cancelled: "+err.Error(), http.StatusRequestTimeout)
	}
}

func main() {
	http.HandleFunc("/", handler)

	server := &http.Server{
		Addr: ":8080",
	}

	go func() {
		log.Println("Starting server on :8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe(): %v", err)
		}
	}()

	// Graceful shutdown on interrupt signal
	stop := make(chan struct{})
	go func() {
		time.Sleep(30 * time.Second)
		log.Println("Shutting down server...")
		server.Shutdown(context.Background())
		close(stop)
	}()

	<-stop
	log.Println("Server stopped")
}
