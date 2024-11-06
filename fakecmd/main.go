package main

import (
	"context"
	"fmt"
	"log"

	"github.com/art22m/dengovie/internal/pkg/store"
	"github.com/art22m/dengovie/internal/pkg/store/repository/pg"
	"github.com/art22m/dengovie/internal/pkg/usecase"
)

func main() {
	ctx := context.Background()

	//db, err := config.CreateDatabase(ctx)
	db, err := store.CreateDatabase(ctx)
	if err != nil {
		log.Fatal(err)
	}

	debtsRepo := pg.NewDebts(store.NewDatabase(db.GetPool()))
	eventsRepo := pg.NewEvents(store.NewDatabase(db.GetPool()))
	usersRepo := pg.NewUsers(store.NewDatabase(db.GetPool()))
	chatsRepo := pg.NewChats(store.NewDatabase(db.GetPool()))

	useCase := usecase.NewUseCase(db.GetPool(), debtsRepo, eventsRepo, usersRepo, chatsRepo)

	al := "art22m"
	err = useCase.Register(ctx, usecase.RegisterUserRequest{
		TelegramUserID: "123",
		PhoneNumber:    "+123123123",
		TelegramAlias:  &al,
	})
	fmt.Println("err=", err)
}
