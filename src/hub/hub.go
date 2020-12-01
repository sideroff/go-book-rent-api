package hub

import (
    "log"
	"net/http"

	"github.com/sideroff/go-book-rent-api/src/services"
)

// AccessRightCode - type of the access right
type AccessRightCode int

// AccessRightList - comment
type AccessRightList struct { 
    Guest AccessRightCode
    User AccessRightCode
    Admin AccessRightCode
}

// AccessRight - enum of app access rights
var AccessRight = &AccessRightList{ 
    Guest: 0,
    User: 1,
    Admin: 2,
}


// put any service u want made available here
var serviceList = []services.Service{
    services.Greet,
    // services....
}

// a map of key: string, value: service
// keys should be url regex
var servicesMap = make(map[string]*services.Service)

// Initialize - a function that starts up the hub and services
func Initialize(l *log.Logger) {
    
}

// ExecuteService - executes the service specified by url
// TODO: should be a goroutine
func ExecuteService(url string, rw http.ResponseWriter, r *http.Request) {
    if services.Greet.URLPattern.MatchString(url) {
        services.Greet.Execute(rw,r)
    }
}
