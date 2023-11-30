package main

import (
	"info-bot/config"
	"info-bot/internal"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	app := internal.New(cfg)
	if err = app.Run(); err != nil {
		panic(err)
	}
}
