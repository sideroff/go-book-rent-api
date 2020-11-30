package hub

import (
	"log"

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

serviceList := []*services.Service{services.Greet}

// Initialize - a function that starts up the hub and services
func Initialize(l *log.Logger) {

}

// DoesServiceExist - checks if a service that expects this specific url exists
func DoesServiceExist(url string) bool {

	return false
}

// ExecuteService - executes the service specified by url
// TODO: should be a goroutine
func ExecuteService(url string) {

}
