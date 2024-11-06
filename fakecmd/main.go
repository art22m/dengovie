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

	//err = useCase.RegisterUser(ctx, usecase.RegisterUserRequest{
	//	TelegramUserID: "triplo",
	//	PhoneNumber:    "+123123123",
	//	TelegramAlias:  lo.ToPtr("triplo"),
	//})
	//if err != nil {
	//	log.Fatal(err)
	//}

	//err = useCase.RegisterChat(ctx, usecase.RegisterChatRequest{
	//	TelegramChatID: "3",
	//	Info:           "test info",
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

	//err = useCase.SplitDebt(ctx, usecase.SplitDebtRequest{
	//	CollectorID: 1,
	//	DebtorIDs:   []int64{2, 3},
	//	ChatID:      3,
	//	TotalAmount: 10,
	//	Description: "beer",
	//})
	//if err != nil {
	//	log.Fatal(err)
	//}

	//resp, err := useCase.ListDebts(ctx, usecase.ListDebtsRequest{
	//	TelegramUserID: "123",
	//	TelegramChatID: "3",
	//})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println("!!!", lo.FromPtr(resp.CollectorTelegramAlias), resp.CollectorTelegramID)
	//for _, info := range resp.DebtsInfo {
	//	fmt.Println(lo.FromPtr(info.DebtorTelegramAlias), info.DebtorTelegramID, info.Amount, info.LastModified)
	//}

	allUsers, err := useCase.ListAllUsers(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("!!!", allUsers)
}
