package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// Creates the MySQL database connection
func Connection() *sql.DB {

	// Loads the environment variables from the .env file
	err := godotenv.Load()

	// If err does not equal nil(zero value) throw an err
	if err != nil {
		log.Fatalf("Failed to load the .env file %v", err)
	}

	// Returns the MySQL username stored in the .env file
	username := os.Getenv("DB_USER_NAME")

	// Returns the MySQL password stored in the .env file
	password := os.Getenv("DB_PASSWORD")

	// Builds the MySQL connection string
	connectionStr := fmt.Sprintf("%v:%v@tcp(127.0.0.1:3306)/FaucetDB", username, password)

	// Creates the MySQL database connection
	db, err := sql.Open("mysql", connectionStr)

	// If err does not equal nil(zero value) throw an error
	if err != nil {
		log.Fatalf("Failed database connection %v", err)
	}

	// Returns the MySQL connection instance
	return db
}
