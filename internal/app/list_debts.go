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
	if chat == nil || chat.Type != telebot.ChatGroup {
		s.Log.Printf("/list not in a group chat. Author: %d", c.Sender().ID)
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
	case errors.Is(err, usecase.ErrChatNotFound):
		c.Send("Неизвестный чат!")
	case err != nil:
		c.Send("Что-то пошло не так...")
		return err
	}

	if resp == nil {
		c.Send("Что-то пошло не так...")
		return nil
	}

	c.Send(buildResponse(*resp))
	return nil
}

func buildResponse(resp usecase.ListDebtsResponse) string {
	builder := strings.Builder{}
	collectorMention := fmt.Sprintf("[коллектора](tg://user?id=%s)\n\n", resp.CollectorTelegramID)
	if resp.CollectorTelegramAlias != nil {
		collectorMention = fmt.Sprintf(
			"[%s](tg://user?id=%s)\n\n",
			lo.FromPtr(resp.CollectorTelegramAlias),
			resp.CollectorTelegramID,
		)
	}

	builder.Write([]byte(fmt.Sprintf("Список долгов для %s", collectorMention)))
	for i, debt := range resp.DebtsInfo {
		debtorMention := fmt.Sprintf("[unknown](tg://user?id=%s)", debt.DebtorTelegramID)
		if debt.DebtorTelegramAlias != nil {
			debtorMention = fmt.Sprintf(
				"[%s](tg://user?id=%s)",
				lo.FromPtr(debt.DebtorTelegramAlias),
				debt.DebtorTelegramID,
			)
		}
		var debtText string
		switch {
		case debt.Amount > 0:
			debtText = fmt.Sprintf(
				"%d) %s должен %s %d,%d\n",
				i,
				debtorMention,
				collectorMention,
				debt.Amount/100,
				debt.Amount%100,
			)
		case debt.Amount < 0:
			debtText = fmt.Sprintf(
				"%d) %s должен %s %d,%d\n",
				i,
				collectorMention,
				debtorMention,
				debt.Amount/100,
				debt.Amount%100,
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
