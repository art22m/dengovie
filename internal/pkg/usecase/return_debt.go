package usecase

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"

	"github.com/art22m/dengovie/internal/pkg/models"
)

type ReturnDebtRequest struct {
	CollectorID int64
	DebtorID    int64
	ChatID      int64
	Amount      int64
	Description string
}

func (uc *UseCase) ReturnDebt(ctx context.Context, req ReturnDebtRequest) error {
	collector := &models.Debt{
		CollectorID: req.CollectorID,
		DebtorID:    req.DebtorID,
		ChatID:      req.ChatID,
	}

	debtor := &models.Debt{
		CollectorID: req.DebtorID,
		DebtorID:    req.CollectorID,
		ChatID:      req.ChatID,
	}

	event := &models.Event{
		CollectorID: req.CollectorID,
		DebtorID:    req.DebtorID,
		ChatID:      req.ChatID,
		Amount:      -req.Amount,
		Description: req.Description,
	}

	err := pgx.BeginFunc(ctx, uc.db, func(tx pgx.Tx) error {
		existingDebt, err := uc.debtsRepo.GetTX(ctx, tx, req.CollectorID, req.DebtorID, req.ChatID)
		if err != nil {
			return nil
		}

		collector.Amount = existingDebt.Amount - req.Amount
		_, err = uc.debtsRepo.UpdateTX(ctx, tx, collector)
		if err != nil {
			return errors.Wrap(err, "failed to update debt")
		}

		debtor.Amount = -collector.Amount
		_, err = uc.debtsRepo.UpdateTX(ctx, tx, debtor)
		if err != nil {
			return errors.Wrap(err, "failed to update debt")
		}

		if err = uc.eventsRepo.CreateTX(ctx, tx, event); err != nil {
			return errors.Wrap(err, "failed to create event")
		}

		return nil
	})

	if err != nil {
		return errors.Wrap(err, "failed to add debt")
	}

	return nil
}
