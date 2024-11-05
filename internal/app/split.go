package app

import (
	"fmt"

	"gopkg.in/telebot.v4"
)

func (s *Service) Split(c telebot.Context) error {
	c.Send(fmt.Sprintf("Hello! Split detected. Text: %s", c.Text()))
	return nil
}
