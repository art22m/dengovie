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

type ChatRepository struct {
	db store.Database
}

func (r *ChatRepository) Create(ctx context.Context, chat *models.Chat) error {
	stmt, args := table.Chats.INSERT().MODEL(chat).Sql()

	_, err := r.db.Exec(ctx, stmt, args...)
	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return nil
	default:
		return err
	}
}

func (r *ChatRepository) Update(ctx context.Context, chat *models.Chat) (bool, error) {
	stmt, args := table.Chats.
		UPDATE(
			table.Chats.Description,
		).
		SET(chat.Description).
		Sql()

	if _, err := r.db.Exec(ctx, stmt, args...); err != nil {
		return false, err
	}

	return true, nil
}

func (r *ChatRepository) Delete(ctx context.Context, id int64) (bool, error) {
	stmt, args := table.Chats.DELETE().WHERE(table.Chats.ChatID.EQ(postgres.Int64(id))).Sql()

	if _, err := r.db.Exec(ctx, stmt, args...); err != nil {
		return false, err
	}

	return true, nil
}

func (r *ChatRepository) GetByTelegramChatID(ctx context.Context, id string) (*models.Chat, error) {
	chat := &models.Chat{}

	stmt, args := table.Chats.
		SELECT(table.Chats.AllColumns).
		WHERE(table.Chats.TgChatID.EQ(postgres.String(id))).
		Sql()

	if err := r.db.Get(ctx, chat, stmt, args...); err != nil {
		return nil, err
	}

	return chat, nil
}
