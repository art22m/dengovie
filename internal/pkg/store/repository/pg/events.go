package pg

import (
	"context"

	"github.com/art22m/dengovie/internal/generated/dengovie/dengovie/public/model"
	"github.com/art22m/dengovie/internal/generated/dengovie/dengovie/public/table"
	"github.com/art22m/dengovie/internal/pkg/models"
	"github.com/art22m/dengovie/internal/pkg/store"
	"github.com/jackc/pgx/v5"
)

type EventsRepo struct {
	db store.DatabaseOperations
}

func NewEvents(db store.DatabaseOperations) *EventsRepo {
	return &EventsRepo{
		db: db,
	}
}

func (r *EventsRepo) Create(ctx context.Context, tx pgx.Tx, event model.Events) error {
	stmt, args := table.Events.
		INSERT(
			table.Events.AllColumns.Except(table.Events.EventID),
		).
		MODEL(event).
		ON_CONFLICT().DO_NOTHING().
		Sql()

	_, err := tx.Exec(ctx, stmt, args...)

	return err
}

func (r *EventsRepo) List(ctx context.Context) ([]*models.Event, error) {
	events := make([]*models.Event, 0)
	err := r.db.Select(ctx, &events,
		"SELECT event_id, collector_id, debtor_id, chat_id, amount, description, created_at FROM events",
	)
	return events, err
}

func (r *EventsRepo) ListByCollectorID(ctx context.Context, id int64) ([]*models.Event, error) {
	events := make([]*models.Event, 0)
	err := r.db.Select(ctx, &events,
		"SELECT event_id, collector_id, debtor_id, chat_id, amount, description, created_at FROM events WHERE collector_id = $1",
		id,
	)
	return events, err
}

func (r *EventsRepo) ListByDebtorID(ctx context.Context, id int64) ([]*models.Event, error) {
	events := make([]*models.Event, 0)
	err := r.db.Select(ctx, &events,
		"SELECT event_id, collector_id, debtor_id, chat_id, amount, description, created_at FROM events WHERE debtor_id = $1",
		id,
	)
	return events, err
}

func (r *EventsRepo) ListByChatID(ctx context.Context, id int64) ([]*models.Event, error) {
	events := make([]*models.Event, 0)
	err := r.db.Select(ctx, &events,
		"SELECT event_id, collector_id, debtor_id, chat_id, amount, description, created_at FROM events WHERE chat_id = $1",
		id,
	)
	return events, err
}
