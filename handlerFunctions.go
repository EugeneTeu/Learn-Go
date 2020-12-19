package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func testPage(w http.ResponseWriter, r *http.Request) {
	testPageAction(w)
}

func retrieveRockets(w http.ResponseWriter, r *http.Request) {
	retrieveRocketAction(w)
}

func retrieveSingleRocket(w http.ResponseWriter, r *http.Request) {
	variables := mux.Vars(r)
	key := variables["id"]
	retrieveSingleRocketAction(w, key)
}

func createSingleRocket(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var rocket Rocket
	json.Unmarshal(reqBody, &rocket)
	createSingleRocketAction(w, rocket)
}

func updateSingleRocket(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: update single Rocket")
	id := mux.Vars(r)["id"]
	reqBody, _ := ioutil.ReadAll(r.Body)
	var newRocket Rocket
	json.Unmarshal(reqBody, &newRocket)
	updateSingleRocketAction(id, newRocket)
}

func deleteSingleRocket(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	deleteSingleRocketAction(id)
}
