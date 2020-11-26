package handlers

import (
	"log"
	"net/http"
)

type Greetings struct {
	l *log.Logger
}

// accepts a pointer to a logger, returns a pointer to a Greetings obj
func NewGreetings(l *log.Logger) *Greetings {
	// return the pointer of this obj
	return &Greetings{l: l}
}

func (g *Greetings) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	g.l.Println("received a request")

	rw.Write([]byte("Hello World!"))
}
