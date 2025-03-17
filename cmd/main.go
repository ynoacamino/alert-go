package main

import (
	"log"

	"github.com/ynoacamino/alert-go/config"
	"github.com/ynoacamino/alert-go/internal/app"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	app.Run(cfg)
}
