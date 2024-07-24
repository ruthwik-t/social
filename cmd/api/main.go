package main

import (
	"log"

	"github.com/ruthwik-t/social/internal/env"
	"github.com/ruthwik-t/social/internal/store"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
	}

	storage := store.NewStorage(nil)

	app := &application{
		config:  cfg,
		storage: storage,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
