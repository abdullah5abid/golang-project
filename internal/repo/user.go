package repo

import (
	"context"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/spongeling/admin-api/internal/dao"
)

// GetAllUsers is a function for getting all users from db
func (r *Repo) GetAllUsers(ctx context.Context) ([]*dao.User, error) {
	var users []*dao.User
	err := pgxscan.Select(ctx, r.Conn.Get(ctx), &users, `SELECT * FROM "user"`)
	return users, wrap(err)
}
