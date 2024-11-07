package usecase

import (
	"context"

	"github.com/pkg/errors"

	"github.com/art22m/dengovie/internal/pkg/store"
)

func (uc *UseCase) GetUser(ctx context.Context, telegramID string) (*UserInfo, error) {
	user, err := uc.usersRepo.GetByTelegramUserID(ctx, telegramID)
	switch {
	case errors.Is(err, store.UserNotFound):
		return nil, ErrUserNotFound
	case err != nil:
		return nil, errors.Wrap(err, "failed to get user")
	}
	info := &UserInfo{
		TelegramUserID: user.TelegramUserID,
		PhoneNumber:    user.PhoneNumber,
		TelegramAlias:  user.TelegramAlias,
		CreatedAt:      user.CreatedAt,
	}
	return info, nil
}
