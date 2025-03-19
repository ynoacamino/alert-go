package server

import (
	"context"

	"github.com/google/uuid"
	genMailAddress "github.com/ynoacamino/alert-go/gen/mail_addresses"
	"github.com/ynoacamino/alert-go/internal/core/services"
)

type MailAddressEndPointService struct {
	mailAddresses []*genMailAddress.Mail
	repo          *services.MailAddressService
}

func NewMailAddressEndPointService(repo *services.MailAddressService) genMailAddress.Service {
	return &MailAddressEndPointService{
		repo: repo,
	}
}

func (m *MailAddressEndPointService) ListMailAddresses(
	ctx context.Context,
	payload *genMailAddress.ListMailAddressesPayload,
) (*genMailAddress.Maillist, error) {

	mailAddresses, err := m.repo.ListMailAddresses(ctx, int64(payload.Limit), int64(payload.Page))
	if err != nil {
		return nil, err
	}

	var res []*genMailAddress.Mail = make([]*genMailAddress.Mail, len(mailAddresses))

	for i, mailAddress := range mailAddresses {
		res[i] = &genMailAddress.Mail{
			ID:        mailAddress.ID,
			Address:   mailAddress.Address,
			Active:    mailAddress.Active,
			CreatedAt: mailAddress.CreatedAt.String(),
			UpdateAt:  mailAddress.UpdatedAt.String(),
		}
	}

	return &genMailAddress.Maillist{
		Data:  res,
		Page:  payload.Page,
		Limit: payload.Limit,
		Total: len(res),
	}, nil
}

func (m *MailAddressEndPointService) GetMailAddresses(ctx context.Context, payload *genMailAddress.GetMailAddressesPayload) (*genMailAddress.Mail, error) {
	mailAddress, err := m.repo.GetMailAddress(ctx, payload.ID)
	if err != nil {
		return nil, err
	}

	return &genMailAddress.Mail{
		ID:        mailAddress.ID,
		Address:   mailAddress.Address,
		CreatedAt: mailAddress.CreatedAt.String(),
		UpdateAt:  mailAddress.UpdatedAt.String(),
	}, nil
}

func (m *MailAddressEndPointService) CreateMailAddress(ctx context.Context, payload *genMailAddress.MailPayload) (*genMailAddress.Mail, error) {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	mailAddress, err := m.repo.CreateMailAddress(ctx, uuid.String(), payload.Address, payload.Active)
	if err != nil {
		return nil, err
	}

	return &genMailAddress.Mail{
		ID:        mailAddress.ID,
		Address:   mailAddress.Address,
		Active:    mailAddress.Active,
		CreatedAt: mailAddress.CreatedAt.String(),
		UpdateAt:  mailAddress.UpdatedAt.String(),
	}, nil
}

func (m *MailAddressEndPointService) UpdateMailAddress(ctx context.Context, payload *genMailAddress.UpdateMailAddressPayload) error {
	err := m.repo.UpdateMailAddress(ctx, payload.ID, payload.Address, payload.Active)
	if err != nil {
		return err
	}

	return nil
}

func (m *MailAddressEndPointService) DeleteMailAddress(ctx context.Context, payload *genMailAddress.DeleteMailAddressPayload) error {
	err := m.repo.DeleteMailAddress(ctx, payload.ID)
	if err != nil {
		return err
	}

	return nil
}
