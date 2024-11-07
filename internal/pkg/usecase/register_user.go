package usecase

import (
	"context"

	"github.com/pkg/errors"

	"github.com/art22m/dengovie/internal/pkg/models"
	"github.com/art22m/dengovie/internal/pkg/store"
)

type RegisterUserRequest struct {
	UserID      int64
	PhoneNumber string
	Alias       *string
}

func (uc *UseCase) RegisterUser(ctx context.Context, req RegisterUserRequest) error {
	_, err := uc.usersRepo.Get(ctx, req.UserID)
	switch {
	case errors.Is(err, store.UserNotFound):
		// ok
	case err != nil:
		return errors.Wrap(err, "failed to check user")
	case err == nil:
		return ErrUserAlreadyExists
	}

	user := &models.User{
		UserID:      req.UserID,
		PhoneNumber: req.PhoneNumber,
		Alias:       req.Alias,
	}
	if err = uc.usersRepo.Create(ctx, user); err != nil {
		return errors.Wrap(err, "failed to register user")
	}
	return nil
}
