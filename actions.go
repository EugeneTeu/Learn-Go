package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func testPageAction(w http.ResponseWriter) {
	fmt.Fprintf(w, "Welcome to test page\n")
	for _, rocket := range rockets {
		fmt.Fprintf(w, "%v\n", rocket)
	}
}

func retrieveRocketAction(w http.ResponseWriter) {
	json.NewEncoder(w).Encode(rockets)
}

func retrieveSingleRocketAction(w http.ResponseWriter, key string) {
	for _, row := range rockets {
		if row.ID == key {
			json.NewEncoder(w).Encode(row)
		}
	}
}

func createSingleRocketAction(w http.ResponseWriter, rocket Rocket) {
	//fmt.Fprintf(w, "%+v", string(reqBody))
	rockets = append(rockets, rocket)
	json.NewEncoder(w).Encode(rockets)
	//fmt.Fprintf(w, "%+v", string(reqBody))
}

func updateSingleRocketAction(id string, newRocket Rocket) {
	for index, rocket := range rockets {
		if rocket.ID == id {
			rocket.RocketName = newRocket.RocketName
			rockets[index] = rocket
		}
	}
}

func deleteSingleRocketAction(id string) {
	for index, row := range rockets {
		if row.ID == id {
			rockets = append(rockets[:index], rockets[index+1:]...)
		}
	}
}
