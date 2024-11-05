package app

import (
	"log"
	"os"

	"github.com/art22m/dengovie/internal/config"
	tele "gopkg.in/telebot.v4"
)

type Service struct {
	Bot *tele.Bot
}

func NewService(
	tg config.Telegram,
) (*Service) {
	bot, err := provideBotApi(tg)
	if err != nil {
		log.Fatal(err)
	}

	return &Service{
		Bot: bot,
	}
}

func provideBotApi(conf config.Telegram) (*tele.Bot, error) {
	pref := tele.Settings{
		Token:  os.Getenv("TOKEN"),
		Poller: &tele.LongPoller{
			Timeout: conf.PollTimeout,
		},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		return nil, err
	}

	return b, nil
}
