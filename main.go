package main

//TODO: structure project as per https://medium.com/hackernoon/golang-clean-archithecture-efd6d7c43047
import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// Rocket struct
type Rocket struct {
	ID         string `json:"Id"`
	RocketName string `json:"Rocket_Name"`
	Weight     int    `json:"Weight"`
}

var rockets []Rocket

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
