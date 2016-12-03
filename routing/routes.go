package routing

import (
	"net/http"
)

//Route . . .
//struct used to define a route,
//Name: description of the route
//Method: http method
//Pattern: actual route endpoint
//HandlerFunc: function to handle the route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

//Res struct to be used when responding with json on request
type Res struct {
	Status  int8        `json:"status"`
	Message interface{} `json:"message"`
}

//Routes slice of routes that are registered with the mux router. All new routes can be defined here.
type Routes []Route

var routes = Routes{}
