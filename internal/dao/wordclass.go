package dao

import "github.com/spongeling/admin-api/internal/errors"

type WordClass struct {
	WordId  uint64 `db:"word_id"`
	ClassId uint   `db:"class_id"`
}

func (*WordClass) GetID() uint64 {
	return 0
}

func (*WordClass) GetTable() string {
	return "word_class"
}

func (wc *WordClass) Validate() error {
	if wc.WordId == 0 {
		return errors.BadRequest("word id is required")
	} else if wc.ClassId == 0 {
		return errors.BadRequest("class id is required")
	}
	return nil
}
