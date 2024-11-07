package usecase

import (
	"context"

	"github.com/pkg/errors"

	"github.com/art22m/dengovie/internal/pkg/store"
)

func (uc *UseCase) GetUser(ctx context.Context, id int64) (*UserInfo, error) {
	user, err := uc.usersRepo.Get(ctx, id)
	switch {
	case errors.Is(err, store.UserNotFound):
		return nil, ErrUserNotFound
	case err != nil:
		return nil, errors.Wrap(err, "failed to get user")
	}
	info := &UserInfo{
		UserID:      user.UserID,
		PhoneNumber: user.PhoneNumber,
		Alias:       user.Alias,
		CreatedAt:   user.CreatedAt,
	}
	return info, nil
}
