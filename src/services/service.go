package services

import (
	"net/http"
	"regexp"
)

// Service - the default service interface
type Service struct {
	URLPattern *regexp.Regexp
	// empty string = all methods, otherwise use http.Method*
	Method string
	RequiredRole int
	Handler func(rw http.ResponseWriter, r *http.Request)
}

