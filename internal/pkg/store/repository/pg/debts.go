package pg

import (
	"context"

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
	q := "INSERT INTO debts(collector_id, debtor_id, chat_id, amount) VALUES ($1, $2, $3, $4) ON CONFLICT DO NOTHING"
	_, err := r.db.Exec(
		ctx, q,
		debt.CollectorID, debt.DebtorID, debt.ChatID, debt.Amount,
	)
	return err
}

func (r *DebtRepository) Update(ctx context.Context, debt *models.Debt) (bool, error) {
	q := "UPDATE debts SET amount = $1, updated_at = NOW() WHERE collector_id = $2 AND debtor_id = $3 AND chat_id = $4"
	result, err := r.db.Exec(
		ctx, q,
		debt.Amount, debt.CollectorID, debt.DebtorID, debt.ChatID,
	)
	return result.RowsAffected() > 0, err
}

func (r *DebtRepository) ListByChatID(ctx context.Context, chatID int64) ([]*models.Debt, error) {
	debts := make([]*models.Debt, 0)
	err := r.db.Select(ctx, &debts,
		"SELECT collector_id, debtor_id, chat_id, amount, updated_at, created_at  FROM debts WHERE chat_id = $1",
		chatID,
	)
	return debts, err
}

func (r *DebtRepository) ListByCollectorID(ctx context.Context, collectorID int64) ([]*models.Debt, error) {
	debts := make([]*models.Debt, 0)
	err := r.db.Select(ctx, &debts,
		"SELECT collector_id, debtor_id, chat_id, amount, updated_at, created_at  FROM debts WHERE collector_id = $1",
		collectorID,
	)
	return debts, err
}

func (r *DebtRepository) ListByDebtorID(ctx context.Context, debtorID int64) ([]*models.Debt, error) {
	debts := make([]*models.Debt, 0)
	err := r.db.Select(ctx, &debts,
		"SELECT collector_id, debtor_id, chat_id, amount, updated_at, created_at  FROM debts WHERE debtor_id = $1",
		debtorID,
	)
	return debts, err
}

func (r *DebtRepository) Delete(ctx context.Context, collectorID, debtorID, chatID int64) (bool, error) {
	q := "DELETE FROM debts WHERE collector_id = $2 AND debtor_id = $3 AND chat_id = $4"
	result, err := r.db.Exec(
		ctx, q, collectorID, debtorID, chatID,
	)
	return result.RowsAffected() > 0, err
}
