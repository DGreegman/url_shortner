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
			redirect_type TEXT NOT NULL DEFAULT '302',
			clicks INT DEFAULT 0,
			created_at TIMESTAMP DEFAULT NOW(),
			expire_at TIMESTAMP NOT NULL
		);

		ALTER TABLE urls 
		ADD COLUMN IF NOT EXISTS redirect_type TEXT NOT NULL DEFAULT '302';

		DO $$
		BEGIN
			IF NOT EXISTS (
				SELECT 1
				FROM pg_constraint
				WHERE conname = 'urls_redirect_type_check'
			) THEN
				ALTER TABLE urls
				ADD CONSTRAINT urls_redirect_type_check 
				CHECK (redirect_type IN ('301', '302', '307'));
			END IF;
		END
		$$;
			
	
	`

	_, err := DB.Exec(context.Background(), query)
	if err != nil {
		log.Fatalf("Migration Failed: %v", err)
	}

	log.Println("Database migration Completed")
}