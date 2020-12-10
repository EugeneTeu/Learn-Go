package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Row struct {
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

func handleRequests() {
	http.HandleFunc("/", homePage);
	http.HandleFunc("/row", retriveRows);
	log.Fatal(http.ListenAndServe(":10000", nil));
}

func main() {
	arr = []Row{
		{Name: "alpha"},
		{Name: "beta"},
	}
	handleRequests();
}