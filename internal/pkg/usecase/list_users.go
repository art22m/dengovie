package usecase

import (
	"context"
	"time"
)

type UserInfo struct {
	UserID      int64
	PhoneNumber string
	Alias       *string
	CreatedAt   time.Time
}

func (uc *UseCase) ListAllUsers(ctx context.Context) ([]UserInfo, error) {
	users, err := uc.usersRepo.List(ctx)
	if err != nil {
		return nil, err
	}

	res := make([]UserInfo, 0, len(users))
	for _, user := range users {
		res = append(res, UserInfo{
			UserID:      user.UserID,
			PhoneNumber: user.PhoneNumber,
			Alias:       user.Alias,
			CreatedAt:   user.CreatedAt,
		})
	}
	return res, nil
}
