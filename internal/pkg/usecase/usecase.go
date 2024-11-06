package usecase

import (
	"github.com/art22m/dengovie/internal/pkg/store/repository"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UseCase struct {
	db *pgxpool.Pool

	debtsRepo  repository.DebtsRepo
	eventsRepo repository.EventsRepo
	usersRepo  repository.UsersRepo
	chatsRepo  repository.ChatsRepo
}

func NewUseCase(
	db *pgxpool.Pool,
	debtsRepo repository.DebtsRepo,
	eventsRepo repository.EventsRepo,
	usersRepo repository.UsersRepo,
	chatsRepo repository.ChatsRepo,
) *UseCase {
	return &UseCase{
		db:         db,
		debtsRepo:  debtsRepo,
		eventsRepo: eventsRepo,
		usersRepo:  usersRepo,
		chatsRepo:  chatsRepo,
	}
}
