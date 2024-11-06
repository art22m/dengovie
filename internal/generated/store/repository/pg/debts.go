package pg

//
//import (
//	"context"
//	"time"
//
//	"github.com/pkg/errors"
//
//	"github.com/go-jet/jet/v2/postgres"
//	"github.com/jackc/pgx/v5"
//
//	"github.com/art22m/dengovie/internal/generated/dengovie/dengovie/public/model"
//	"github.com/art22m/dengovie/internal/generated/dengovie/dengovie/public/table"
//	"github.com/art22m/dengovie/internal/pkg/models"
//	"github.com/art22m/dengovie/internal/pkg/store"
//)
//
//type DebtRepository struct {
//	db store.DatabaseOperations
//}
//
//func NewDebts(db store.DatabaseOperations) *DebtRepository {
//	return &DebtRepository{
//		db: db,
//	}
//}
//
//func (r *DebtRepository) GetForUpdate(ctx context.Context, tx pgx.Tx, collectorID, debtorID, chatID int64) (*model.Debts, error) {
//	stmt, args := table.Debts.
//		SELECT(table.Debts.AllColumns).
//		WHERE(
//			table.Debts.CollectorID.EQ(postgres.Int64(collectorID)).
//				AND(table.Debts.DebtorID.EQ(postgres.Int64(debtorID)).
//					AND(table.Debts.ChatID.EQ(postgres.Int64(chatID)))),
//		).
//		FOR(postgres.UPDATE().SKIP_LOCKED()).
//		Sql()
//
//	var debt *model.Debts
//	row := tx.QueryRow(ctx, stmt, args...)
//	if err := scanDebt(row, debt); err != nil {
//		if errors.Is(err, pgx.ErrNoRows) {
//			return nil, nil
//		}
//
//		return nil, errors.Wrap(err, "failed to scan debt")
//	}
//
//	return debt, nil
//}
//
//func (r *DebtRepository) Create(ctx context.Context, tx pgx.Tx, debt model.Debts) error {
//	stmt, args := table.Debts.
//		INSERT().
//		MODEL(debt).
//		Sql()
//
//	_, err := tx.Exec(ctx, stmt, args...)
//	return err
//}
//
//func (r *DebtRepository) Update(ctx context.Context, tx pgx.Tx, debt model.Debts) error {
//	stmt, args := table.Debts.
//		UPDATE(
//			table.Debts.Amount,
//			table.Debts.UpdatedAt,
//		).
//		SET(
//			debt.Amount,
//			time.Now(),
//		).
//		Sql()
//
//	_, err := tx.Exec(ctx, stmt, args...)
//	return err
//}
//
//func (r *DebtRepository) ListByChatID(ctx context.Context, chatID int64) ([]*models.Debt, error) {
//	var debts []*models.Debt
//
//	stmt, args := table.Debts.
//		SELECT(table.Debts.AllColumns).
//		WHERE(table.Debts.ChatID.EQ(postgres.Int64(chatID))).
//		Sql()
//
//	err := r.db.Select(ctx, &debts, stmt, args...)
//	switch {
//	case errors.Is(err, pgx.ErrNoRows):
//		return nil, nil
//	case err != nil:
//		return nil, err
//	}
//
//	return debts, nil
//}
//
//func (r *DebtRepository) ListByCollectorID(ctx context.Context, collectorID int64) ([]*models.Debt, error) {
//	var debts []*models.Debt
//
//	stmt, args := table.Debts.
//		SELECT(table.Debts.AllColumns).
//		WHERE(table.Debts.CollectorID.EQ(postgres.Int64(collectorID))).
//		Sql()
//
//	err := r.db.Select(ctx, &debts, stmt, args...)
//	switch {
//	case errors.Is(err, pgx.ErrNoRows):
//		return nil, nil
//	case err != nil:
//		return nil, err
//	}
//
//	return debts, nil
//}
//
//func (r *DebtRepository) ListByDebtorID(ctx context.Context, debtorID int64) ([]*models.Debt, error) {
//	var debts []*models.Debt
//
//	stmt, args := table.Debts.
//		SELECT(table.Debts.AllColumns).
//		WHERE(table.Debts.DebtorID.EQ(postgres.Int64(debtorID))).
//		Sql()
//
//	err := r.db.Select(ctx, &debts, stmt, args...)
//	switch {
//	case errors.Is(err, pgx.ErrNoRows):
//		return nil, nil
//	case err != nil:
//		return nil, err
//	}
//
//	return debts, nil
//}
//
//func (r *DebtRepository) Delete(ctx context.Context, collectorID, debtorID, chatID int64) (bool, error) {
//	stmt, args := table.Debts.
//		DELETE().
//		WHERE(
//			table.Debts.CollectorID.EQ(postgres.Int64(collectorID)).AND(
//				table.Debts.DebtorID.EQ(postgres.Int64(debtorID))).AND(
//				table.Debts.ChatID.EQ(postgres.Int64(chatID))),
//		).
//		Sql()
//
//	result, err := r.db.Exec(ctx, stmt, args...)
//
//	return result.RowsAffected() > 0, err
//}
//
//func scanDebt(row pgx.Row, dest *model.Debts) error {
//	return row.Scan(
//		&dest.CollectorID,
//		&dest.DebtorID,
//		&dest.ChatID,
//		&dest.Amount,
//		&dest.UpdatedAt,
//		&dest.CreatedAt,
//	)
//}
