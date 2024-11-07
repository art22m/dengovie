package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"

	"github.com/art22m/dengovie/internal/pkg/store"
)

type ListDebtsRequest struct {
	TelegramUserID string
	TelegramChatID string
}

type ListDebtsResponse struct {
	CollectorTelegramID    string
	CollectorTelegramAlias *string
	DebtsInfo              []DebtInfo
}

type DebtInfo struct {
	DebtorTelegramID    string
	DebtorTelegramAlias *string

	Amount       int64
	LastModified time.Time
}

func (uc *UseCase) ListDebts(ctx context.Context, req ListDebtsRequest) (*ListDebtsResponse, error) {
	chat, err := uc.chatsRepo.GetByTelegramChatID(ctx, req.TelegramChatID)
	switch {
	case errors.Is(err, store.ChatNotFound):
		return nil, ErrChatNotFound
	case err != nil:
		return nil, errors.Wrap(err, "failed to get chat")
	}

	collector, err := uc.usersRepo.GetByTelegramUserID(ctx, req.TelegramUserID)
	switch {
	case errors.Is(err, store.UserNotFound):
		return nil, ErrUserNotFound
	case err != nil:
		return nil, errors.Wrap(err, "failed to get user")
	}

	debts, err := uc.debtsRepo.List(ctx, &collector.UserID, &chat.ChatID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get collectors debts")
	}

	resp := ListDebtsResponse{
		CollectorTelegramID:    collector.TelegramUserID,
		CollectorTelegramAlias: collector.TelegramAlias,
		DebtsInfo:              make([]DebtInfo, len(debts)),
	}
	for i, debt := range debts {
		debtor, err := uc.usersRepo.Get(ctx, debt.DebtorID)
		switch {
		case errors.Is(err, store.UserNotFound):
			return nil, errors.Wrap(ErrUserNotFound, fmt.Sprintf("inernal state error with debtor_id=%v", debt.DebtorID))
		case err != nil:
			return nil, errors.Wrap(err, "failed to get userr")
		}

		resp.DebtsInfo[i] = DebtInfo{
			DebtorTelegramID:    debtor.TelegramUserID,
			DebtorTelegramAlias: debtor.TelegramAlias,
			Amount:              debt.Amount,
			LastModified:        debt.UpdatedAt,
		}
	}

	return &resp, nil
}
