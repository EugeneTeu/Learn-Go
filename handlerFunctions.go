package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

/*

type Rocket struct {
	ID            string `json:"id"`
	RocketName    string `json:"rocket_name"`
	PayloadWeight int    `json:"payload_weight"`
	RocketType    string `json:"rocket_type"`
}

*/

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
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Body read error, %v", err)
		w.WriteHeader(500) // Return 500 Internal Server Error.
		return
	}

	var rocket Rocket
	err = json.Unmarshal(reqBody, &rocket)
	if err != nil {
		log.Printf("Body parse error, %v", err)
		w.WriteHeader(400) // Return 400 Bad Request.
		return
	}
	createSingleRocketAction(w, rocket)
}

func updateSingleRocket(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: update single Rocket")
	id := mux.Vars(r)["id"]
	reqBody, _ := ioutil.ReadAll(r.Body)
	var newRocket Rocket
	json.Unmarshal(reqBody, &newRocket)
	updateSingleRocketAction(id, newRocket)
	w.WriteHeader(200)
}

func deleteSingleRocket(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	deleteSingleRocketAction(id)
}
