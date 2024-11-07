package pg

import (
	"context"

	"github.com/jackc/pgx/v5"

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

func (r *DebtRepository) CreateTX(ctx context.Context, tx pgx.Tx, debt *models.Debt) error {
	q := "INSERT INTO debts(collector_id, debtor_id, chat_id, amount) VALUES ($1, $2, $3, $4) ON CONFLICT DO NOTHING"
	_, err := r.db.ExecTX(
		ctx, tx, q,
		debt.CollectorID, debt.DebtorID, debt.ChatID, debt.Amount,
	)
	return err
}

func (r *DebtRepository) ClearTX(ctx context.Context, tx pgx.Tx, chatID int64) error {
	q := "DELETE FROM debts WHERE chat_id = $1"
	_, err := r.db.ExecTX(
		ctx, tx, q, chatID,
	)
	return err
}

func (r *DebtRepository) GetTX(ctx context.Context, tx pgx.Tx, collectorID, debtorID, chatID int64) (*models.Debt, error) {
	q := "SELECT collector_id, debtor_id, chat_id, amount, updated_at, created_at FROM debts WHERE collector_id = $1 AND debtor_id = $2 AND chat_id = $3 FOR UPDATE"
	debts := make([]*models.Debt, 0)
	err := r.db.SelectTX(ctx, tx, &debts, q, collectorID, debtorID, chatID)
	if err != nil {
		return nil, err
	}
	if len(debts) == 0 {
		return nil, store.DebtNotFound
	}
	return debts[0], err
}

func (r *DebtRepository) UpdateTX(ctx context.Context, tx pgx.Tx, debt *models.Debt) (bool, error) {
	q := "UPDATE debts SET amount = $1, updated_at = NOW() WHERE collector_id = $2 AND debtor_id = $3 AND chat_id = $4"
	result, err := r.db.ExecTX(
		ctx, tx, q,
		debt.Amount, debt.CollectorID, debt.DebtorID, debt.ChatID,
	)
	return result.RowsAffected() > 0, err
}

func (r *DebtRepository) List(ctx context.Context, collectorID *int64, chatID *int64) ([]*models.Debt, error) {
	args := make([]interface{}, 0)
	q := "SELECT collector_id, debtor_id, chat_id, amount, updated_at, created_at FROM debts WHERE 1 = 1"
	switch {
	case collectorID != nil && chatID != nil:
		args = append(args, *collectorID, *chatID)
		q += " AND collector_id = $1 AND chat_id = $2"
	case collectorID != nil:
		args = append(args, *collectorID)
		q += " AND collector_id = $1"
	case chatID != nil:
		args = append(args, *chatID)
		q += " AND chat_id = $1"
	}
	debts := make([]*models.Debt, 0)
	err := r.db.Select(
		ctx, &debts,
		q, args...,
	)
	//fmt.Println(len(debts), lo.FromPtr(collectorID), lo.FromPtr(chatID), q)
	return debts, err
}

func (r *DebtRepository) Delete(ctx context.Context, collectorID, debtorID, chatID int64) (bool, error) {
	q := "DELETE FROM debts WHERE collector_id = $1 AND debtor_id = $2 AND chat_id = $3"
	result, err := r.db.Exec(
		ctx, q, collectorID, debtorID, chatID,
	)
	return result.RowsAffected() > 0, err
}

func (r *DebtRepository) DeleteTX(ctx context.Context, tx pgx.Tx, collectorID, debtorID, chatID int64) (bool, error) {
	q := "DELETE FROM debts WHERE collector_id = $1 AND debtor_id = $2 AND chat_id = $3"
	result, err := r.db.ExecTX(
		ctx, tx, q, collectorID, debtorID, chatID,
	)
	return result.RowsAffected() > 0, err
}
