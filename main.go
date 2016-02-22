package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const dbname = "mysqltest"
const tablename = "tabletest"

func main() {
	// Connect to server
	db, err := sql.Open("mysql", "demo:demo@(127.0.0.1:3306)/")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create new database dynamically
	fmt.Println("Creating Database")
	_, err = db.Exec("CREATE DATABASE " + dbname)
	if err != nil {
		log.Printf("Database %s already exists", dbname)
	}

	// Test DB connection
	err = db.Ping()
	if err != nil {
		fmt.Printf("Ping DB error: %v\n", err)
	}

	// Select table to use
	_, err = db.Exec("USE " + dbname)
	if err != nil {
		panic(err)
	}

	fmt.Println("Creating table Data")

	// Dynamically create table
	_, err = db.Exec("CREATE TABLE " + tablename + " ( name varchar(32), typeOf varchar(32) )")
	if err != nil {
		log.Print(err)
	}

	// Modifying Data
	fmt.Println("Modifying Data")
	stmt, err := db.Prepare("INSERT INTO " + tablename + "(name, typeOf) VALUES(?, ?)")
	if err != nil {
		log.Printf(err.Error())
	}

	fmt.Println("Inserting Data")
	res, err := stmt.Exec("Brian", "Survivor")
	if err != nil {
		log.Printf(err.Error())
	}

	id, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}
	fmt.Printf("ID: %d\n", id)

	fmt.Println("Querying Data")
	rows, err := db.Query("select name, typeOf from " + tablename)
	if err != nil {
		fmt.Printf("Query Err: %v\n", err)
	}
	defer rows.Close()
	for rows.Next() {
		var name string
		var typeOf string
		err := rows.Scan(&name, &typeOf)
		if err != nil {
			fmt.Printf("Scan Err: %v\n", err)
			panic(err)
		}
		log.Printf("Name: %s Type Of: %s\n", name, typeOf)
	}
	err = rows.Err()
	if err != nil {
		log.Printf("Rows err %s\n", err)
	}

	db.Prepare("DELETE from " + tablename + "where id=?")
	if err != nil {
		log.Printf("Delete rows err %s\n", err)
	}

	_, err = stmt.Exec(id)

	fmt.Println("Deleting rows")
	_, err = db.Exec("DELETE FROM " + tablename)
	if err != nil {
		log.Printf(err.Error())
	}
}
