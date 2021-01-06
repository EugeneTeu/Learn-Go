package main

//TODO: structure project as per https://medium.com/hackernoon/golang-clean-archithecture-efd6d7c43047

//TODO: Convert to mysql
//TODO: Add function signatures
import (
	"database/sql"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// Rocket struct
type Rocket struct {
	ID            string `json:"id"`
	RocketName    string `json:"rocket_name"`
	PayloadWeight int    `json:"payload_weight"`
	RocketType    string `json:"rocket_type"`
}

var rockets []Rocket

func errorHandler(err error, message string) {
	if err != nil {
		log.Println(message)
		panic(err.Error())
	}
}

var db *sql.DB

func main() {
	rockets = []Rocket{
		/*{ID: "1", RocketName: "alpha", PayloadWeight: 5, RocketType: "apollo"},
		{ID: "2", RocketName: "beta", PayloadWeight: 10, RocketType: "gemini"},*/
	}
	log.Println("Starting Server")
	err := godotenv.Load()
	errorHandler(err, "error with loading env variables")
	db, err = sql.Open("mysql", os.Getenv("dev"))
	errorHandler(err, "error with db connection")

	// run sql create table
	dropTable, err := ioutil.ReadFile("./sql/drop-table.sql")
	errorHandler(err, "error reading sql file")
	//log.Printf(string(query))
	db.Query(string(dropTable))
	createTable, err := ioutil.ReadFile("./sql/create-table.sql")
	//log.Printf(string(createTable))
	_, err = db.Query(string(createTable))
	errorHandler(err, "create table failed")
	defer db.Close()

	// init router singleton
	myRouter := Router()
	log.Printf("Running on port %v\n", os.Getenv("PORT"))
	err = http.ListenAndServe(os.Getenv("PORT"), myRouter)
	errorHandler(err, "Error with starting port")

}
