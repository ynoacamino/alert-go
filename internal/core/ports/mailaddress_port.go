package ports

import (
	"context"

	"github.com/ynoacamino/alert-go/internal/core/domain"
)

type MailAddressRepository interface {
	Save(ctx context.Context, mail *domain.MailAddress) (*domain.MailAddress, error)
	FindByID(ctx context.Context, id string) (*domain.MailAddress, error)
	List(ctx context.Context, limit, offset int64) ([]*domain.MailAddress, error)
	Update(ctx context.Context, mail *domain.MailAddress) error
	Delete(ctx context.Context, id string) error
}
