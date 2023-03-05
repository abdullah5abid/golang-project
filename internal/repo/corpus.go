package repo

import (
	"context"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/spongeling/admin-api/internal/dao"
)

func (r *Repo) GetIdOfPos(ctx context.Context, pos *dao.POS) (uint64, error) {
	var id uint64
	err := pgxscan.Get(ctx, r.Conn.Get(ctx), &id, "SELECT id FROM pos WHERE "+pos.ToPOS().GetConditions())
	return id, wrap(err)
}

func (r *Repo) GetWordsByPosId(ctx context.Context, posId uint64) ([]*dao.Word, error) {
	var words []*dao.Word
	err := pgxscan.Select(ctx, r.Conn.Get(ctx), &words, `
	SELECT w.* FROM word w
	JOIN word_pos wp ON w.id = wp.word_id
	WHERE wp.pos_id = $1`, posId,
	)
	return words, wrap(err)
}

func (r *Repo) GetIdOfWord(ctx context.Context, word string) (uint64, error) {
	var id uint64
	err := pgxscan.Get(ctx, r.Conn.Get(ctx), &id, `SELECT id FROM word WHERE word = $1`, word)
	return id, wrap(err)
}

func (r *Repo) GetPosByWordId(ctx context.Context, wordId uint64) ([]*dao.POS, error) {
	var pos []*dao.POS
	err := pgxscan.Select(ctx, r.Conn.Get(ctx), &pos, `
	SELECT p.* FROM pos p
	JOIN word_pos wp ON p.id = wp.pos_id
	WHERE wp.word_id = $1 
	`, wordId,
	)
	return pos, wrap(err)
}
