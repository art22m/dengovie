package usecase

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
	"github.com/samber/lo"

	"github.com/art22m/dengovie/internal/pkg/models"
	"github.com/art22m/dengovie/internal/pkg/store"
)

type AddDebtRequest struct {
	CollectorID int64
	DebtorID    int64
	ChatID      int64
	Amount      int64
	Description *string
}

func (uc *UseCase) AddDebt(ctx context.Context, req AddDebtRequest) error {
	collector := &models.Debt{
		CollectorID: req.CollectorID,
		DebtorID:    req.DebtorID,
		ChatID:      req.ChatID,
		Amount:      req.Amount,
	}

	debtor := &models.Debt{
		CollectorID: req.DebtorID,
		DebtorID:    req.CollectorID,
		ChatID:      req.ChatID,
		Amount:      -req.Amount,
	}

	event := &models.Event{
		CollectorID: req.CollectorID,
		DebtorID:    req.DebtorID,
		ChatID:      req.ChatID,
		Amount:      req.Amount,
		Description: lo.FromPtr(req.Description),
	}

	err := pgx.BeginFunc(ctx, uc.db, func(tx pgx.Tx) error {
		existingDebt, err := uc.debtsRepo.GetTX(ctx, tx, req.CollectorID, req.DebtorID, req.ChatID)
		switch {
		case err == nil:
		case errors.Is(err, store.DebtNotFound):
		default:
			return err
		}

		if errors.Is(err, store.DebtNotFound) {
			err = uc.debtsRepo.CreateTX(ctx, tx, collector)
			if err != nil {
				return errors.Wrap(err, "failed to create debt")
			}

			err = uc.debtsRepo.CreateTX(ctx, tx, debtor)
			if err != nil {
				return errors.Wrap(err, "failed to create debt")
			}
		} else {
			collector.Amount = existingDebt.Amount + req.Amount
			_, err = uc.debtsRepo.UpdateTX(ctx, tx, collector)
			if err != nil {
				return errors.Wrap(err, "failed to update debt")
			}

			debtor.Amount = -collector.Amount
			_, err = uc.debtsRepo.UpdateTX(ctx, tx, debtor)
			if err != nil {
				return errors.Wrap(err, "failed to update debt")
			}
		}

		fmt.Println("add event")
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
