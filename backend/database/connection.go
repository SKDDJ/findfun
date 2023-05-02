package database

import (
	"database/sql"
	"fmt"
	"log"

	// Import the appropriate database driver.
	_ "github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB
)

func ConnectDB() {
	// Replace the placeholders with your actual database credentials.
	username := "shiyiming"
	password := "shiyiming"
	host := "localhost"
	port := "3306"
	database := "findfun"

	// Create the database connection string.
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, database)

	// Connect to the database.
	var err error
	db, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}

	// Test the database connection.
	err = db.Ping()
	if err != nil {
		log.Fatal("Error pinging the database: ", err)
	}

	log.Println("Connected to the database.")
}

func GetDB() *sql.DB {
	return db
}
