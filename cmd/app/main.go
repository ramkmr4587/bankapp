package main

import (
	"bankapp/database"
	"bankapp/internal/server"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// Build full path to your .env in project root
	envPath := filepath.Join(dir, ".env")

	// Load it explicitly
	if err := godotenv.Load(envPath); err != nil {
		log.Printf("Warning: could not load .env from %s: %v", envPath, err)
	} else {
		log.Printf(".env loaded from %s", envPath)
	}

	secret := os.Getenv("JWT_SECRET")
	log.Printf(" JWT_SECRET: %s", secret)
	database.Init()
	server.Start()
}
