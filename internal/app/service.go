package app

import (
	"log"
	"os"

	"github.com/art22m/dengovie/internal/pkg/usecase"

	"github.com/art22m/dengovie/internal/config"
	tele "gopkg.in/telebot.v4"
)

type Service struct {
	Bot *tele.Bot
	Log *log.Logger
	Usecase *usecase.UseCase
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
		Log: log.Default(),
		Usecase: nil, // TODO(artemreyt): занести provideUseCase
	}
}

func provideBotApi(conf config.Telegram) (*tele.Bot, error) {
	pref := tele.Settings{
		Token:  os.Getenv("TOKEN"),
		Poller: &tele.LongPoller{
			Timeout: conf.PollTimeout,
		},
		Verbose: true,
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (s *Service) BindBotHandlers() {
	s.bindRegisterHandlers()
}
