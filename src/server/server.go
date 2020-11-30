package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// Start starts the server
func Start() {
	l := log.New(os.Stdout, string(os.Getpid())+" ", log.LstdFlags)
	serveMux := http.NewServeMux()

	// since no specific routes are added
	// and "/" is a prefix it matches everything
	serveMux.HandleFunc("/", handleRequest)

	server := &http.Server{
		Addr:         ":3000",
		Handler:      serveMux,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	l.Println("Starting server on port 3000")
	// server runs but wont block this "thread"
	// sadly no way to log a "sever is listening at..." msg
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			l.Fatal("Encountered a problem while serving.")
		}
	}()

	// blocks
	listenForOSSignal(server, l)
}

func handleRequest(responseWriter http.ResponseWriter, request *http.Request) {
	// check hub if service exists by url match
	// call hub.execute()
	// auth if service requires auth
}

func listenForOSSignal(server *http.Server, l *log.Logger) {
	signalChannel := make(chan os.Signal)
	signal.Notify(signalChannel, os.Interrupt)
	signal.Notify(signalChannel, os.Kill)

	// blocks
	sig := <-signalChannel
	l.Println("Received an interrupt, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)

	server.Shutdown(tc)
}
