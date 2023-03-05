package dao

import "github.com/spongeling/admin-api/internal/errors"

type Category struct {
	Id       uint   `db:"id"`
	Name     string `db:"name"`
	ParentId *uint  `db:"parent_id"`
}

func (c *Category) GetID() uint {
	return c.Id
}

func (_ *Category) GetTable() string {
	return "category"
}

func (c *Category) Validate() error {
	if c.Name == "" {
		return errors.BadRequest("category name is required")
	}
	return nil
}
