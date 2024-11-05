package app

import (
	"gopkg.in/telebot.v4"
)

func (s *Service) Register(c telebot.Context) error {
	c.Send("Добро пожаловать! Теперь вы зарегестированы в боте")
	return nil
}
