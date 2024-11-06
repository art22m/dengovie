package usecase

import (
	"context"
	"time"
)

type UserInfo struct {
	TelegramUserID string    `json:"tg_user_id" db:"tg_user_id"`
	PhoneNumber    string    `json:"phone_number" db:"phone_number"`
	TelegramAlias  *string   `json:"tg_alias" db:"tg_alias"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
}

func (uc *UseCase) ListAllUsers(ctx context.Context) ([]UserInfo, error) {
	users, err := uc.usersRepo.List(ctx)
	if err != nil {
		return nil, err
	}

	res := make([]UserInfo, 0, len(users))
	for _, user := range users {
		res = append(res, UserInfo{
			TelegramUserID: user.TelegramUserID,
			PhoneNumber:    user.PhoneNumber,
			TelegramAlias:  user.TelegramAlias,
			CreatedAt:      user.CreatedAt,
		})
	}
	return res, nil
}
