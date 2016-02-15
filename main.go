package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "demo:demo@(127.0.0.1:3306)/hello")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Printf("Ping DB error: %v\n", err)
	}

	var (
		name string
	)
	rows, err := db.Query("select name from hi")
	if err != nil {
		fmt.Printf("Query Err: %v\n", err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&name)
		if err != nil {
			fmt.Printf("Scan Err: %v\n", err)
		}
		log.Printf("Hi %s\n", name)
	}
	err = rows.Err()
	if err != nil {
		log.Printf("Rows err %s\n", err)
	}
}
