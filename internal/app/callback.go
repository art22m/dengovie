package app

import (
	"encoding/json"
	"fmt"

	"gopkg.in/telebot.v4"
)


type ButtonDataType = int
const (
	ButtonDataSplitSelect ButtonDataType = iota + 1
	ButtonDataSplitOk
	ButtonDataReturn
)

type CallbackData struct {
	Type ButtonDataType `json:"t"`
}

func (s *Service) onCallback(c telebot.Context) error {
	cb := c.Callback()

	cbData := CallbackData{}
	if err := json.Unmarshal([]byte(cb.Data), &cbData); err != nil {
		return fmt.Errorf("Can't unmarshall callback data type: %w", err)
	}

	s.Log.Printf("Callback data: %d", cbData.Type)

	switch (cbData.Type) {
	case ButtonDataSplitSelect:
		return s.SplitSelect(c)
	case ButtonDataSplitOk:
		return s.SplitOk(c)
	case ButtonDataReturn:
		return s.ReturnSelect(c)
	default:
		return fmt.Errorf("Unknown callback data button: %d", cbData.Type)
	}
}

func (s *Service) bindCallbackHandlers() {
	s.Bot.Handle(telebot.OnCallback, s.onCallback)
}