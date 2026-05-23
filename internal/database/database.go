package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

var DB *pgx.Conn 

func Connect() {
	// load environment variables

	err := godotenv.Load() 

	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dbUrl := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
	)

	// connect to PostgreSQL
	conn, err := pgx.Connect(context.Background(), dbUrl)

	if err != nil {
		log.Fatalf("Unable to connect to Database: %v", err)
	}

	DB = conn

	log.Println("Connected to Database successfully")
}