package register_user

import (
	"context"
	"time"

	"github.com/art22m/dengovie/internal/generated/dengovie/dengovie/public/model"
	"github.com/pkg/errors"
	"github.com/samber/lo"
)

type RegisterUserRequest struct {
	TelegramUserID string  `json:"tg_user_id"`
	PhoneNumber    string  `json:"phone_number"`
	TelegramAlias  *string `json:"tg_alias"`
}

type UserRepo interface {
	Create(ctx context.Context, user model.Users) error
}

type UseCase struct {
	repo UserRepo
}

func NewUseCase(repo UserRepo) *UseCase {
	return &UseCase{repo: repo}
}

func (uc *UseCase) Register(ctx context.Context, req RegisterUserRequest) error {
	user := model.Users{
		TgUserID:    req.TelegramUserID,
		PhoneNumber: req.PhoneNumber,
		TgAlias:     lo.FromPtr(req.TelegramAlias),
		CreatedAt:   time.Now(),
	}

	if err := uc.repo.Create(ctx, user); err != nil {
		return errors.Wrap(err, "failed to register user")
	}

	return nil
}
