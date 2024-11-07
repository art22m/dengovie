package usecase

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
)

type SplitDebtRequest struct {
	TelegramCollectorID int64
	TelegramDebtorIDs   []int64
	TelegramChatID      int64
	TotalAmount         int64
	Description         string
}

func (uc *UseCase) SplitDebt(ctx context.Context, req SplitDebtRequest) error {
	amountForEach := req.TotalAmount * 100 / int64(len(req.TelegramDebtorIDs))
	err := pgx.BeginFunc(ctx, uc.db, func(tx pgx.Tx) error {
		for _, debtorID := range req.TelegramDebtorIDs {
			if debtorID == req.TelegramCollectorID {
				continue
			}
			addReq := AddDebtRequest{
				TelegramCollectorID: req.TelegramCollectorID,
				TelegramDebtorID:    debtorID,
				TelegramChatID:      req.TelegramChatID,
				Amount:              amountForEach,
				Description:         req.Description,
			}
			err := uc.addDebtTX(ctx, tx, addReq)
			if err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		return errors.Wrap(err, "failed to split debt")
	}
	return nil
}
