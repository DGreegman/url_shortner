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
			
	
		CREATE TABLE IF NOT EXISTS click_events (
		event_id SERIAL PRIMARY KEY,
		link_id INT NOT NULL REFERENCES urls(id) ON DELETE CASCADE,
		ip VARCHAR(45),
		user_agent TEXT,
		referrer TEXT,
		device_type VARCHAR(20),
		country VARCHAR(50),
		ts TIMESTAMP DEFAULT NOW()
	);

	CREATE INDEX IF NOT EXISTS idx_click_events_link_id ON click_events(link_id);
	CREATE INDEX IF NOT EXISTS idx_click_events_timestamp ON click_events(ts);
	`

	_, err := DB.Exec(context.Background(), query)
	if err != nil {
		log.Fatalf("Migration Failed: %v", err)
	}

	log.Println("Database migration Completed")
}