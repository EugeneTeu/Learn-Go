package main

import "github.com/gorilla/mux"

// exports routing
func Router() *mux.Router {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/all", retriveRows)
	myRouter.HandleFunc("/row", createSingleRow).Methods("POST")
	myRouter.HandleFunc("/row/{id}", deleteSingleRow).Methods("DELETE")
	myRouter.HandleFunc("/row/{id}", updateSingleRow).Methods("PUT")
	myRouter.HandleFunc("/row/{id}", retriveSingleRow)
	return myRouter
}
