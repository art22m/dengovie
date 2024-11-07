package app

import (
	"context"
	"fmt"

	"gopkg.in/telebot.v4"
)

func (s *Service) Optimize(c telebot.Context) error {
	chat := c.Chat()
	if chat == nil || (chat.Type != telebot.ChatGroup && chat.Type != telebot.ChatSuperGroup) {
		s.Log.Printf("/optimize not in a group chat. Author: %d, ChatType: %s", c.Sender().ID, chat.Type)
		return nil
	}

	fromCount, toCount, err := s.Usecase.OptimizeDebts(context.TODO(), chat.ID)
	if err != nil {
		c.Send(fmt.Sprintf("Ошибка при оптимизации: %v", err))
		return nil
	}

	if fromCount == toCount {
		c.Send("Нечего оптимизировать:)")
	} else {
		c.Send(fmt.Sprintf("Уменьшено количество транзакций с %v до %v", fromCount, toCount))
	}
	return nil
}

func (s *Service) bindOptimizeDebts() {
	s.Bot.Handle("/optimize", s.Optimize)
}
