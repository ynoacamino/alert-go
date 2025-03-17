package domain

import (
	"net/mail"
	"time"
)

type MailAddress struct {
	ID        string
	Address   string
	Active    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewMailAddress(id, address string, active bool) (*MailAddress, error) {
	if id == "" {
		return nil, ErrMailAddress(ErrInvalidID)
	}

	return &MailAddress{
		ID:      id,
		Address: address,
		Active:  active,
	}, nil
}

func (m *MailAddress) Activate() {
	m.Active = true
}

func (m *MailAddress) Deactivate() {
	m.Active = false
}

func (m *MailAddress) UpdateAddress(address string) error {
	err := validateEmail(address)

	if err != nil {
		return err
	}

	m.Address = address

	return nil
}

func validateEmail(email string) error {
	_, err := mail.ParseAddress(email)

	if err != nil {
		return ErrMailAddress(ErrInvalidEmail)
	}

	return nil
}
