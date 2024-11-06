package repository

import (
	"context"

	"github.com/jackc/pgx/v5"

	"github.com/art22m/dengovie/internal/pkg/models"
)

type UsersRepo interface {
	Create(ctx context.Context, user *models.User) error
	Delete(ctx context.Context, id int64) (bool, error)
	GetByTelegramUserID(ctx context.Context, id string) (*models.User, error)
	Get(ctx context.Context, id int64) (*models.User, error)
}

type ChatsRepo interface {
	Create(ctx context.Context, chat *models.Chat) error
	Delete(ctx context.Context, id int64) (bool, error)
	GetByTelegramChatID(ctx context.Context, id string) (*models.Chat, error)
}

type DebtsRepo interface {
	List(ctx context.Context, collectorID *int64, chatID *int64) ([]*models.Debt, error)
	Delete(ctx context.Context, collectorID, debtorID, chatID int64) (bool, error)

	CreateTX(ctx context.Context, tx pgx.Tx, debt *models.Debt) error
	GetTX(ctx context.Context, tx pgx.Tx, collectorID, debtorID, chatID int64) (*models.Debt, error)
	UpdateTX(ctx context.Context, tx pgx.Tx, debt *models.Debt) (bool, error)
}

type EventsRepo interface {
	Create(ctx context.Context, event *models.Event) error
	CreateTX(ctx context.Context, tx pgx.Tx, event *models.Event) error
	List(ctx context.Context) ([]*models.Event, error)
	ListByCollectorID(ctx context.Context, id int64) ([]*models.Event, error)
	ListByDebtorID(ctx context.Context, id int64) ([]*models.Event, error)
	ListByChatID(ctx context.Context, id int64) ([]*models.Event, error)
}
