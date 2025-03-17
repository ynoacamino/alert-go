package ports

import (
	"context"

	"github.com/ynoacamino/alert-go/internal/core/domain"
)

type ResultRepository interface {
	Save(ctx context.Context, result *domain.Result) (*domain.Result, error)
	FindByID(ctx context.Context, id string) (*domain.Result, error)
	List(ctx context.Context, limit, offset int64) ([]*domain.Result, error)
	Update(ctx context.Context, result *domain.Result) error
	Delete(ctx context.Context, id string) error
}
