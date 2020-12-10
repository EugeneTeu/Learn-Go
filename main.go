package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Row struct
type Row struct {
	ID	 string `json:"Id"`;
	Name string `json:"Name"`;
}
var arr []Row;

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the home page");
	fmt.Printf("endpoint: home page");

}

func retriveRows(w http.ResponseWriter, r *http.Request) {
	 fmt.Println("Endpoint Hit: returnAllRows")
    json.NewEncoder(w).Encode(arr);
}

func retriveSingleRow(w http.ResponseWriter, r *http.Request) {
	variables := mux.Vars(r);
	key := variables["id"];
	//fmt.Fprintf(w, "key: " + key);
	for _, row := range arr {
		if row.ID == key {
			json.NewEncoder(w).Encode(row);
		}
	}
}

func createSingleRow(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	fmt.Fprintf(w, "%+v", string(reqBody))
}

func handleRequests() {
	http.HandleFunc("/", homePage);
	http.HandleFunc("/row", retriveRows);
	log.Fatal(http.ListenAndServe(":10000", nil));
}

func handleRequestsWithRouter() {
	myRouter := mux.NewRouter().StrictSlash(true);
	myRouter.HandleFunc("/", homePage);
	myRouter.HandleFunc("/all", retriveRows);
	myRouter.HandleFunc("/row", createSingleRow).Methods("POST")
	myRouter.HandleFunc("/row/{id}", retriveSingleRow);
	log.Fatal((http.ListenAndServe(":10000", myRouter)));
}

func main() {
	arr = []Row{
		{ID: "1" , Name: "alpha"},
		{ID: "2" , Name: "beta"},
	}
	handleRequestsWithRouter();
}