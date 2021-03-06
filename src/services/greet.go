package services

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/sideroff/go-book-rent-api/src/config"
)

// Greet - service handling greeting of users
var Greet = Service{
	URLPattern: regexp.MustCompile(`/`),
	Method: "",
	RequiredRole: config.Roles.Guest,
	Handler: func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "Hello world!")
	},
}
