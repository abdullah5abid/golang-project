package service

import (
	"context"

	"github.com/spongeling/admin-api/internal/dao"
)

func (s *Service) GetWordsByPos(ctx context.Context, pos *dao.POS) ([]*dao.Word, error) {
	id, err := s.repo.GetIdOfPos(ctx, pos)
	if err != nil {
		return nil, err
	}

	words, err := s.repo.GetWordsByPosId(ctx, id)
	if err != nil {
		return nil, err
	}

	return words, nil
}

func (s *Service) GetPosOfWord(ctx context.Context, word string) ([]*dao.POS, error) {
	id, err := s.repo.GetIdOfWord(ctx, word)
	if err != nil {
		return nil, err
	}

	return s.repo.GetPosByWordId(ctx, id)
}
