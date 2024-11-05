package pg

import (
	"context"
	"database/sql"

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
		user.UserID, user.TelegramUserID, user.TelegramAlias,
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
	q := "SELECT user_id, tg_user_id, phone_number, tg_alias, created_at FROM users WHERE user_id = $1"
	var u models.User
	err := r.db.Get(ctx, &u, q, id)
	if err == sql.ErrNoRows {
		return nil, store.UserNotFound
	}
	return &u, err
}
