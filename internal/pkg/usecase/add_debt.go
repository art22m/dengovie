package usecase

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"

	"github.com/art22m/dengovie/internal/pkg/models"
	"github.com/art22m/dengovie/internal/pkg/store"
)

type AddDebtRequest struct {
	CollectorID int64
	DebtorID    int64
	ChatID      int64
	Amount      int64
	Description string
}

func (uc *UseCase) AddDebt(ctx context.Context, req AddDebtRequest, logEvent bool) error {
	err := pgx.BeginFunc(ctx, uc.db, func(tx pgx.Tx) error {
		return uc.addDebtTX(ctx, tx, req, logEvent)
	})
	if err != nil {
		return errors.Wrap(err, "failed to add debt")
	}

	return nil
}

func (uc *UseCase) addDebtTX(ctx context.Context, tx pgx.Tx, req AddDebtRequest, logEvent bool) error {
	collectorUser, err := uc.usersRepo.Get(ctx, req.CollectorID)
	switch {
	case errors.Is(err, store.UserNotFound):
		return ErrUserNotFound
	case err != nil:
		return errors.Wrap(err, "failed to get user")
	}

	debtorUser, err := uc.usersRepo.Get(ctx, req.DebtorID)
	switch {
	case errors.Is(err, store.UserNotFound):
		return ErrUserNotFound
	case err != nil:
		return errors.Wrap(err, "failed to get user")
	}

	collector := &models.Debt{
		CollectorID: collectorUser.UserID,
		DebtorID:    debtorUser.UserID,
		ChatID:      req.ChatID,
		Amount:      req.Amount,
	}

	debtor := &models.Debt{
		CollectorID: debtorUser.UserID,
		DebtorID:    collectorUser.UserID,
		ChatID:      req.ChatID,
		Amount:      -req.Amount,
	}

	event := &models.Event{
		CollectorID: collectorUser.UserID,
		DebtorID:    debtorUser.UserID,
		ChatID:      req.ChatID,
		Amount:      req.Amount,
		Description: req.Description,
	}

	existingDebt, err := uc.debtsRepo.GetTX(ctx, tx, collectorUser.UserID, debtorUser.UserID, req.ChatID)
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

	if err = uc.eventsRepo.CreateTX(ctx, tx, event); err != nil {
		return errors.Wrap(err, "failed to create event")
	}

	return nil
}
