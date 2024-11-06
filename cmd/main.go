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

	debtsRepo := pg.NewDebts(store.NewDatabase(db))
	eventsRepo := pg.NewEvents(store.NewDatabase(db))
	usersRepo := pg.NewUsers(store.NewDatabase(db))
	chatsRepo := pg.NewChats(store.NewDatabase(db))

	useCase := usecase.NewUseCase(db, debtsRepo, eventsRepo, usersRepo, chatsRepo)

	service := app.NewService(
		config.Telegram{
			PollTimeout: 10 * time.Second,
		},
		useCase,
	)
	service.BindBotHandlers()
	service.Bot.Start()
}
