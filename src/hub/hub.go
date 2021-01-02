package hub

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sideroff/go-book-rent-api/src/config"
	"github.com/sideroff/go-book-rent-api/src/customerrors"
	"github.com/sideroff/go-book-rent-api/src/services"
)

// AccessRightList comment
type AccessRightList struct { 
    Guest int
    User int
    Admin int
}

// AccessRight - enum of app access rights
var AccessRight = &AccessRightList{ 
    Guest: 0,
    User: 1,
    Admin: 2,
}

// put any service u want made available here


// Hub - handles service initialization and execution
type Hub struct {
    logger *log.Logger
    serviceList []*services.Service
}

// NewHub - a constructor function for the hub
func NewHub(logger *log.Logger) *Hub {
    // a map of key: string, value: service
    // keys should be url regex
    var serviceList = []*services.Service{
        &services.Greet,
        &services.TestPost,
    }

    return &Hub{logger, serviceList}
}

func (hub *Hub) getService(url string, method string) (*services.Service, error) {
    for _, service := range hub.serviceList {
        if (service.URLPattern.MatchString(url) && (service.Method == method || method == "")) {
            return service, nil
        }
    }

    return nil, customerrors.ErrorServiceNotFound
}

// ExecuteService - executes the service specified by url
// TODO: should be a goroutine
func (hub *Hub) ExecuteService(url string, rw http.ResponseWriter, r *http.Request) {
    var service, err = hub.getService(url, r.Method)
    var userRole = config.Roles.Guest

    if (err == customerrors.ErrorServiceNotFound) {
        hub.logger.Print(err.Error())
        fmt.Fprint(rw, err.Error())

        return
    }

    // unrecognised error - log it and return 500 to user
    if (err != nil) {
        hub.logger.Print(err.Error())
        rw.WriteHeader(http.StatusInternalServerError)
        fmt.Fprint(rw, http.StatusText(http.StatusInternalServerError))

        return
    }

    if (service.RequiredRole > userRole) {
        msg := http.StatusText(http.StatusUnauthorized)
        hub.logger.Print(msg)
        fmt.Fprint(rw, msg)

        return
    }

    service.Handler(rw, r)
}