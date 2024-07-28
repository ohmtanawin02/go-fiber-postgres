package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	databaseName := os.Getenv("DB_NAME")
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, username, password, databaseName)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return fmt.Errorf("error opening database connection: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return fmt.Errorf("error pinging database: %w", err)
	}

	DB = db
	log.Println("Successfully connected to the database")
	return nil
}

// func createTables() error {
// 	createUserTable := `
// 	CREATE TABLE IF NOT EXISTS users (
// 		id SERIAL PRIMARY KEY,
// 		username VARCHAR(50) UNIQUE NOT NULL,
// 		password VARCHAR(100) NOT NULL
// 	);`

// 	_, err := DB.Exec(createUserTable)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
