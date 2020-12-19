package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

/*
//TODO: handle enums
type EndpointName string

const (
	homepage     EndpointName = "homePage"
	all          EndpointName = "all"
	rocket       EndpointName = "rocket"
	singleRocket EndpointName = "singleRocket"
)

func (name *EndpointName) getValue(b []byte) string {
	var result string
	json.Unmarshal(b, &result)
	endpointName := EndpointName(result)
	switch endpointName {
	case homepage, all, rocket, singleRocket:
		*name = endpointName
	}
	return result

}
*/

var endpoints = map[string]string{
	"homePage":     "/",
	"all":          "/all",
	"rocket":       "/rocket",
	"singleRocket": "/rocket/{id}",
}

func makeHandler(fn func(http.ResponseWriter, *http.Request), title string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Endpoint hit: %s\n", title)
		fn(w, r)
	}
}

// Router handles routing
func Router() *mux.Router {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc(endpoints["homePage"], makeHandler(homePage, endpoints["homePage"]))
	myRouter.HandleFunc(endpoints["all"], makeHandler(retriveRockets, endpoints["all"]))
	myRouter.HandleFunc(endpoints["rocket"], makeHandler(createSingleRocket, endpoints["rocket"])).Methods("POST")
	myRouter.HandleFunc(endpoints["singleRocket"], makeHandler(deleteSingleRocket, endpoints["singleRocket"])).Methods("DELETE")
	myRouter.HandleFunc(endpoints["singleRocket"], makeHandler(updateSingleRocket, endpoints["singleRocket"])).Methods("PUT")
	myRouter.HandleFunc(endpoints["singleRocket"], makeHandler(retrieveSingleRocket, endpoints["singleRocket"]))
	return myRouter
}
