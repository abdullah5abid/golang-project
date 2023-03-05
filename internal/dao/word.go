package dao

import "github.com/spongeling/admin-api/internal/errors"

type Word struct {
	ID   uint64 `db:"id"`
	Word string `db:"word"`
}

func (w *Word) GetID() uint64 {
	return w.ID
}

func (*Word) GetTable() string {
	return "word"
}

func (w *Word) Validate() error {
	if w.Word == "" {
		return errors.BadRequest("word is required")
	}

	return nil
}
