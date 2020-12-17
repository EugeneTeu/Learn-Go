package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

// Rocket struct
type Rocket struct {
	ID         string `json:"Id"`
	RocketName string `json:"Rocket_Name"`
	Weight     int    `json:"Weight"`
}

var rockets []Rocket

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the home page\n")
	for _, rocket := range rockets {
		fmt.Fprintf(w, "%v\n", rocket)
	}
	fmt.Printf("endpoint: home page")
}

func retriveRockets(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllRows")
	json.NewEncoder(w).Encode(rockets)
}

func retriveSingleRocket(w http.ResponseWriter, r *http.Request) {
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

func errorHandler(err error, message string) {
	if err != nil {
		log.Println(message)
		panic(err.Error)
	}
}

func main() {
	rockets = []Rocket{
		{ID: "1", RocketName: "alpha", Weight: 5},
		{ID: "2", RocketName: "beta", Weight: 10},
	}
	log.Println("Starting Server")
	err := godotenv.Load()
	errorHandler(err, "error with loading env variables")

	// init router singleton
	myRouter := Router()
	log.Printf("Running on port %v\n", os.Getenv("PORT"))
	err = http.ListenAndServe(os.Getenv("PORT"), myRouter)
	errorHandler(err, "Error with starting port")

}
