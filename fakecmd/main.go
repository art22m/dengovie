package main

import (
	"context"
	"fmt"
	"log"

	"github.com/samber/lo"

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

	useCase := usecase.NewUseCase(db.GetPool(), debtsRepo, eventsRepo, usersRepo)

	//err = useCase.RegisterUser(ctx, usecase.RegisterUserRequest{
	//	UserID:      5050,
	//	PhoneNumber: "+123123123",
	//	Alias:       lo.ToPtr("edik"),
	//})
	//if err != nil {
	//	log.Fatal(err)
	//}

	//err = useCase.AddDebt(ctx, usecase.AddDebtRequest{
	//	CollectorID: 42,
	//	DebtorID:    1488,
	//	ChatID:      5,
	//	Amount:      30,
	//	Description: "test50",
	//})
	//if err != nil {
	//	log.Fatal(err)
	//}

	//err = useCase.ReturnDebt(ctx, usecase.ReturnDebtRequest{
	//	CollectorID: 42,
	//	DebtorID:    1488,
	//	ChatID:      4,
	//	Amount:      29,
	//	Description: "beer",
	//})
	//if err != nil {
	//	log.Fatal(err)
	//}

	//err = useCase.SplitDebt(ctx, usecase.SplitDebtRequest{
	//	CollectorID: 42,
	//	DebtorIDs:   []int64{1488, 5050},
	//	ChatID:      4,
	//	TotalAmount: 8,
	//	Description: "beer",
	//})
	//if err != nil {
	//	log.Fatal(err)
	//}

	resp, err := useCase.ListDebts(ctx, usecase.ListDebtsRequest{
		UserID: 42,
		ChatID: 4,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("!!!", lo.FromPtr(resp.CollectorTelegramAlias), resp.CollectorTelegramID)
	for _, info := range resp.DebtsInfo {
		fmt.Println(lo.FromPtr(info.DebtorTelegramAlias), info.DebtorTelegramID, info.Amount, info.LastModified)
	}

	//allUsers, err := useCase.ListAllUsers(ctx)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println("!!!", allUsers)
}
