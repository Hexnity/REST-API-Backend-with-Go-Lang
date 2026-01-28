package main

import (
	"database/sql"
	"fmt" // Added for Sprintf
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {

	// Load environment variables from .env file
	err := godotenv.Load("internal/env/.env")
	if err != nil {
		log.Println("No .env file found in env/ folder, proceeding with system envs")
	}

	if len(os.Args) < 2 {
		log.Fatalln("No command provided. Use 'up' or 'down'.")
	}

	direction := os.Args[1]

	// 1. Fetch values from environment
	dbUser := os.Getenv("DB_USERNAME")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_DATABASE")
	sslMode := os.Getenv("DB_SSLMODE")

	// Set defaults if empty
	if dbPort == "" {
		dbPort = "5432"
	}
	if sslMode == "" {
		sslMode = "disable"
	}

	// 2. Format the DSN (Data Source Name)
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		dbUser, dbPass, dbHost, dbPort, dbName, sslMode)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v\n", err)
	}
	defer db.Close()

	// ... (rest of your migration logic remains the same)
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("Failed to create migration instance: %v\n", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://cmd/migrate/migrations",
		"postgres",
		driver,
	)

	m.Force(2)

	if err != nil {
		log.Fatalf("Failed to create migrate instance: %v\n", err)
	}

	if direction == "up" {
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("Migration failed: %v", err)
		}
	} else {
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("Migration failed: %v", err)
		}
	}
}
