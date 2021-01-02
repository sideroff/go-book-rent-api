package main

import (
	"log"
	"os"
	"strconv"

	"github.com/sideroff/go-book-rent-api/src/bookserver"
	"github.com/sideroff/go-book-rent-api/src/hub"
)

func main() {
	l := log.New(os.Stdout, strconv.Itoa(os.Getpid()) + ": ", log.LstdFlags)

	serviceHub := hub.NewHub(l)

	server := bookserver.NewBookServer(l, serviceHub)

	server.Start()
}
