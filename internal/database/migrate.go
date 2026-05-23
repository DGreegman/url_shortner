package database

import (
	"context"
	"log"
)


func Migrate() {
	query := `
		CREATE TABLE IF NOT EXISTS urls(
			id SERIAL PRIMARY KEY,
			code TEXT NOT NULL UNIQUE,
			target_url TEXT NOT NULL,
			clicks INT DEFAULT 0,
			created_at TIMESTAMP DEFAULT NOW(),
			expire_at TIMESTAMP NOT NULL
		);
	
	`

	_, err := DB.Exec(context.Background(), query)
	if err != nil {
		log.Fatalf("Migration Failed: %v", err)
	}

	log.Println("Database migration Completed")
}