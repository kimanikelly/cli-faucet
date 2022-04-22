package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func Database() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Failed to load the .env file %v", err)
	}

	username := os.Getenv("DB_USER_NAME")
	password := os.Getenv("DB_PASSWORD")

	connectionStr := fmt.Sprintf("%v:%v@tcp(127.0.0.1:3306)/accounts", username, password)

	db, err := sql.Open("mysql", connectionStr)

	if err != nil {
		log.Fatalf("Failed database connection %v", err)
	}

	defer db.Close()
}
