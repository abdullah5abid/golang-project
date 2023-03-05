package repo

import (
	"context"

	"github.com/doug-martin/goqu/v9"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/spongeling/admin-api/internal/dao"
)

// GetWords is a procedure for fetching words
func (r *Repo) GetWords(ctx context.Context, words []string) ([]*dao.Word, error) {
	var m []*dao.Word
	sql, _, _ := goqu.From("word").
		Where(goqu.Ex{"word": words}).
		ToSQL()
	err := pgxscan.Select(ctx, r.Conn.Get(ctx), &m, sql)
	return m, err
}

// AddWords is a procedure for saving multiple words
func (r *Repo) AddWords(ctx context.Context, words []string) ([]uint64, error) {
	var ids []uint64

	var rows []interface{}
	for _, word := range words {
		rows = append(rows, goqu.Record{"word": word})
	}

	q, _, err := goqu.Insert("word").
		Rows(rows...).
		Returning("id").
		ToSQL()
	if err != nil {
		return nil, err
	}

	err = pgxscan.Select(ctx, r.Conn.Get(ctx), &ids, q)
	return ids, err
}
