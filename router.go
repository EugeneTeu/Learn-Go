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
const (
	homePage     string = "homePage"
	all          string = "all"
	rocket       string = "rocket"
	singleRocket string = "singleRocket"
)

var endpoints = map[string]string{
	homePage:     "/",
	all:          "/all",
	rocket:       "/rocket",
	singleRocket: "/rocket/{id}",
}

// func wrapper
func makeHandler(fn http.HandlerFunc, title string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Endpoint hit: %s\n", title)
		fn(w, r)
	}
}

// middle ware
func loggingMiddleware(fn http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Request uri: %v\n", r.RequestURI)
		fn.ServeHTTP(w, r)
	})
}

// Router handles routing
func Router() *mux.Router {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.Use(loggingMiddleware)
	myRouter.HandleFunc(endpoints[homePage], makeHandler(testPage, endpoints[homePage]))
	myRouter.HandleFunc(endpoints[all], makeHandler(retrieveRockets, endpoints[all]))
	myRouter.HandleFunc(endpoints[rocket], makeHandler(createSingleRocket, endpoints["rocket"])).Methods("POST")
	myRouter.HandleFunc(endpoints[singleRocket], makeHandler(deleteSingleRocket, endpoints[singleRocket])).Methods("DELETE")
	myRouter.HandleFunc(endpoints[singleRocket], makeHandler(updateSingleRocket, endpoints[singleRocket])).Methods("PUT")
	myRouter.HandleFunc(endpoints[singleRocket], makeHandler(retrieveSingleRocket, endpoints[singleRocket]))
	return myRouter
}
