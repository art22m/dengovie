package main

import (
	"time"

	"github.com/art22m/dengovie/internal/app"
	"github.com/art22m/dengovie/internal/config"
)

func main() {
	service := app.NewService(config.Telegram{
		PollTimeout: 10 * time.Second,
	})

	service.Bot.Handle("/register", service.Register)
	service.Bot.Handle("/split", service.Split)

	service.Bot.Start()
}
