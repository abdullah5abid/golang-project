package request

import "github.com/spongeling/admin-api/internal/errors"

type GetPosOfWord struct {
	Word string `json:"word"`
}

func (w *GetPosOfWord) Validate() error {
	if w.Word == "" {
		return errors.BadRequest("word must not be empty")
	}
	return nil
}
