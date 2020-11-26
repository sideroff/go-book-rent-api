package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/sideroff/go-book-rent-api/src/server/handlers"
)

// Start starts the server
func Start() {
	l := log.New(os.Stdout, string(os.Getpid())+" ", log.LstdFlags)

	g := handlers.NewGreetings(l)

	serveMux := http.NewServeMux()

	serveMux.Handle("/", g)

	server := &http.Server{
		Addr:         ":3000",
		Handler:      serveMux,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	l.Println("Starting server on port 3000")
	// server runs but wont block this "thread"
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			l.Fatal("Encountered a problem while serving.")
		}
	}()

	signalChannel := make(chan os.Signal)
	signal.Notify(signalChannel, os.Interrupt)
	signal.Notify(signalChannel, os.Kill)

	// block thread listening for os signal
	sig := <-signalChannel
	l.Println("Received an interrupt, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)

	server.Shutdown(tc)
}
