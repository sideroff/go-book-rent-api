package services

import (
	"fmt"
	"net/http"

	"github.com/sideroff/go-book-rent-api/src/config"
)

// Greet - service handling greeting of users
var Greet = &Service{
	urlPattern: "/",
	requiredRole: config.Roles.Guest,
	middlewares: nil, 
	handler: func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "Hello world!")
	},
}