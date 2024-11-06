package app

import (
	"context"
	"errors"
	"strconv"

	"gopkg.in/telebot.v4"

	"github.com/art22m/dengovie/internal/pkg/usecase"
)

var (
	registerMenu       = &telebot.ReplyMarkup{}
	shareContactButton = registerMenu.Contact("Поделиться телефоном телефоном")
)

func (s *Service) Register(c telebot.Context) error {
	chat := c.Chat()
	if chat == nil || chat.Type != telebot.ChatPrivate {
		s.Log.Printf("/register not in a private chat '%s'. Author: %d", chat.Title, c.Sender().ID)
		return nil
	}

	s.Log.Printf("/register: id: %d, %s", chat.ID, chat.Username)

	c.Send("Добро пожаловать! Чтобы завершить регистрацию пришлите свой контакт.", registerMenu)
	return nil
}

func (s *Service) ShareContact(c telebot.Context) error {
	contact := c.Message().Contact
	chat := c.Chat()

	if chat == nil || chat.Type != telebot.ChatPrivate {
		s.Log.Printf("ShareContact button:  not in a private chat '%s'. Author: %d", chat.Title, c.Sender().ID)
		return nil
	}

	if contact.UserID != c.Chat().ID {
		c.Send("Это не ваш контакт. Мне он неинтересен")
		return nil
	}

	req := usecase.RegisterUserRequest{
		TelegramUserID: strconv.FormatInt(contact.UserID, 10),
		PhoneNumber:    contact.PhoneNumber,
		TelegramAlias:  &c.Sender().Username,
	}

	if err := s.Usecase.RegisterUser(context.TODO(), req); err != nil {
		if errors.Is(err, usecase.ErrUserAlreadyExists) {
			c.Send("Вы уже зарегистрированы.")
		}
		return err
	}

	c.Send("Вы успешно зарегестрированны!")

	return nil
}

func (s *Service) bindRegisterHandlers() {

	registerMenu.Reply(
		registerMenu.Row(shareContactButton),
	)
	s.Bot.Handle("/start", s.Register)
	s.Bot.Handle("/register", s.Register)
	s.Bot.Handle(telebot.OnContact, s.ShareContact)
}
