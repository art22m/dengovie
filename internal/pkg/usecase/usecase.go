package usecase

import (
	"context"

	"github.com/art22m/dengovie/internal/generated/dengovie/dengovie/public/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DebtRepo interface {
	GetForUpdate(ctx context.Context, tx pgx.Tx, collectorID, debtorID, chatID int64) (*model.Debts, error)
	Create(ctx context.Context, tx pgx.Tx, debt model.Debts) error
	Update(ctx context.Context, tx pgx.Tx, debt model.Debts) error
}

type EventRepo interface {
	Create(ctx context.Context, tx pgx.Tx, event model.Events) error
}

type UserRepo interface {
	Create(ctx context.Context, user model.Users) error
}

type UseCase struct {
	db        *pgxpool.Pool
	debtRepo  DebtRepo
	eventRepo EventRepo
	userRepo  UserRepo
}

func NewUseCase(db *pgxpool.Pool, debtRepo DebtRepo, eventRepo EventRepo, userRepo UserRepo) *UseCase {
	return &UseCase{db: db, debtRepo: debtRepo, eventRepo: eventRepo, userRepo: userRepo}
}
