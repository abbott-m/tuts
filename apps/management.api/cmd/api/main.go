package main

import (
	"flag"
	"log"

	"github.com/go-chi/chi/v5"

	"management.api/internal/config"
	"management.api/internal/handlers"
)

func main() {
	c := config.New()
	flag.StringVar(&c.Port, "port", c.Port, "override port")
	flag.Parse()

	if err := run(c); err != nil {
		log.Fatal(err)
	}
}

func run(c *config.Config) error {
	mux := chi.NewMux()
	opts := handlers.APIOptions{Port: c.Port}
	api := handlers.NewAPI(mux, opts)

	return api.Start()
}
