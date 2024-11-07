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
}

func NewUseCase(
	db *pgxpool.Pool,
	debtsRepo repository.DebtsRepo,
	eventsRepo repository.EventsRepo,
	usersRepo repository.UsersRepo,
) *UseCase {
	return &UseCase{
		db:         db,
		debtsRepo:  debtsRepo,
		eventsRepo: eventsRepo,
		usersRepo:  usersRepo,
	}
}
