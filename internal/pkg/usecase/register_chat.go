package usecase

import (
	"context"

	"github.com/pkg/errors"

	"github.com/art22m/dengovie/internal/pkg/models"
	"github.com/art22m/dengovie/internal/pkg/store"
)

type RegisterChatRequest struct {
	TelegramChatID string
	Info           string
}

func (uc *UseCase) RegisterChat(ctx context.Context, req RegisterChatRequest) error {
	_, err := uc.chatsRepo.GetByTelegramChatID(ctx, req.TelegramChatID)
	switch {
	case errors.Is(err, store.ChatNotFound):
		// ok
	case err != nil:
		return errors.Wrap(err, "failed to check user")
	case err == nil:
		return ErrChatAlreadyExists
	}

	chat := &models.Chat{
		TelegramChatID: req.TelegramChatID,
		Description:    req.Info,
	}
	if err = uc.chatsRepo.Create(ctx, chat); err != nil {
		return errors.Wrap(err, "failed to register user")
	}
	return nil
}
