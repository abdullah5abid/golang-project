package dao

import "github.com/spongeling/admin-api/internal/errors"

type Class struct {
	Id          uint64   `db:"id"`
	Name        string   `db:"name"`
	WordId      *uint64  `db:"word_id"`
	Description string   `db:"description"`
	Words       []string `db:"words"`
}

func (c *Class) GetID() uint64 {
	return c.Id
}

func (*Class) GetTable() string {
	return "class"
}

func (c *Class) Validate() error {
	if c.Name == "" {
		return errors.BadRequest("name is required")
	}
	return nil
}
