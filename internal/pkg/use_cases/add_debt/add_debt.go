package add_debt

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type AddDebtRepo interface{}

type UseCase struct {
	db *pgxpool.Pool
}
