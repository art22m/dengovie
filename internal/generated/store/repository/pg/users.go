package pg

//
//import (
//	"context"
//	"database/sql"
//
//	"github.com/art22m/dengovie/internal/generated/dengovie/dengovie/public/model"
//	"github.com/art22m/dengovie/internal/generated/dengovie/dengovie/public/table"
//	"github.com/art22m/dengovie/internal/pkg/models"
//	"github.com/art22m/dengovie/internal/pkg/store"
//)
//
//type UsersRepo struct {
//	db store.DatabaseOperations
//}
//
//func NewUsers(db store.DatabaseOperations) *UsersRepo {
//	return &UsersRepo{
//		db: db,
//	}
//}
//
//func (r *UsersRepo) Create(ctx context.Context, user model.Users) error {
//	stmt, args := table.Users.
//		INSERT(table.Users.AllColumns.Except(table.Users.UserID)).
//		MODEL(user).
//		ON_CONFLICT().DO_NOTHING().
//		Sql()
//
//	_, err := r.db.Exec(ctx, stmt, args...)
//	return err
//}
//
//func (r *UsersRepo) Delete(ctx context.Context, id int64) (bool, error) {
//	result, err := r.db.Exec(
//		ctx, "DELETE FROM users WHERE user_id = $1", id,
//	)
//	return result.RowsAffected() > 0, err
//}
//
//func (r *UsersRepo) GetByTelegramUserID(ctx context.Context, id string) (*models.User, error) {
//	q := "SELECT user_id, tg_user_id, phone_number, tg_alias, created_at FROM users WHERE user_id = $1"
//	var u models.User
//	err := r.db.Get(ctx, &u, q, id)
//	if err == sql.ErrNoRows {
//		return nil, store.UserNotFound
//	}
//	return &u, err
//}
