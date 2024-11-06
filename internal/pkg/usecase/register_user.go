package usecase

import (
	"context"

	"github.com/pkg/errors"

	"github.com/art22m/dengovie/internal/pkg/models"
	"github.com/art22m/dengovie/internal/pkg/store"
)

type RegisterUserRequest struct {
	TelegramUserID string
	PhoneNumber    string
	TelegramAlias  *string
}

func (uc *UseCase) RegisterUser(ctx context.Context, req RegisterUserRequest) error {
	_, err := uc.usersRepo.GetByTelegramUserID(ctx, req.TelegramUserID)
	switch {
	case errors.Is(err, store.UserNotFound):
		// ok
	case err != nil:
		return errors.Wrap(err, "failed to check user")
	case err == nil:
		return ErrUserAlreadyExists
	}

	user := &models.User{
		TelegramUserID: req.TelegramUserID,
		PhoneNumber:    req.PhoneNumber,
		TelegramAlias:  req.TelegramAlias,
	}
	if err = uc.usersRepo.Create(ctx, user); err != nil {
		return errors.Wrap(err, "failed to register user")
	}
	return nil
}
