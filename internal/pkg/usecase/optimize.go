package usecase

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type transfer struct {
	from, to int64
	amount   int64
}

type balance struct {
	id     int64
	amount int64
}

func selectCandidates(balances map[int64]int64) (ok bool, negative, positive balance) {
	if len(balances) == 0 {
		return
	}
	for k, v := range balances {
		if v > 0 {
			positive = balance{
				id:     k,
				amount: v,
			}
		} else {
			negative = balance{
				id:     k,
				amount: v,
			}
		}
	}
	ok = true
	return
}

func deleteZeroBalances(balances map[int64]int64) {
	for k, v := range balances {
		if v == 0 {
			delete(balances, k)
		}
	}
}

// OptimizeDebts https://www.win.tue.nl/~wstomv/publications/settling-debts-problems.pdf
func (uc *UseCase) OptimizeDebts(ctx context.Context, chatID int64) (fromCount int, toCount int, err error) {
	debts, err := uc.debtsRepo.List(ctx, nil, &chatID)
	if err != nil {
		return 0, 0, err
	}

	balances := make(map[int64]int64, len(debts)/2)
	for _, debt := range debts {
		balances[debt.CollectorID] += debt.Amount
	}

	for k, v := range balances {
		fmt.Printf("user=%v, balance=%v \n", k, v)
	}

	deleteZeroBalances(balances)

	transfers := make([]transfer, 0)
	for {
		ok, neg, pos := selectCandidates(balances)
		if !ok {
			break
		}

		m := min(-neg.amount, pos.amount)
		transfers = append(transfers, transfer{
			from:   neg.id,
			to:     pos.id,
			amount: m,
		})

		balances[neg.id] += m
		balances[pos.id] -= m

		deleteZeroBalances(balances)
	}

	err = pgx.BeginFunc(ctx, uc.db, func(tx pgx.Tx) error {
		txErr := uc.debtsRepo.ClearTX(ctx, tx, chatID)
		if txErr != nil {
			return txErr
		}
		for _, t := range transfers {
			fmt.Printf("from=%v, to=%v, amount=%v \n", t.from, t.to, t.amount)
			txErr = uc.addDebtTX(ctx, tx, AddDebtRequest{
				CollectorID: t.to,
				DebtorID:    t.from,
				ChatID:      chatID,
				Amount:      t.amount,
			}, false)
			if txErr != nil {
				return txErr
			}
		}
		return nil
	})
	if err != nil {
		return 0, 0, err
	}

	return len(debts) / 2, len(transfers), nil
}
