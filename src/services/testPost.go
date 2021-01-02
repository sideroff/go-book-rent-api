package services

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/sideroff/go-book-rent-api/src/config"
)

// TestPost - service test
var TestPost = Service{
	URLPattern: regexp.MustCompile(`/testPost`),
	Method: "POST",
	RequiredRole: config.Roles.Guest,
	Handler: func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "Hello world from testPost!")
	},
}
