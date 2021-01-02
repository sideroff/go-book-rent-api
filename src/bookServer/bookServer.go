package bookserver

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/sideroff/go-book-rent-api/src/hub"
)

// BookServer the book server
type BookServer struct {
	logger *log.Logger
	serviceHub *hub.Hub
}

// NewBookServer constructor fn for the server
func NewBookServer(logger *log.Logger, serviceHub *hub.Hub) *BookServer {
	return &BookServer{
		logger,
		serviceHub,
	}
}

// Start starts the server
func (bookServer *BookServer) Start() {
	serveMux := http.NewServeMux()
	// since no specific routes are added
	// and "/" is a prefix it matches everything
	serveMux.HandleFunc("/", bookServer.handleRequest)

	server := &http.Server{
		Addr:         ":3000",
		Handler:      serveMux,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	bookServer.logger.Println("Starting server on port: 3000")
	// server runs but wont block this "thread"
	// sadly no way to log a "sever is listening at..." msg
	go func() {
		err := server.ListenAndServe()

		if err != nil {
			bookServer.logger.Fatalf("Server shutting down because of %s", err.Error())
		}
	}()

	// blocks
	listenForOSSignal(server, bookServer.logger)
}

func (bookServer *BookServer) handleRequest(responseWriter http.ResponseWriter, request *http.Request) {
	bookServer.logger.Printf("received request %s %s ", request.Method, request.URL.Path)

	bookServer.serviceHub.ExecuteService(request.URL.Path, responseWriter, request)
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
