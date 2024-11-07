package app

import (
	"context"
	"strconv"

	"github.com/art22m/dengovie/internal/pkg/usecase"
	"github.com/pkg/errors"
	"gopkg.in/telebot.v4"
)

func (s *Service) RegisterChat(c telebot.Context) error {
	chat := c.Chat()
	if chat == nil || (chat.Type != telebot.ChatGroup && chat.Type != telebot.ChatSuperGroup) {
		s.Log.Printf("adding bot not in a group chat. Author: %d, ChatType: %s", c.Sender().ID, chat.Type)
		return nil
	}

	if err := s.Usecase.RegisterChat(context.Background(), usecase.RegisterChatRequest{
		TelegramChatID: strconv.FormatInt(chat.ID, 10),
		Info:           chat.Description,
	}); err != nil {
		switch {
		case errors.Is(err, usecase.ErrChatAlreadyExists):
			c.Send("Бот уже добавлен в чат")
			return nil
		default:
			c.Send("Не удалось добавить бота в чат")
			return err
		}

	}

	c.Send("Бот успешно добавлен в чат")
	return nil
}

func (s *Service) bindRegisterChat() {
	s.Bot.Handle(telebot.OnAddedToGroup, s.RegisterChat)
}
