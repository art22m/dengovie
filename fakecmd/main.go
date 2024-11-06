package main

import (
	"context"
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

	//err = useCase.Register(ctx, usecase.RegisterUserRequest{
	//	TelegramUserID: "123",
	//	PhoneNumber:    "+123123123",
	//	TelegramAlias:  "art22m",
	//})
	//if err != nil {
	//	log.Fatal(err)
	//}

	//err = useCase.AddDebt(ctx, usecase.AddDebtRequest{
	//	CollectorID: 2,
	//	DebtorID:    1,
	//	ChatID:      3,
	//	Amount:      30,
	//	Description: "test50",
	//})
	//if err != nil {
	//	log.Fatal(err)
	//}

	//err = useCase.ReturnDebt(ctx, usecase.ReturnDebtRequest{
	//	CollectorID: 1,
	//	DebtorID:    2,
	//	ChatID:      3,
	//	Amount:      5,
	//	Description: "beer",
	//})
	//if err != nil {
	//	log.Fatal(err)
	//}

	err = useCase.SplitDebt(ctx, usecase.SplitDebtRequest{
		CollectorID: 1,
		DebtorIDs:   []int64{2, 3},
		ChatID:      3,
		TotalAmount: 10,
		Description: "beer",
	})
	if err != nil {
		log.Fatal(err)
	}
}
