package usecase

//
//import (
//	"context"
//
//	"github.com/pkg/errors"
//	"github.com/samber/lo"
//
//	"github.com/art22m/dengovie/internal/generated/dengovie/dengovie/public/model"
//)
//
//type RegisterUserRequest struct {
//	TelegramUserID string
//	PhoneNumber    string
//	TelegramAlias  *string
//}
//
//func (uc *UseCase) RegisterUser(ctx context.Context, req RegisterUserRequest) error {
//	user := model.Users{
//		TgUserID:    req.TelegramUserID,
//		PhoneNumber: req.PhoneNumber,
//		TgAlias:     lo.FromPtr(req.TelegramAlias),
//	}
//
//	if err := uc.userRepo.Create(ctx, user); err != nil {
//		return errors.Wrap(err, "failed to register user")
//	}
//
//	return nil
//}
