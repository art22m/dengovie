package usecase

//
//import (
//	"context"
//	"time"
//
//	"github.com/jackc/pgx/v5"
//	"github.com/pkg/errors"
//	"github.com/samber/lo"
//
//	"github.com/art22m/dengovie/internal/generated/dengovie/dengovie/public/model"
//)
//
//type SplitDebtRequest struct {
//	CollectorID int64
//	DebtorIDs   []int64
//	ChatID      int64
//	TotalAmount int64
//	Description *string
//}
//
//func (uc *UseCase) SplitDebt(ctx context.Context, req SplitDebtRequest) error {
//	amountForEach := req.TotalAmount / int64(len(req.DebtorIDs))
//
//	if err := pgx.BeginFunc(ctx, uc.db, func(tx pgx.Tx) error {
//		for _, debtorID := range req.DebtorIDs {
//			err := uc.addDebtTx(ctx, tx, req, debtorID, amountForEach)
//			if err != nil {
//				return err
//			}
//		}
//
//		return nil
//	}); err != nil {
//		return errors.Wrap(err, "failed to add debt")
//	}
//
//	return nil
//}
//
//func (uc *UseCase) addDebtTx(ctx context.Context, tx pgx.Tx, req SplitDebtRequest, debtorID int64, amountForEach int64) error {
//	collector := model.Debts{
//		CollectorID: req.CollectorID,
//		DebtorID:    debtorID,
//		ChatID:      req.ChatID,
//		Amount:      lo.ToPtr(amountForEach),
//		UpdatedAt:   time.Now(),
//		CreatedAt:   time.Now(),
//	}
//
//	debtor := model.Debts{
//		CollectorID: debtorID,
//		DebtorID:    req.CollectorID,
//		ChatID:      req.ChatID,
//		Amount:      lo.ToPtr(-amountForEach),
//		UpdatedAt:   time.Now(),
//		CreatedAt:   time.Now(),
//	}
//
//	event := model.Events{
//		CollectorID: req.CollectorID,
//		DebtorID:    debtorID,
//		ChatID:      req.ChatID,
//		Amount:      lo.ToPtr(req.TotalAmount),
//		Description: req.Description,
//		CreatedAt:   time.Now(),
//	}
//	existingDebt, err := uc.debtRepo.GetForUpdate(ctx, tx, req.CollectorID, debtorID, req.ChatID)
//	if err != nil {
//		return err
//	}
//
//	switch existingDebt {
//	// Если долга не существует -> нужно его создать
//	case nil:
//		err = uc.debtRepo.Create(ctx, tx, collector)
//		if err != nil {
//			return errors.Wrap(err, "failed to create collector debt")
//		}
//
//		err = uc.debtRepo.Create(ctx, tx, debtor)
//		if err != nil {
//			return errors.Wrap(err, "failed to create debtor debt")
//		}
//	// Иначе, обновляем его
//	default:
//		collector.Amount = lo.ToPtr(lo.FromPtr(existingDebt.Amount) + amountForEach)
//		err = uc.debtRepo.Update(ctx, tx, collector)
//		if err != nil {
//			return errors.Wrap(err, "failed to update collector debt")
//		}
//
//		debtor.Amount = lo.ToPtr(lo.FromPtr(existingDebt.Amount) - amountForEach)
//		err = uc.debtRepo.Update(ctx, tx, debtor)
//		if err != nil {
//			return errors.Wrap(err, "failed to update debtor debt")
//		}
//	}
//
//	// Создаем запись в таблице events в любом случае
//	if err = uc.eventRepo.Create(ctx, tx, event); err != nil {
//		return errors.Wrap(err, "failed to create event")
//	}
//	return nil
//}
