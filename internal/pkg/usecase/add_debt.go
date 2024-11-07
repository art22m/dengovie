package usecase

import (
	"context"
	"strconv"

	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"

	"github.com/art22m/dengovie/internal/pkg/models"
	"github.com/art22m/dengovie/internal/pkg/store"
)

type AddDebtRequest struct {
	TelegramCollectorID int64
	TelegramDebtorID    int64
	TelegramChatID      int64
	Amount              int64
	Description         string
}

func (uc *UseCase) AddDebt(ctx context.Context, req AddDebtRequest) error {
	err := pgx.BeginFunc(ctx, uc.db, func(tx pgx.Tx) error {
		return uc.addDebtTX(ctx, tx, req)
	})
	if err != nil {
		return errors.Wrap(err, "failed to add debt")
	}

	return nil
}

func (uc *UseCase) addDebtTX(ctx context.Context, tx pgx.Tx, req AddDebtRequest) error {
	chat, err := uc.chatsRepo.GetByTelegramChatID(ctx, strconv.FormatInt(req.TelegramChatID, 10))
	switch {
	case errors.Is(err, store.ChatNotFound):
		return ErrChatNotFound
	case err != nil:
		return errors.Wrap(err, "failed to get chat")
	}

	collectorUser, err := uc.usersRepo.GetByTelegramUserID(ctx, strconv.FormatInt(req.TelegramCollectorID, 10))
	switch {
	case errors.Is(err, store.UserNotFound):
		return ErrUserNotFound
	case err != nil:
		return errors.Wrap(err, "failed to get user")
	}

	debtorUser, err := uc.usersRepo.GetByTelegramUserID(ctx, strconv.FormatInt(req.TelegramDebtorID, 10))
	switch {
	case errors.Is(err, store.UserNotFound):
		return ErrUserNotFound
	case err != nil:
		return errors.Wrap(err, "failed to get user")
	}

	collector := &models.Debt{
		CollectorID: collectorUser.UserID,
		DebtorID:    debtorUser.UserID,
		ChatID:      chat.ChatID,
		Amount:      req.Amount,
	}

	debtor := &models.Debt{
		CollectorID: debtorUser.UserID,
		DebtorID:    collectorUser.UserID,
		ChatID:      chat.ChatID,
		Amount:      -req.Amount,
	}

	event := &models.Event{
		CollectorID: collectorUser.UserID,
		DebtorID:    debtorUser.UserID,
		ChatID:      chat.ChatID,
		Amount:      req.Amount,
		Description: req.Description,
	}

	existingDebt, err := uc.debtsRepo.GetTX(ctx, tx, collectorUser.UserID, debtorUser.UserID, chat.ChatID)
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
