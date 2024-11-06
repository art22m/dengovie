package usecase

import (
	"context"
	"time"

	"github.com/art22m/dengovie/internal/generated/dengovie/dengovie/public/model"
	"github.com/pkg/errors"
	"github.com/samber/lo"
)

type RegisterUserRequest struct {
	TelegramUserID string
	PhoneNumber    string
	TelegramAlias  *string
}

func (uc *UseCase) Register(ctx context.Context, req RegisterUserRequest) error {
	user := model.Users{
		TgUserID:    req.TelegramUserID,
		PhoneNumber: req.PhoneNumber,
		TgAlias:     lo.FromPtr(req.TelegramAlias),
		CreatedAt:   time.Now(),
	}

	if err := uc.userRepo.Create(ctx, user); err != nil {
		return errors.Wrap(err, "failed to register user")
	}

	return nil
}
