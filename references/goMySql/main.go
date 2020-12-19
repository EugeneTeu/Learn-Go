package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// Database name = GO_SQL
// Table: test
// CREATE TABLE test ( Num int, Name varchar(255));

func main() {
	err := godotenv.Load()
	fmt.Println("Go and MySql tutorial")
	//fmt.Println(os.Getenv("dev"))
	db, err := sql.Open("mysql", os.Getenv("dev"))
	if err != nil {
		panic(err.Error())
	}

	//fmt.Println(os.Environ())

	defer db.Close()
	insert, err := db.Query("INSERT INTO test VALUES (2,'TEST') ")
	if err != nil {
		//fmt.Println("inserrt error")
		panic(err.Error())
	}

	defer insert.Close()

}
