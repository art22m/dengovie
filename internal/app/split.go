package app

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/art22m/dengovie/internal/pkg/usecase"
	"gopkg.in/telebot.v4"
)

const (
	MaxUsersInColumn  = 5
	MaxDescriptionLen = 32

	SplitButtonText    = "OK"
	SelectUserSymbol   = "✅"
	DeselectUserSymbol = "❌"
)

func (s *Service) Split(c telebot.Context) error {
	chat := c.Chat()
	switch chat.Type {
	case telebot.ChatGroup, telebot.ChatSuperGroup:
		return s.splitInGroup(c)
	default:
		return errors.New("Unsupported chat type for split")
	}
}

func (s *Service) splitInGroup(c telebot.Context) error {
	chat := c.Chat()

	amount, description, err := validateSplitArgs(c)
	if err != nil {
		return fmt.Errorf("Args are not valid: %w", err)
	}

	members, err := s.getAllRegisteredMembers(chat)
	if err != nil {
		return fmt.Errorf("Can't get members of chat: %w", err)
	}

	selector, err := makeKeyboardForSplit(c.Sender().ID, members, amount, description)
	if err != nil {
		c.Send("Прозошла ошибка, попробуйте еще раз! :(")
		return fmt.Errorf("Can't create selector keyboard for split: %w", err)
	}

	c.Bot().Reply(c.Message(), "Выберите на кого разбить долг:", selector)
	return nil
}

func validateSplitArgs(c telebot.Context) (int64, string, error) {
	args := c.Args()

	if len(args) != 2 {
		c.Bot().Reply(c.Message(), "Expected 2 args for split")
		return 0, "", fmt.Errorf("Expected 2 args for split, got %d", len(args))
	}

	amount, err := strconv.ParseInt(args[0], 10, 64)
	if err != nil {
		c.Bot().Reply(c.Message(), "Bad amount")
		return 0, "", fmt.Errorf("Bad amount: %w", err)
	}

	desciption := args[1]
	if len(desciption) > MaxDescriptionLen {
		c.Bot().Reply(c.Message(), "Too big description")
		return 0, "", fmt.Errorf("Too big description, max allowed: %d, got %d", MaxDescriptionLen, len(desciption))
	}

	return amount, desciption, nil
}

func makeKeyboardForSplit(authorID int64, users []usecase.UserInfo, amount int64, description string) (*telebot.ReplyMarkup, error) {
	selector := telebot.ReplyMarkup{}

	rows := make([]telebot.Row, MaxUsersInColumn)
	for i, user := range users {
		row_i := i % MaxUsersInColumn
		col_i := len(rows[row_i])

		btn, err := makeUserButton(&user, row_i, col_i)
		if err != nil {
			return nil, fmt.Errorf("Can't create selector button: %w", err)
		}
		rows[row_i] = append(rows[row_i], btn)

		i++
	}

	okBtn, err := createOkButton(authorID, amount, description)
	if err != nil {
		return nil, fmt.Errorf("Can't create OK button for split")
	}
	rows = append(rows, telebot.Row{okBtn})

	selector.Inline(rows...)
	return &selector, nil
}

func userInfoScreenName(ui *usecase.UserInfo) string {
	if ui.TelegramAlias != nil {
		return *ui.TelegramAlias
	}

	return ui.TelegramUserID
}

type okBtnData struct {
	CollectorID int64  `json:"id"`
	Amount      int64  `json:"amount"`
	Description string `json:"desc"`
}

func createOkButton(collectorID int64, amount int64, description string) (telebot.Btn, error) {
	data := okBtnData{
		CollectorID: collectorID,
		Amount:      amount,
		Description: description,
	}

	dataRaw, err := json.Marshal(data)
	if err != nil {
		return telebot.Btn{}, err
	}

	return telebot.Btn{
		Text: SplitButtonText,
		Data: string(dataRaw),
	}, nil
}

type selectorBtnData struct {
	Chosen         bool   `json:"chosen"`
	UserScreenName string `json:"sn"`
	RowIndex       int    `json:"i"`
	ColumnIndex    int    `json:"j"`
	TelegramID     int64  `json:"id"`
}

func (data selectorBtnData) Text() string {
	text := data.UserScreenName
	if data.Chosen {
		text += DeselectUserSymbol
	} else {
		text += SelectUserSymbol
	}

	return text
}

func (data *selectorBtnData) Press() {
	data.Chosen = !data.Chosen
}

func makeUserButton(user *usecase.UserInfo, i int, j int) (telebot.Btn, error) {
	tgID, err := strconv.ParseInt(user.TelegramUserID, 10, 64)
	if err != nil {
		return telebot.Btn{}, err
	}

	data := selectorBtnData{
		Chosen:         false,
		UserScreenName: userInfoScreenName(user),
		RowIndex:       i,
		ColumnIndex:    j,
		TelegramID:     tgID,
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

type userRecipient struct {
	usecase.UserInfo
}

func (u userRecipient) Recipient() string {
	return u.TelegramUserID
}

func isMember(m *telebot.ChatMember) bool {
	return m.Role != telebot.Left && m.Role != telebot.Kicked
}

// TODO(artemreyt): redo - fullscan is very bad
func (s *Service) getAllRegisteredMembers(chat *telebot.Chat) ([]usecase.UserInfo, error) {
	users, err := s.Usecase.ListAllUsers(context.TODO())
	if err != nil {
		return nil, err
	}

	result := []usecase.UserInfo{}

	for _, user := range users {
		member, err := s.Bot.ChatMemberOf(chat, userRecipient{user})
		if err != nil {
			return nil, err
		}

		s.Log.Printf("Check user %s is chat member of %d - %t", *user.TelegramAlias, chat.ID, isMember(member))

		if isMember(member) {
			result = append(result, user)
		}
	}
	return result, nil
}

func (s *Service) onCallback(c telebot.Context) error {
	cb := c.Callback()

	okData, err := getOkData(c)
	if err != nil {
		return fmt.Errorf("error getting okData: %w", err)
	}

	if cb.Sender.ID != okData.CollectorID {
		c.RespondAlert("Это не ваш сплит")
		return nil
	}

	// TODO: kostyl
	if strings.Contains(c.Callback().Data, "\"amount\":") {
		err := s.completeSplit(c)
		if err != nil {
			c.Send("Sorry, can't split right now")
		}

		c.Bot().Delete(c.Callback().Message)
		return nil
	}

	var data selectorBtnData
	if err := json.Unmarshal([]byte(cb.Data), &data); err != nil {
		return fmt.Errorf("Unknown callback data: %w", err)
	}

	data.Press()
	newDataBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	markup := cb.Message.ReplyMarkup
	if markup == nil {
		return errors.New("Empty markup")
	}

	pressedBtn := markup.InlineKeyboard[data.RowIndex][data.ColumnIndex]
	pressedBtn.Data = string(newDataBytes)
	pressedBtn.Text = data.Text()

	markup.InlineKeyboard[data.RowIndex][data.ColumnIndex] = pressedBtn

	c.Bot().EditReplyMarkup(cb, markup)
	c.RespondText("Ok")
	return nil
}

func (s *Service) completeSplit(c telebot.Context) error {
	keyboard := c.Callback().Message.ReplyMarkup.InlineKeyboard
	if keyboard == nil {
		return fmt.Errorf("Keyboard is empty in completeSplit")
	}

	okDataStr := c.Callback().Data
	okData := okBtnData{}
	if err := json.Unmarshal([]byte(okDataStr), &okData); err != nil {
		return fmt.Errorf("Can't parse ok button data: %w", err)
	}

	debtorIDs := make([]int64, 0)
	debtorScreenames := make([]string, 0)
	for i, row := range keyboard {
		for _, btn := range row {

			if i == len(keyboard)-1 {
				// OK button
				continue
			}

			data := selectorBtnData{}
			if err := json.Unmarshal([]byte(btn.Data), &data); err != nil {
				return fmt.Errorf("Can't parse selector data")
			}

			if data.Chosen {
				debtorIDs = append(debtorIDs, data.TelegramID)
				debtorScreenames = append(debtorScreenames, data.UserScreenName)
			}
		}
	}

	if len(debtorIDs) == 0 {
		c.RespondAlert("Ни одного должника не выбрано")
		return nil
	}

	req := usecase.SplitDebtRequest{
		TelegramCollectorID: okData.CollectorID,
		TelegramDebtorIDs:   debtorIDs,
		TelegramChatID:      c.Chat().ID,
		TotalAmount:         okData.Amount,
		Description:         okData.Description,
	}

	if err := s.Usecase.SplitDebt(context.TODO(), req); err != nil {
		return fmt.Errorf("Error in usecase split debt: %w", err)
	}

	notification := fmt.Sprintf("Новый долг в %d рублей распределен между", req.TotalAmount)
	for i := range debtorIDs {
		notification += fmt.Sprintf(" [%s](tg://user?id=%d)", debtorScreenames[i], debtorIDs[i])
	}

	c.Send(notification, &telebot.SendOptions{
		ParseMode: telebot.ModeMarkdown,
	})

	return nil
}

func getOkData(c telebot.Context) (okBtnData, error) {
	keyboard := c.Callback().Message.ReplyMarkup.InlineKeyboard
	okDataStr := keyboard[len(keyboard)-1][0].Data
	okData := okBtnData{}
	if err := json.Unmarshal([]byte(okDataStr), &okData); err != nil {
		return okBtnData{}, fmt.Errorf("Can't parse ok button data: %w", err)
	}

	return okData, nil
}

func (s *Service) bindSplitHandlers() {
	s.Bot.Handle("/split", s.Split)
	s.Bot.Handle(telebot.OnCallback, s.onCallback)
}
