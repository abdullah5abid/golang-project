package repo

import (
	"context"
	"fmt"

	"github.com/doug-martin/goqu/v9"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/spongeling/admin-api/internal/dao"
)

// GetAllClasses is a procedure for getting all classes with words
func (r *Repo) GetAllClasses(ctx context.Context) ([]*dao.Class, error) {
	var classes []*dao.Class
	err := pgxscan.Select(ctx, r.Conn.Get(ctx), &classes, `
	SELECT c.name, c.description, array_agg(w.word) AS words
	FROM word_class wc
	JOIN word w ON wc.word_id = w.id
	JOIN "class" c ON wc.class_id = c.id
	GROUP BY c.id
	`)
	return classes, err
}

// AddClass is a procedure for saving a class
func (r *Repo) AddClass(ctx context.Context, m *dao.Class) (uint, error) {
	var id uint
	sql, _, err := goqu.Insert("class").
		Rows(goqu.Record{"name": m.Name, "word_id": m.WordId, "description": m.Description}).
		Returning("id").
		ToSQL()
	if err != nil {
		return 0, err
	}

	err = pgxscan.Get(ctx, r.Conn.Get(ctx), &id, sql)
	return id, err
}

// DeleteClass is a procedure for deleting word class
func (r *Repo) DeleteClass(ctx context.Context, classId uint) (uint, error) {
	var id uint
	sql, _, err := goqu.From("class").
		Where(goqu.C("id").Eq(fmt.Sprintf("%v", classId))).
		Delete().
		Returning("id").
		ToSQL()
	if err != nil {
		return 0, err
	}
	err = pgxscan.Get(ctx, r.Conn.Get(ctx), &id, sql)
	return id, err
}

// AddWordClasses is a procedure for saving word classes
func (r *Repo) AddWordClasses(ctx context.Context, m []dao.WordClass) error {
	q, params, err := goqu.Insert("word_class").Rows(m).ToSQL()
	if err != nil {
		return nil
	}
	_, err = r.Conn.Get(ctx).Exec(ctx, q, params...)
	return err
}
