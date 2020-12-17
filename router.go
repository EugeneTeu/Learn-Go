package main

import "github.com/gorilla/mux"

// Router handles routing
func Router() *mux.Router {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/all", retriveRockets)
	myRouter.HandleFunc("/rocket", createSingleRocket).Methods("POST")
	myRouter.HandleFunc("/rocket/{id}", deleteSingleRocket).Methods("DELETE")
	myRouter.HandleFunc("/rocket/{id}", updateSingleRocket).Methods("PUT")
	myRouter.HandleFunc("/rocket/{id}", retriveSingleRocket)
	return myRouter
}
