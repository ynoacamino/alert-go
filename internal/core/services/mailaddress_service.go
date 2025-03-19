package services

import (
	"context"
	"fmt"

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
		return nil, fmt.Errorf("%w: %s", ErrMailAddressDomain, err)
	}

	row, err := s.repo.Save(ctx, mail)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrMailAddressRepository, err)
	}

	return row, nil
}

func (s *MailAddressService) GetMailAddress(ctx context.Context, id string) (*domain.MailAddress, error) {
	err := domain.ValidateMailAddressID(id)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrMailAddressDomain, err)
	}

	mail, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrMailAddressRepository, err)
	}

	return mail, nil
}

func (s *MailAddressService) ListMailAddresses(ctx context.Context, limit, offset int64) ([]*domain.MailAddress, error) {
	mails, err := s.repo.List(ctx, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrMailAddressRepository, err)
	}

	return mails, nil
}

func (s *MailAddressService) UpdateMailAddress(ctx context.Context, id, address string, active bool) error {
	err := domain.ValidateMailAddressID(id)
	if err != nil {
		return fmt.Errorf("%w: %s", ErrMailAddressDomain, err)
	}

	err = domain.ValidateEmail(address)
	if err != nil {
		return fmt.Errorf("%w: %s", ErrMailAddressDomain, err)
	}

	mail, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return fmt.Errorf("%w: %s", ErrMailAddressRepository, err)
	}

	mail.Address = address
	mail.Active = active

	err = s.repo.Update(ctx, mail)
	if err != nil {
		return fmt.Errorf("%w: %s", ErrMailAddressRepository, err)
	}

	return nil
}

func (s *MailAddressService) DeleteMailAddress(ctx context.Context, id string) error {
	err := domain.ValidateMailAddressID(id)
	if err != nil {
		return fmt.Errorf("%w: %s", ErrMailAddressDomain, err)
	}

	err = s.repo.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("%w: %s", ErrMailAddressRepository, err)
	}

	return nil
}
