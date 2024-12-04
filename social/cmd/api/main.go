package main

import (
	"log"

	"github.com/Unchana19/social/internal/env"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":9000"),
	}

	app := &application{
		config: cfg,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
