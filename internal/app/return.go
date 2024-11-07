package app

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"gopkg.in/telebot.v4"

	"github.com/art22m/dengovie/internal/pkg/usecase"
)

func (s *Service) Return(c telebot.Context) error {
	chat := c.Chat()
	switch chat.Type {
	case telebot.ChatGroup:
		return s.returnInGroup(c)
	default:
		return errors.New("Unsupported chat type for return")
	}
}

func (s *Service) ReturnSelect(c telebot.Context) error {
	cb := c.Callback()

	data := returnBtnData{}
	if err := json.Unmarshal([]byte(cb.Data), &data); err != nil {
		return fmt.Errorf("Can't unmarshal return data button: %w", err)
	}

	req := usecase.ReturnDebtRequest{
		CollectorID: data.TelegramID,
		DebtorID:    c.Sender().ID,
		ChatID:      c.Chat().ID,
		Amount:      data.Amount * 100,
	}

	if err := s.Usecase.ReturnDebt(context.TODO(), req); err != nil {
		c.Send("Ошибочка вышла сорян")
		return fmt.Errorf("Errors during usecase: %w", err)
	}

	c.Send(fmt.Sprintf("@%s вернул(-a) %d рублей @%s", c.Sender().Username, data.Amount, data.UserScreenName))
	c.Bot().Delete(cb.Message)
	return nil
}

func (s *Service) returnInGroup(c telebot.Context) error {
	req := usecase.ListDebtsRequest{
		UserID: c.Sender().ID,
		ChatID: c.Chat().ID,
	}

	debtResp, err := s.Usecase.ListDebts(context.TODO(), req)
	if err != nil {
		c.Send("Sorry, can't return right now")
	}

	markup, err := makeKeyboardForReturn(debtResp.DebtsInfo)
	if err != nil {
		c.Send("Ошибка, попробуйте позже :(")
		return fmt.Errorf("Can't create keyboard for return: ", err)
	}

	c.Send("Выберите кому вернуть деньги", markup)
	return nil
}

func makeKeyboardForReturn(infos []usecase.DebtInfo) (*telebot.ReplyMarkup, error) {
	selector := telebot.ReplyMarkup{}

	rows := make([]telebot.Row, MaxUsersInColumn)
	for i, info := range infos {
		row_i := i % MaxUsersInColumn

		if info.Amount >= 0 {
			continue
		}

		btn, err := makeUserButtonReturn(info)

		if err != nil {
			return nil, fmt.Errorf("Can't create selector button for return: %w", err)
		}
		rows[row_i] = append(rows[row_i], btn)

		i++
	}
	selector.Inline(rows...)
	return &selector, nil
}

type returnBtnData struct {
	CallbackData
	UserScreenName string `json:"sn"`
	TelegramID     int64  `json:"id"`
	Amount         int64  `json:"amount"`
}

func (data returnBtnData) Text() string {
	return fmt.Sprintf("%s (%d рублей)", data.UserScreenName, data.Amount)
}

func makeUserButtonReturn(info usecase.DebtInfo) (telebot.Btn, error) {
	data := returnBtnData{
		CallbackData:   CallbackData{Type: ButtonDataReturn},
		UserScreenName: *info.DebtorTelegramAlias,
		TelegramID:     info.DebtorTelegramID,
		Amount:         -info.Amount / 100,
	}

	dataRaw, err := json.Marshal(data)
	if err != nil {
		return telebot.Btn{}, err
	}

	return telebot.Btn{
		Text: data.Text(),
		Data: string(dataRaw),
	}, nil
}

func (s *Service) bindReturnHandlers() {
	s.Bot.Handle("/return", s.Return)
}
