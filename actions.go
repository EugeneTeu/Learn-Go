package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func testPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to test page\n")
	for _, rocket := range rockets {
		fmt.Fprintf(w, "%v\n", rocket)
	}
}

func retriveRockets(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllRows")
	json.NewEncoder(w).Encode(rockets)
}

func retrieveSingleRocket(w http.ResponseWriter, r *http.Request) {
	variables := mux.Vars(r)
	key := variables["id"]
	//fmt.Fprintf(w, "key: " + key);
	for _, row := range rockets {
		if row.ID == key {
			json.NewEncoder(w).Encode(row)
		}
	}
}

func createSingleRocket(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var rocket Rocket
	//fmt.Fprintf(w, "%+v", string(reqBody))
	json.Unmarshal(reqBody, &rocket)

	rockets = append(rockets, rocket)
	json.NewEncoder(w).Encode(rockets)
	//fmt.Fprintf(w, "%+v", string(reqBody))
}

func updateSingleRocket(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: update single Rocket")
	id := mux.Vars(r)["id"]
	reqBody, _ := ioutil.ReadAll(r.Body)
	var newRocket Rocket
	json.Unmarshal(reqBody, &newRocket)

	for index, rocket := range rockets {
		if rocket.ID == id {
			rocket.RocketName = newRocket.RocketName
			rockets[index] = rocket
		}
	}

}

func deleteSingleRocket(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	for index, row := range rockets {
		if row.ID == id {
			rockets = append(rockets[:index], rockets[index+1:]...)
		}
	}
}
