package repo

import (
	"context"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/spongeling/admin-api/internal/dao"
)

// GetAllTopLevelCategories is a procedure for fetching all top level categories
func (r *Repo) GetAllTopLevelCategories(ctx context.Context) ([]*dao.Category, error) {
	var categories []*dao.Category
	err := pgxscan.Select(ctx, r.Conn.Get(ctx), &categories, `SELECT * FROM category WHERE parent_id is NULL`)
	return categories, wrap(err)
}

// GetSubCategories is a procedure for fetching all subcategories for a category
func (r *Repo) GetSubCategories(ctx context.Context, categoryId uint) ([]*dao.Category, error) {
	var categories []*dao.Category
	err := pgxscan.Select(ctx, r.Conn.Get(ctx), &categories, `SELECT * FROM category WHERE parent_id=$1`, categoryId)
	return categories, wrap(err)
}

// GetCategory is a procedure for fetching a category
func (r *Repo) GetCategory(ctx context.Context, categoryId uint) (dao.Category, error) {
	var category dao.Category
	err := pgxscan.Get(ctx, r.Conn.Get(ctx), &category, `SELECT * FROM category WHERE id=$1`, categoryId)
	return category, wrap(err)
}
