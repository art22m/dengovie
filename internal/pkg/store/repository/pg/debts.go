package pg

import (
	"context"
	"time"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
	"github.com/samber/lo"

	"github.com/art22m/dengovie/internal/generated/dengovie/dengovie/public/model"
	"github.com/art22m/dengovie/internal/generated/dengovie/dengovie/public/table"
	"github.com/art22m/dengovie/internal/pkg/models"
	"github.com/art22m/dengovie/internal/pkg/store"
)

type DebtRepository struct {
	db store.DatabaseOperations
}

func NewDebts(db store.DatabaseOperations) *DebtRepository {
	return &DebtRepository{
		db: db,
	}
}

func (r *DebtRepository) Create(ctx context.Context, debt *models.Debt) error {
	jetDebt := model.Debts{
		CollectorID: debt.CollectorID,
		DebtorID:    debt.ChatID,
		ChatID:      debt.ChatID,
		Amount:      lo.ToPtr(debt.Amount),
		UpdatedAt:   debt.UpdatedAt,
		CreatedAt:   debt.CreatedAt,
	}

	stmt, args := table.Debts.INSERT().MODEL(jetDebt).Sql()

	_, err := r.db.Exec(ctx, stmt, args...)
	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return nil
	case err != nil:
		return err
	}

	return nil
}

func (r *DebtRepository) ListByChatID(ctx context.Context, chatID int64) ([]*models.Debt, error) {
	var debts []*models.Debt

	stmt, args := table.Debts.
		SELECT(table.Debts.AllColumns).
		WHERE(table.Debts.ChatID.EQ(postgres.Int64(chatID))).
		Sql()

	err := r.db.Select(ctx, &debts, stmt, args...)
	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return nil, nil
	case err != nil:
		return nil, err
	}

	return debts, nil
}

func (r *DebtRepository) ListByCollectorID(ctx context.Context, collectorID int64) ([]*models.Debt, error) {
	var debts []*models.Debt

	stmt, args := table.Debts.
		SELECT(table.Debts.AllColumns).
		WHERE(table.Debts.CollectorID.EQ(postgres.Int64(collectorID))).
		Sql()

	err := r.db.Select(ctx, &debts, stmt, args...)
	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return nil, nil
	case err != nil:
		return nil, err
	}

	return debts, nil
}

func (r *DebtRepository) ListByDebtorID(ctx context.Context, debtorID int64) ([]*models.Debt, error) {
	var debts []*models.Debt

	stmt, args := table.Debts.
		SELECT(table.Debts.AllColumns).
		WHERE(table.Debts.DebtorID.EQ(postgres.Int64(debtorID))).
		Sql()

	err := r.db.Select(ctx, &debts, stmt, args...)
	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return nil, nil
	case err != nil:
		return nil, err
	}

	return debts, nil
}

func (r *DebtRepository) Update(ctx context.Context, debt *models.Debt) (bool, error) {
	stmt, args := table.Debts.
		UPDATE(
			table.Debts.Amount,
			table.Debts.UpdatedAt,
		).
		SET(
			debt.Amount,
			time.Now(),
		).
		Sql()

	result, err := r.db.Exec(ctx, stmt, args...)
	if errors.Is(err, pgx.ErrNoRows) {
		return false, nil
	}

	return result.RowsAffected() > 0, err
}

func (r *DebtRepository) Delete(ctx context.Context, collectorID, debtorID, chatID int64) (bool, error) {
	stmt, args := table.Debts.
		DELETE().
		WHERE(
			table.Debts.CollectorID.EQ(postgres.Int64(collectorID)).AND(
				table.Debts.DebtorID.EQ(postgres.Int64(debtorID))).AND(
				table.Debts.ChatID.EQ(postgres.Int64(chatID))),
		).
		Sql()

	result, err := r.db.Exec(ctx, stmt, args...)

	return result.RowsAffected() > 0, err
}
