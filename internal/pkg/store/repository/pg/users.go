package pg

import (
	"context"

	"github.com/art22m/dengovie/internal/pkg/models"
	"github.com/art22m/dengovie/internal/pkg/store"
)

type UsersRepo struct {
	db store.DatabaseOperations
}

func NewUsers(db store.DatabaseOperations) *UsersRepo {
	return &UsersRepo{
		db: db,
	}
}

func (r *UsersRepo) Create(ctx context.Context, user *models.User) error {
	q := "INSERT INTO users(tg_user_id, phone_number, tg_alias) VALUES ($1, $2, $3) ON CONFLICT DO NOTHING"
	_, err := r.db.Exec(
		ctx, q,
		user.TelegramUserID, user.PhoneNumber, user.TelegramAlias,
	)
	return err
}

func (r *UsersRepo) Delete(ctx context.Context, id int64) (bool, error) {
	result, err := r.db.Exec(
		ctx, "DELETE FROM users WHERE user_id = $1", id,
	)
	return result.RowsAffected() > 0, err
}

func (r *UsersRepo) GetByTelegramUserID(ctx context.Context, id string) (*models.User, error) {
	q := "SELECT user_id, tg_user_id, phone_number, tg_alias, created_at FROM users WHERE tg_user_id = $1"
	users := make([]*models.User, 0)
	err := r.db.Select(ctx, &users, q, id)
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, store.UserNotFound
	}
	return users[0], err
}

func (r *UsersRepo) Get(ctx context.Context, id int64) (*models.User, error) {
	q := "SELECT user_id, tg_user_id, phone_number, tg_alias, created_at FROM users WHERE user_id = $1"
	users := make([]*models.User, 0)
	err := r.db.Select(ctx, &users, q, id)
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, store.UserNotFound
	}
	return users[0], err
}

func (r *UsersRepo) List(ctx context.Context) ([]*models.User, error) {
	q := "SELECT user_id, tg_user_id, phone_number, tg_alias, created_at FROM users"
	users := make([]*models.User, 0)
	err := r.db.Select(ctx, &users, q)
	if err != nil {
		return nil, err
	}
	return users, err
}
