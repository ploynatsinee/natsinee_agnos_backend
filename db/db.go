package db

import (
	"fmt"
	"log"
	"os"

	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
)

type Log struct {
	ID           int
	Method       string
	URL          string
	Status       int
	ResponseTime int
}

func Connect() *pg.DB {
	db := connectToDB()
	createTable(db)
	return db
}

func connectToDB() *pg.DB {
	options := &pg.Options{
		Addr:     getEnv("DB_HOST", "127.0.0.1") + ":" + getEnv("DB_PORT", "5432"),
		User:     getEnv("DB_USERNAME", "postgres"),
		Password: getEnv("DB_PASSWORD", "postgres"),
		Database: getEnv("DB_DATABASE", "db"),
	}

	fmt.Printf("DB User: %s\n", options.User)
	fmt.Printf("DB Password: %s\n", options.Password)

	db := pg.Connect(options)
	if db == nil {
		log.Fatal("Failed to connect to the database")
		os.Exit(1)
	}
	log.Printf("Connected to the database")
	return db
}

func createTable(db *pg.DB) {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}

	createError := db.CreateTable(&Log{}, opts)
	if createError != nil {
		log.Printf("Error while creating table, reason: %v\n", createError)
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
