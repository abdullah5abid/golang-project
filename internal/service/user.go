package service

import (
	"context"

	"github.com/spongeling/admin-api/internal/dao"
)

// GetAllUsers is a service for getting all users
func (s *Service) GetAllUsers(ctx context.Context) ([]*dao.User, error) {
	return s.repo.GetAllUsers(ctx)
}
