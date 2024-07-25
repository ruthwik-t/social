package main

import (
	"log"

	"github.com/ruthwik-t/social/internal/config"
	"github.com/ruthwik-t/social/internal/db"
	"github.com/ruthwik-t/social/internal/store"
)

const version = "0.0.1"

func main() {
	cfg := config.LoadConfig()

	database, err := db.NewDB(
		cfg.DB.Address,
		cfg.DB.MaxOpenConns,
		cfg.DB.MaxIdleConns,
		cfg.DB.MaxIdleTime,
	)
	if err != nil {
		log.Panic(err)
	}

	storage := store.NewStorage(database)

	app := &application{
		config:  cfg,
		storage: storage,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
