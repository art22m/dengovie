package app

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/art22m/dengovie/internal/pkg/usecase"
	"github.com/pkg/errors"
	"github.com/samber/lo"
	"gopkg.in/telebot.v4"
)

func (s *Service) ListDebts(c telebot.Context) error {
	chat := c.Chat()
	if chat == nil || (chat.Type != telebot.ChatGroup && chat.Type != telebot.ChatSuperGroup) {
		s.Log.Printf("/list not in a group chat. Author: %d, ChatType: %s", c.Sender().ID, chat.Type)
		return nil
	}

	resp, err := s.Usecase.ListDebts(context.Background(),
		usecase.ListDebtsRequest{
			TelegramUserID: strconv.FormatInt(c.Sender().ID, 10),
			TelegramChatID: strconv.FormatInt(chat.ID, 10),
		},
	)

	switch {
	case errors.Is(err, usecase.ErrUserNotFound):
		c.Send("Вы не зарегистрированы!")
		return err
	case errors.Is(err, usecase.ErrChatNotFound):
		c.Send("Неизвестный чат!")
		return err
	case err != nil:
		c.Send("Что-то пошло не так...")
		return err
	}

	if resp == nil {
		c.Send("Что-то пошло не так...")
		return nil
	}

	c.Send(s.buildResponse(*resp))
	return nil
}

func (s *Service) buildResponse(resp usecase.ListDebtsResponse) string {
	builder := strings.Builder{}
	collectorMention := fmt.Sprintf("[коллектора](tg://user?id=%s)\n\n", resp.CollectorTelegramID)
	if resp.CollectorTelegramAlias != nil {
		collectorMention = fmt.Sprintf(
			"@%s",
			lo.FromPtr(resp.CollectorTelegramAlias),
		)
	}

	builder.Write([]byte(fmt.Sprintf("Список долгов для %s\n\n", collectorMention)))

	for i, debt := range resp.DebtsInfo {
		debtorMention := fmt.Sprintf("[unknown](tg://user?id=%s)", debt.DebtorTelegramID)
		if debt.DebtorTelegramAlias != nil {
			debtorMention = fmt.Sprintf(
				"@%s",
				lo.FromPtr(debt.DebtorTelegramAlias),
			)
		}

		var debtText string
		switch {
		case debt.Amount > 0:
			debtText = fmt.Sprintf(
				"%d) %s должен %s %d,%d\n",
				i+1,
				debtorMention,
				collectorMention,
				debt.Amount/100,
				debt.Amount%100,
			)
		case debt.Amount < 0:
			amount := -debt.Amount
			debtText = fmt.Sprintf(
				"%d) %s должен %s %d,%d\n",
				i+1,
				collectorMention,
				debtorMention,
				amount/100,
				amount%100,
			)
		default:
			continue
		}

		builder.Write([]byte(debtText))
	}

	return builder.String()
}

func (s *Service) bindListDebts() {
	s.Bot.Handle("/list", s.ListDebts)
}
