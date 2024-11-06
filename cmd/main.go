package main

import (
	"context"
	"log"
	"time"

	"github.com/art22m/dengovie/internal/app"
	"github.com/art22m/dengovie/internal/config"
	"github.com/art22m/dengovie/internal/generated/dengovie/dengovie/public/model"
	"github.com/art22m/dengovie/internal/generated/dengovie/dengovie/public/table"
)

func main() {
	ctx := context.Background()

	// КОД ДЛЯ ТЕСТОВ КОННЕКТА К БД ДЛЯ ЭДУАРДА
	db, err := config.CreateDatabase(ctx)
	if err != nil {
		log.Fatal(err)
	}

	stmt, args := table.Chats.SELECT(table.Chats.AllColumns).LIMIT(1).Sql()
	var chat model.Chats
	if err = db.Get(ctx, &chat, stmt, args...); err != nil {
		log.Fatal(err)
	}

	service := app.NewService(config.Telegram{
		PollTimeout: 10 * time.Second,
	})

	service.Bot.Handle("/register", service.Register)
	service.Bot.Handle("/split", service.Split)

	service.Bot.Start()
}
