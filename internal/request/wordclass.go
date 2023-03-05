package request

import "github.com/spongeling/admin-api/internal/errors"

type WordClass struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Words       []string `json:"words"`
}

func (wc *WordClass) Validate() error {
	if wc.Name == "" && len(wc.Words) == 0 {
		return errors.BadRequest("name should not be empty")
	}
	return nil
}
