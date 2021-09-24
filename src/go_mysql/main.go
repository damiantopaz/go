package main

import (
	// Importing needed libries
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Prinitng available Drivesr on the project
	fmt.Println("Drivers:", sql.Drivers())

	// Conecting to database after creating
	// DataBase User: root
	// Database password ""
	// Database name test
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Fatal("unable to open connection to db")
	}
	defer db.Close()
	// SQL statement to query database
	result, err := db.Query("Select * from products")
	if err != nil {
		log.Fatal("Error when fetching product table rows:", err)
	}
	// closing conection to free result
	defer result.Close()
	//Looping through the whole data
	for result.Next() {
		// creating Variables to asign fetched data
		var (
			id       int
			name     string
			price    int
			quantity int
		)
		err = result.Scan(&id, &name, &price, &quantity)

		if err != nil {
			log.Fatal("Unable to parse row:", err)
		}
		// Printing out fetched assigned data
		fmt.Printf("ID: %d,Name: '%s', Price: %d, Quantity: %d\n", id, name, price, quantity)
	}

	// Fetching only single data
	// creating variable to accept fetch data
	var (
		id       int
		name     string
		price    int
		quantity int
	)
	// SQL Statment to query database
	err = db.QueryRow("Select * from products where id = 1").Scan(&id, &name, &price, &quantity)
	if err != nil {
		log.Fatal("Unable to Pass row:", err)
	}
	// printing out the fetch data
	fmt.Printf("ID: %d, Name: '%s', Price: %d, Quantity: %d\n", id, name, price, quantity)

	// Inserting Data to the table
	NewProducts := []struct {
		name     string
		price    int
		quantity int
	}{
		{"LightBulb", 200, 15},
		{"Mic", 14000, 6},
		{"Speker", 120000, 19},
		{"Wire Row", 14000, 38},
	}
	// Here I used prepared statemt
	stmt, err := db.Prepare("INSERT INTO products(name, price, qauntity) VALUES(?,?,?)")
	defer stmt.Close()
	if err != nil {
		log.Fatal("Unable to execute statement:", err)
	}
	for _, product := range NewProducts {
		_, err = stmt.Exec(product.name, product.price, product.quantity)
		if err != nil {
			log.Fatal("Unable to execute statement:", err)
		}
	}

}
