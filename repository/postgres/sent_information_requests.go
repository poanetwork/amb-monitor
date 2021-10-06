package postgres

import (
	"amb-monitor/db"
	"amb-monitor/entity"
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

type sentInformationRequestsRepo struct {
	table string
	db    *db.DB
}

func NewSentInformationRequestsRepo(table string, db *db.DB) entity.SentInformationRequestsRepo {
	return &sentInformationRequestsRepo{
		table: table,
		db:    db,
	}
}

func (r *sentInformationRequestsRepo) Ensure(ctx context.Context, msg *entity.SentInformationRequest) error {
	q, args, err := sq.Insert(r.table).
		Columns("log_id", "bridge_id", "message_id").
		Values(msg.LogID, msg.BridgeID, msg.MessageID).
		Suffix("ON CONFLICT (log_id) DO UPDATE SET updated_at = NOW()").
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return fmt.Errorf("can't build query: %w", err)
	}
	_, err = r.db.ExecContext(ctx, q, args...)
	if err != nil {
		return fmt.Errorf("can't insert sent information request: %w", err)
	}
	return nil
}