package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dysdimas/internal/handler"
	"github.com/dysdimas/internal/middleware"
)

func StartServer(handler *handler.FizzBuzzHandler) {
	mux := http.NewServeMux()
	mux.HandleFunc("/range-fizzbuzz", handler.RangeFizzBuzzHandler)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      middleware.LoggingMiddleware(mux),
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	// Graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		<-stop
		log.Println("Shutting down server...")
		if err := server.Shutdown(context.Background()); err != nil {
			log.Fatalf("Could not gracefully shutdown the server: %v\n", err)
		}
	}()

	log.Println("Starting server on :8080")
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("Could not listen on :8080: %v\n", err)
	}

	log.Println("Server stopped")
}
