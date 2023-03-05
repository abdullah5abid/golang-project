package service

import (
	"github.com/spongeling/admin-api/internal/repo"
)

type Service struct {
	repo *repo.Repo
}

func New(r *repo.Repo) *Service {
	return &Service{r}
}
