package repository

import (
	"context"

	"github.com/art22m/dengovie/internal/pkg/models"
)

type UsersRepo interface {
	Create(ctx context.Context, user *models.User) error
	Delete(ctx context.Context, id int64) (bool, error)
	GetByTelegramUserID(ctx context.Context, id string) (*models.User, error)
}

type ChatsRepo interface {
	Create(ctx context.Context, chat *models.Chat) error
	Delete(ctx context.Context, id int64) (bool, error)
	GetByTelegramChatID(ctx context.Context, id string) (*models.Chat, error)
}

type DebtsRepo interface {
	Create(ctx context.Context, debt *models.Debt) error
	List(ctx context.Context) ([]*models.Debt, error)
	Update(ctx context.Context, debt *models.Debt) (bool, error)
	Delete(ctx context.Context, id int64) (bool, error)
}

type EventsRepo interface {
	Create(ctx context.Context, event *models.Event) error
	List(ctx context.Context) ([]*models.Event, error)
	ListByCollectorID(ctx context.Context, id int64) ([]*models.Event, error)
	ListByDebtorID(ctx context.Context, id int64) ([]*models.Event, error)
	ListByChatID(ctx context.Context, id int64) ([]*models.Event, error)
}
