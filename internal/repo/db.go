package repo

import (
	"context"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type transactionKey string

const (
	txKey transactionKey = `tx`
)

type DB interface {
	Get(ctx context.Context) conn
}

type db struct {
	pool *pgxpool.Pool
}

func (r *db) Get(ctx context.Context) conn {
	var q conn

	tx, ok := ctx.Value(txKey).(pgx.Tx)
	if ok && tx != nil {
		q = tx
	} else {
		q = r.pool
	}

	return q
}

type conn interface {
	Begin(ctx context.Context) (pgx.Tx, error)
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
}
