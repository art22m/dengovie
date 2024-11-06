package pg

//
//import (
//	"context"
//	"database/sql"
//
//	"github.com/art22m/dengovie/internal/pkg/models"
//	"github.com/art22m/dengovie/internal/pkg/store"
//)
//
//type ChatsRepo struct {
//	db store.DatabaseOperations
//}
//
//func NewChats(db store.DatabaseOperations) *UsersRepo {
//	return &UsersRepo{
//		db: db,
//	}
//}
//
//func (r *ChatsRepo) Create(ctx context.Context, chat *models.Chat) error {
//	q := "INSERT INTO chats(chat_id, tg_chat_id, description) VALUES ($1, $2, $3) ON CONFLICT DO NOTHING"
//	_, err := r.db.Exec(
//		ctx, q,
//		chat.ChatID, chat.TelegramChatID, chat.Description,
//	)
//	return err
//}
//
//func (r *ChatsRepo) Delete(ctx context.Context, id int64) (bool, error) {
//	result, err := r.db.Exec(
//		ctx, "DELETE FROM chats WHERE chat_id = $1", id,
//	)
//	return result.RowsAffected() > 0, err
//}
//
//func (r *ChatsRepo) GetByTelegramChatID(ctx context.Context, id string) (*models.Chat, error) {
//	q := "SELECT chat_id, tg_chat_id, description, created_at FROM chats WHERE chat_id = $1"
//	var u models.Chat
//	err := r.db.Get(ctx, &u, q, id)
//	if err == sql.ErrNoRows {
//		return nil, store.UserNotFound
//	}
//	return &u, err
//}
