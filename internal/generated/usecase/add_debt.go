package usecase

//
//import (
//	"context"
//	"time"
//
//	"github.com/art22m/dengovie/internal/generated/dengovie/dengovie/public/model"
//	"github.com/jackc/pgx/v5"
//	"github.com/pkg/errors"
//	"github.com/samber/lo"
//)
//
//type AddDebtRequest struct {
//	CollectorID int64
//	DebtorID    int64
//	ChatID      int64
//	Amount      int64
//	Description *string
//}
//
//func (uc *UseCase) AddDebt(ctx context.Context, req AddDebtRequest) error {
//	collector := model.Debts{
//		CollectorID: req.CollectorID,
//		DebtorID:    req.DebtorID,
//		ChatID:      req.ChatID,
//		Amount:      lo.ToPtr(req.Amount),
//		UpdatedAt:   time.Now(),
//		CreatedAt:   time.Now(),
//	}
//
//	debtor := model.Debts{
//		CollectorID: req.DebtorID,
//		DebtorID:    req.CollectorID,
//		ChatID:      req.ChatID,
//		Amount:      lo.ToPtr(-req.Amount),
//		UpdatedAt:   time.Now(),
//		CreatedAt:   time.Now(),
//	}
//
//	event := model.Events{
//		CollectorID: req.CollectorID,
//		DebtorID:    req.DebtorID,
//		ChatID:      req.ChatID,
//		Amount:      lo.ToPtr(req.Amount),
//		Description: req.Description,
//		CreatedAt:   time.Now(),
//	}
//
//	if err := pgx.BeginFunc(ctx, uc.db, func(tx pgx.Tx) error {
//		existingDebt, err := uc.debtRepo.GetForUpdate(ctx, tx, req.CollectorID, req.DebtorID, req.ChatID)
//		if err != nil {
//			return err
//		}
//
//		switch existingDebt {
//		// Если долга не существует -> нужно его создать
//		case nil:
//			err = uc.debtRepo.Create(ctx, tx, collector)
//			if err != nil {
//				return errors.Wrap(err, "failed to create collector debt")
//			}
//
//			err = uc.debtRepo.Create(ctx, tx, debtor)
//			if err != nil {
//				return errors.Wrap(err, "failed to create debtor debt")
//			}
//		// Иначе, обновляем его
//		default:
//			collector.Amount = lo.ToPtr(lo.FromPtr(existingDebt.Amount) + req.Amount)
//			err = uc.debtRepo.Update(ctx, tx, collector)
//			if err != nil {
//				return errors.Wrap(err, "failed to update collector debt")
//			}
//
//			debtor.Amount = lo.ToPtr(lo.FromPtr(existingDebt.Amount) - req.Amount)
//			err = uc.debtRepo.Update(ctx, tx, debtor)
//			if err != nil {
//				return errors.Wrap(err, "failed to update debtor debt")
//			}
//		}
//
//		// Создаем запись в таблице events в любом случае
//		if err = uc.eventRepo.Create(ctx, tx, event); err != nil {
//			return errors.Wrap(err, "failed to create event")
//		}
//
//		return nil
//	}); err != nil {
//		return errors.Wrap(err, "failed to add debt")
//	}
//
//	return nil
//}
