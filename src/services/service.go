package services

import (
	"net/http"
	"regexp"

	"github.com/sideroff/go-book-rent-api/src/config"
)

// Middleware - a function that is a handler and is used to setup data on the request
// eg. - authentication
type Middleware func(rw http.ResponseWriter, r *http.Request, next Middleware)

// Service - the default service interface
type Service struct {
	URLPattern *regexp.Regexp
	// empty string = all methods, otherwise use http.Method*
	Method string
	RequiredRole config.Role
	Middlewares []Middleware // using nil is advised agains using empty array
	Execute func(rw http.ResponseWriter, r *http.Request)
}

