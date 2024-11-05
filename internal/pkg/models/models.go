package models

import "time"

type User struct {
	UserID         int64     `json:"user_id" db:"user_id"`
	TelegramUserID string    `json:"tg_user_id" db:"tg_user_id"`
	PhoneNumber    string    `json:"phone_number" db:"phone_number"`
	TelegramAlias  *string   `json:"tg_alias" db:"tg_alias"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
}

type Chat struct {
	ChatID         int64     `json:"chat_id" db:"chat_id"`
	TelegramChatID string    `json:"tg_chat_id" db:"tg_chat_id"`
	Description    string    `json:"description" db:"description"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
}

type Debt struct {
	CollectorID int64     `json:"collector_id" db:"collector_id"`
	DebtorID    int64     `json:"debtor_id" db:"debtor_id"`
	ChatID      int64     `json:"chat_id" db:"chat_id"`
	Amount      int64     `json:"amount" db:"amount"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

type Event struct {
	EventID     int64     `json:"event_id" db:"event_id"`
	CollectorID int64     `json:"collector_id" db:"collector_id"`
	DebtorID    int64     `json:"debtor_id" db:"debtor_id"`
	ChatID      int64     `json:"chat_id" db:"chat_id"`
	Amount      int64     `json:"amount" db:"amount"`
	Description string    `json:"description" db:"description"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}
