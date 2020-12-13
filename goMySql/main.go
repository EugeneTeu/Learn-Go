package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Go and MySql tutorial")

	db, err := sql.Open("mysql", "eugeneteu:Acjc2014@tcp(127.0.0.1:3306)/G0_SQL")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
}
