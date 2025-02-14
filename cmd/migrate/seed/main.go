package main

import (
	"github.com/tabrizgulmammadov/go-articles-platform/internal/db"
	"github.com/tabrizgulmammadov/go-articles-platform/internal/env"
	"github.com/tabrizgulmammadov/go-articles-platform/internal/store"
	"log"
)

func main() {
	addr := env.GetString("DB_ADDR", "postgres://postgres:admin@localhost:5432/social?sslmode=disabled")
	conn, err := db.New(addr, 3, 3, "15m")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	store := store.NewStorage(conn)

	db.Seed(store)
}
