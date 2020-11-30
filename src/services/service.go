package services

import (
	"net/http"

	"github.com/sideroff/go-book-rent-api/src/config"
)

// Middleware - a function that is a handler and is used to setup data on the request
// eg. - authentication
type Middleware func(rw http.ResponseWriter, r *http.Request, next Middleware)

// URLPattern - the pattern each service is attached to
type URLPattern string

// Service - the default service interface
type Service struct {
	urlPattern URLPattern
	requiredRole config.Role
	middlewares []Middleware // using nil is advised agains using empty array
	handler func(rw http.ResponseWriter, r *http.Request)
}

