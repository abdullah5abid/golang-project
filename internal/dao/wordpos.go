package dao

type WordPos struct {
	WordId  uint64 `db:"word_id"`
	LemmaId uint64 `db:"lemma_id"`
	PosId   uint64 `db:"pos_id"`
}

func (*WordPos) GetID() uint64 {
	return 0
}

func (*WordPos) GetTable() string {
	return "word_pos"
}

func (wp *WordPos) Validate() error {
	return nil
}
