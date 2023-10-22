package db

import (
	"log"
	"os"

	"github.com/go-pg/pg/v9"
)

func Connect() *pg.DB {
	db := connectToDB()
	createTable(db)
	return db
}

func connectToDB() *pg.DB {
	options := &pg.Options{
		Addr:     getEnv("DB_HOST", "postgres") + ":" + getEnv("DB_PORT", "5432"),
		User:     getEnv("DB_USERNAME", "postgres"),
		Password: getEnv("DB_PASSWORD", "postgres"),
		Database: getEnv("DB_DATABASE", "db"),
	}

	db := pg.Connect(options)
	if db == nil {
		log.Fatal("Failed to connect to the database")
		os.Exit(1)
	}
	log.Printf("Connected to the database")
	return db
}

func createTable(db *pg.DB) {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS logs (
		log_id SERIAL PRIMARY KEY,
    timestamp TIMESTAMPTZ DEFAULT current_timestamp,
    method VARCHAR(10),
    url TEXT,
    request_body TEXT,
    response_status INT,
		response_time INT
	);`)
	if err != nil {
		log.Printf("Error while creating table, reason: %v\n", err)
	} else {
		log.Printf("Table created successfully")
	}

}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
