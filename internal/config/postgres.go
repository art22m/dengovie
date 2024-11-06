package config

import (
	"context"
	"fmt"
	"os"

	"github.com/art22m/dengovie/internal/pkg/store"
	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateDatabase(ctx context.Context) (*store.Database, error) {
	pool, err := pgxpool.New(ctx, generateDsn())
	if err != nil {
		return nil, err
	}
	return store.NewDatabase(pool), nil
}

func generateDsn() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv(DBHost),
		os.Getenv(DBPort),
		os.Getenv(DBUser),
		os.Getenv(DBPass),
		os.Getenv(DBName),
	)
}
