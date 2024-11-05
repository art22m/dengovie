package repository

import (
	"context"

	"github.com/art22m/dengovie/internal/generated/dengovie/dengovie/public/table"
	"github.com/art22m/dengovie/internal/pkg/models"
	"github.com/art22m/dengovie/internal/pkg/store"
	"github.com/go-jet/jet/v2/postgres"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
)

type UsersRepository struct {
	db store.Database
}

func (r *UsersRepository) Create(ctx context.Context, user *models.User) error {
	stmt, args := table.Users.INSERT().MODEL(user).Sql()

	_, err := r.db.Exec(ctx, stmt, args...)
	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return nil
	default:
		return err
	}
}

func (r *UsersRepository) Update(ctx context.Context, user *models.User) (bool, error) {
	stmt, args := table.Users.
		UPDATE(
			table.Users.PhoneNumber,
		).
		SET(
			user.PhoneNumber,
		).
		Sql()

	if _, err := r.db.Exec(ctx, stmt, args...); err != nil {
		return false, err
	}

	return true, nil
}

func (r *UsersRepository) Delete(ctx context.Context, id int64) (bool, error) {
	stmt, args := table.Users.DELETE().WHERE(table.Users.UserID.EQ(postgres.Int64(id))).Sql()

	if _, err := r.db.Exec(ctx, stmt, args...); err != nil {
		return false, err
	}

	return true, nil
}

func (r *UsersRepository) GetByTelegramUserID(ctx context.Context, id string) (*models.User, error) {
	user := &models.User{}

	stmt, args := table.Users.
		SELECT(table.Users.AllColumns).
		WHERE(table.Users.TgUserID.EQ(postgres.StringColumn(id))).
		Sql()

	if err := r.db.Get(ctx, user, stmt, args...); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return user, nil
}
