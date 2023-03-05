package repo

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/doug-martin/goqu/v9"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/spongeling/admin-api/internal/dao"
)

const (
	DefaultPageSize = 25
)

type Repo struct {
	Conn DB
}

func New(conn *pgxpool.Pool) *Repo {
	return &Repo{
		Conn: &db{
			pool: conn,
		},
	}
}

func (r *Repo) GetByID(ctx context.Context, m dao.Model, id uint64) error {
	return wrap(pgxscan.Get(ctx, r.Conn.Get(ctx), m, fmt.Sprintf("SELECT * FROM %s WHERE id = $1", m.GetTable()), id))
}

func (r *Repo) DeleteByID(ctx context.Context, m dao.Model, id uint64) error {
	return wrap(pgxscan.Get(ctx, r.Conn.Get(ctx), m, fmt.Sprintf("DELETE FROM %s WHERE id = $1 RETURNING *", m.GetTable()), id))
}

func (r *Repo) Create(ctx context.Context, m dao.Model) error {
	if err := m.Validate(); err != nil {
		return err
	}

	q, params, err := goqu.Insert(m.GetTable()).Rows(m).ToSQL()
	if err != nil {
		return err
	}

	_, err = r.Conn.Get(ctx).Exec(ctx, q, params...)
	return wrap(err)
}

func (r *Repo) Update(ctx context.Context, m dao.Model) error {
	if err := m.Validate(); err != nil {
		return err
	}

	q, params, err := goqu.Update(m.GetTable()).Where(goqu.Ex{
		"id": m.GetID(),
	}).Set(m).ToSQL()
	if err != nil {
		return err
	}
	_, err = r.Conn.Get(ctx).Exec(ctx, q, params...)
	return wrap(err)
}

func wrap(err error) error {
	if err == nil {
		return nil
	}
	if errors.Is(err, pgx.ErrNoRows) || pgxscan.NotFound(err) {
		return err
	}

	return err
}

func wrapf(format string, a ...interface{}) error {
	return wrap(fmt.Errorf(format, a...))
}

func (r *Repo) RunTx(ctx context.Context, f func(ctx context.Context) error) error {
	tx, err := r.Conn.Get(ctx).Begin(ctx)
	if err != nil {
		return wrapf("error calling begin transaction: %w", err)
	}
	defer func() {
		err := tx.Rollback(ctx)
		if err != nil {
			log.Printf("error rolling back transaction: %v", err)
		}
	}()
	err = f(context.WithValue(ctx, txKey, tx))
	if err != nil {
		return wrapf("error running transaction: %w", err)
	}
	err = tx.Commit(ctx)
	if err != nil {
		return wrapf("error committing transaction: %w", err)
	}
	return nil
}
