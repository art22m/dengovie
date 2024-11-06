package main

import (
	"context"
	"log"
	"time"

	"github.com/art22m/dengovie/internal/app"
	"github.com/art22m/dengovie/internal/config"
	"github.com/art22m/dengovie/internal/pkg/store"
	"github.com/art22m/dengovie/internal/pkg/store/repository/pg"
	"github.com/art22m/dengovie/internal/pkg/usecase"
)

func main() {
	ctx := context.Background()

	db, err := config.CreateDatabase(ctx)
	if err != nil {
		log.Fatal(err)
	}

	eventsRepo := pg.NewEvents(store.NewDatabase(db))
	//chatsRepo := pg.NewChats(store.NewDatabase(db))
	usersRepo := pg.NewUsers(store.NewDatabase(db))
	debtsRepo := pg.NewDebts(store.NewDatabase(db))

	useCase := usecase.NewUseCase(db, debtsRepo, eventsRepo, usersRepo)

	service := app.NewService(
		config.Telegram{
			PollTimeout: 10 * time.Second,
		},
		useCase,
	)

	service.Bot.Handle("/register", service.Register)
	service.Bot.Handle("/split", service.Split)

	service.Bot.Start()
}
