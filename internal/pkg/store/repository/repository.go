package repository

import (
	"context"

	"github.com/art22m/dengovie/internal/pkg/models"
)

type UsersRepo interface {
	Create(ctx context.Context, user *models.User) error
	Update(ctx context.Context, user *models.User) (bool, error)
	Delete(ctx context.Context, id int64) (bool, error)
	GetByTelegramUserID(ctx context.Context, id string) (*models.User, error)
}

type ChatsRepo interface {
	Create(ctx context.Context, chat *models.Chat) error
	Update(ctx context.Context, chat *models.Chat) (bool, error)
	Delete(ctx context.Context, id int64) (bool, error)
	GetByTelegramChatID(ctx context.Context, id string) (*models.Chat, error)
}

type DebtsRepo interface {
	Create(ctx context.Context, chat *models.Debt) error
	Read(ctx context.Context) ([]*models.Debt, error)
	Update(ctx context.Context, chat *models.Debt) (bool, error)
	Delete(ctx context.Context, id int64) (bool, error)
}

type EventsRepo interface {
	Create(ctx context.Context, chat *models.Event) error
	List(ctx context.Context) ([]*models.Event, error)
	ReadByCollectorID(ctx context.Context, id int64) ([]*models.Event, error)
	ReadByDebtorID(ctx context.Context, id int64) ([]*models.Event, error)
	ReadByChatID(ctx context.Context, id int64) ([]*models.Event, error)
}
