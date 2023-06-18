package main

import (
	"flag"

	"management.api/internal/config"
)

func main() {
	c := config.New()

	flag.StringVar(&c.Port, "port", c.Port, "override port")
	flag.Parse()

	if err := run(c); err != nil {
	}
}

func run(c *config.Config) error {
	return nil
}
