package main

import (
	"log"

	"github.com/Unchana19/go-learn/internal/db"
	"github.com/Unchana19/go-learn/internal/env"
	"github.com/Unchana19/go-learn/internal/store"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found")
	}

	addr := env.GetString("DB_ADDR", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")

	conn, err := db.New(addr, 3, 3, "15m")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	store := store.NewStorage(conn)

	db.Seed(store, conn)
}
