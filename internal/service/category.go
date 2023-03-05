package service

import (
	"context"

	"github.com/spongeling/admin-api/internal/dao"
)

// GetAllTopLevelCategories retrieves a list of top-level categories from the database
func (s *Service) GetAllTopLevelCategories(ctx context.Context) ([]*dao.Category, error) {
	return s.repo.GetAllTopLevelCategories(ctx)
}

// GetSubCategories retrieves a list of subcategories from the database
func (s *Service) GetSubCategories(ctx context.Context, categoryId uint) ([]*dao.Category, error) {
	c, err := s.repo.GetCategory(ctx, categoryId)
	if err != nil {
		return nil, err
	}
	return s.repo.GetSubCategories(ctx, c.Id)
}
