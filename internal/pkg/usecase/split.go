package usecase

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
)

type SplitDebtRequest struct {
	CollectorID int64
	DebtorIDs   []int64
	ChatID      int64
	TotalAmount int64
	Description string
}

func (uc *UseCase) SplitDebt(ctx context.Context, req SplitDebtRequest) error {
	amountForEach := req.TotalAmount / int64(len(req.DebtorIDs))
	err := pgx.BeginFunc(ctx, uc.db, func(tx pgx.Tx) error {
		for _, debtorID := range req.DebtorIDs {
			if debtorID == req.CollectorID {
				continue
			}
			addReq := AddDebtRequest{
				CollectorID: req.CollectorID,
				DebtorID:    debtorID,
				ChatID:      req.ChatID,
				Amount:      amountForEach,
				Description: req.Description,
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
