package services

import (
	"context"

	"github.com/ynoacamino/alert-go/internal/core/domain"
	"github.com/ynoacamino/alert-go/internal/core/ports"
)

type ResultService struct {
	repo ports.ResultRepository
}

func NewResultService(repo ports.ResultRepository) *ResultService {
	return &ResultService{
		repo: repo,
	}
}

func (s *ResultService) CreateResult(ctx context.Context, id, status string) (*domain.Result, error) {
	result, err := domain.NewResult(id, status)
	if err != nil {
		return nil, err
	}

	row, err := s.repo.Save(ctx, result)
	if err != nil {
		return nil, err
	}

	return row, nil
}

func (s *ResultService) GetResult(ctx context.Context, id string) (*domain.Result, error) {
	result, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *ResultService) ListResult(ctx context.Context, limit, page int64) ([]*domain.Result, error) {
	results, err := s.repo.List(ctx, limit, page)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (s *ResultService) UpdateResult(ctx context.Context, id, status string) (*domain.Result, error) {
	result, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	result.Status = status
	if err := s.repo.Update(ctx, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (s *ResultService) DeleteResult(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
