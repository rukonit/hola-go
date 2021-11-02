package main

import (
	"fmt"
	"log"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)


func main() {

	// db, err:= sql.Open("mysql", "rukon:123456@tcp(127.0.0.1:3306)/product")

	if err != nil {
		log.Fatal("Unable to open connection to db")
	}

	defer db.Close()
	results, err := db.Query("select * from category")

	if err != nil {
		log.Fatal("Error occured when fetching data")
	}
	defer results.Close()

	for results.Next() {
		var (id int 
		category string
		)
		err = results.Scan(&id, &category)
		if err != nil {
			log.Fatal("Unable to parse row", err)
		}
		fmt.Printf("ID: %d, Category: %s\n", id, category)
	}
	
}