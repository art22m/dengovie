package app

import (
	"context"
	"errors"
	"strconv"

	"github.com/art22m/dengovie/internal/pkg/usecase"
	"gopkg.in/telebot.v4"
)

var (
	registerMenu = &telebot.ReplyMarkup{}
	shareContactButton = registerMenu.Contact("Поделитесь телефоном")
)

func (s *Service) Register(c telebot.Context) error {
	chat := c.Chat()
	if chat.Type != telebot.ChatPrivate {
		s.Log.Printf("/register not in a private chat '%s'. Author: %d", chat.Title, c.Sender().ID)
		return nil
	}

	s.Log.Printf("/register: id: %d, %s", chat.ID, chat.Username)

	registerMenu.Reply(
		registerMenu.Row(shareContactButton),
	)

	c.Send("Добро пожаловать!", registerMenu)
	return nil
}

func (s *Service) ShareContact(c telebot.Context) error {
	contact := c.Message().Contact
	if contact == nil {
		return errors.New("Empty contact")
	}

	req := usecase.RegisterUserRequest{
		TelegramUserID: strconv.FormatInt(contact.UserID, 10),
		PhoneNumber: contact.PhoneNumber,
		TelegramAlias: &c.Sender().Username,
	}

	if err := s.Usecase.Register(context.TODO(), req); err != nil {
		return err
	}

	c.Send("Вы успешно зарегестрированны!")
	
	return nil
}

func (s *Service) bindRegisterHandlers() {
	s.Bot.Handle("/register", s.Register)
	s.Bot.Handle(shareContactButton, s.ShareContact)
}
