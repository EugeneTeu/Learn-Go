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
	ID   string `json:"Id"`
	Name string `json:"Name"`
}

var arr []Row

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the home page")
	fmt.Printf("endpoint: home page")

}

func retriveRows(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllRows")
	json.NewEncoder(w).Encode(arr)
}

func retriveSingleRow(w http.ResponseWriter, r *http.Request) {
	variables := mux.Vars(r)
	key := variables["id"]
	//fmt.Fprintf(w, "key: " + key);
	for _, row := range arr {
		if row.ID == key {
			json.NewEncoder(w).Encode(row)
		}
	}
}

func createSingleRow(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var row Row
	//fmt.Fprintf(w, "%+v", string(reqBody))
	json.Unmarshal(reqBody, &row)

	arr = append(arr, row)
	json.NewEncoder(w).Encode(arr)
	//fmt.Fprintf(w, "%+v", string(reqBody))
}

func updateSingleRow(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: update single Row")
	id := mux.Vars(r)["id"]
	reqBody, _ := ioutil.ReadAll(r.Body)
	var newRow Row
	json.Unmarshal(reqBody, &newRow)

	for index, row := range arr {
		if row.ID == id {
			row.Name = newRow.Name
			arr[index] = row
		}
	}

}

func deleteSingleRow(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	for index, row := range arr {
		if row.ID == id {
			arr = append(arr[:index], arr[index+1:]...)
		}
	}
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/row", retriveRows)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func handleRequestsWithRouter() {
	myRouter := Router()
	log.Fatal((http.ListenAndServe(":10000", myRouter)))
}

func main() {
	arr = []Row{
		{ID: "1", Name: "alpha"},
		{ID: "2", Name: "beta"},
	}
	handleRequestsWithRouter()
}
