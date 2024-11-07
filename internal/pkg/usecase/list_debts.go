package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"

	"github.com/art22m/dengovie/internal/pkg/store"
)

type ListDebtsRequest struct {
	UserID int64
	ChatID int64
}

type ListDebtsResponse struct {
	CollectorTelegramID    int64
	CollectorTelegramAlias *string
	DebtsInfo              []DebtInfo
}

type DebtInfo struct {
	DebtorTelegramID    int64
	DebtorTelegramAlias *string

	Amount       int64
	LastModified time.Time
}

func (uc *UseCase) ListDebts(ctx context.Context, req ListDebtsRequest) (*ListDebtsResponse, error) {
	collector, err := uc.usersRepo.Get(ctx, req.UserID)
	switch {
	case errors.Is(err, store.UserNotFound):
		return nil, ErrUserNotFound
	case err != nil:
		return nil, errors.Wrap(err, "failed to get user")
	}

	debts, err := uc.debtsRepo.List(ctx, &req.UserID, &req.ChatID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get collectors debts")
	}

	resp := ListDebtsResponse{
		CollectorTelegramID:    collector.UserID,
		CollectorTelegramAlias: collector.Alias,
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
			DebtorTelegramID:    debtor.UserID,
			DebtorTelegramAlias: debtor.Alias,
			Amount:              debt.Amount,
			LastModified:        debt.UpdatedAt,
		}
	}

	return &resp, nil
}
