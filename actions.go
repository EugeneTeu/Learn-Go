package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func testPageAction(w http.ResponseWriter) {
	fmt.Fprintf(w, "Welcome to test page\n")
	for _, rocket := range rockets {
		fmt.Fprintf(w, "%v\n", rocket)
	}
}

func retrieveRocketAction(w http.ResponseWriter) {
	rockets := []Rocket{}
	rows, err := db.Query("SELECT * FROM Rocket")
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
	}
	var (
		id             string
		rocket_name    string
		payload_weight int
		rocket_type    string
	)

	for rows.Next() {
		err := rows.Scan(&id, &rocket_name, &payload_weight, &rocket_type)
		rocket := Rocket{id, rocket_name, payload_weight, rocket_type}
		if err != nil {
			log.Println(err)
			w.WriteHeader(500)
		}
		rockets = append(rockets, rocket)
	}
	defer rows.Close()

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

	stmt, err := db.Prepare("INSERT INTO Rocket(rocket_name, payload_weight, rocket_type) VALUES(?, ? , ?)")
	defer stmt.Close()
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(500)
	}

	_, err = stmt.Exec(rocket.RocketName, rocket.PayloadWeight, rocket.RocketType)
	if err != nil {
		log.Printf("error with statement %v", err)
		w.WriteHeader(500)
	}

	/*
		//fmt.Fprintf(w, "%+v", string(reqBody))
		rockets = append(rockets, rocket)
		json.NewEncoder(w).Encode(rockets)
		//fmt.Fprintf(w, "%+v", string(reqBody))*/
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
