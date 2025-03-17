package services

import (
	"context"

	"github.com/ynoacamino/alert-go/internal/core/domain"
	"github.com/ynoacamino/alert-go/internal/core/ports"
)

type MailAddressService struct {
	repo ports.MailAddressRepository
}

func NewMailAddressService(repo ports.MailAddressRepository) *MailAddressService {
	return &MailAddressService{
		repo: repo,
	}
}

func (s *MailAddressService) CreateMailAddress(ctx context.Context, id, address string, active bool) (*domain.MailAddress, error) {
	mail, err := domain.NewMailAddress(id, address, active)
	if err != nil {
		return nil, err
	}

	row, err := s.repo.Save(ctx, mail)
	if err != nil {
		return nil, err
	}

	return row, nil
}

func (s *MailAddressService) GetMailAddress(ctx context.Context, id string) (*domain.MailAddress, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *MailAddressService) ListMailAddresses(ctx context.Context, limit, offset int64) ([]*domain.MailAddress, error) {
	return s.repo.List(ctx, limit, offset)
}

func (s *MailAddressService) UpdateMailAddress(ctx context.Context, id, address string, active bool) (*domain.MailAddress, error) {
	mail, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	mail.Address = address
	mail.Active = active

	err = s.repo.Update(ctx, mail)
	if err != nil {
		return nil, err
	}

	return mail, nil
}

func (s *MailAddressService) DeleteMailAddress(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
