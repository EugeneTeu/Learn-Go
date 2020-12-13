package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	fmt.Println("Go and MySql tutorial")

	db, err := sql.Open("mysql", os.Getenv("dev"))
	if err != nil {
		panic(err.Error())
	}

	//fmt.Println(os.Environ())
	defer db.Close()
}
