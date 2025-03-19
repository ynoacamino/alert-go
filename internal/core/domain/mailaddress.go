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

func ValidateMailAddressID(id string) error {
	if id == "" {
		return ErrResultInvalidID
	}
	return nil
}

func ValidateEmail(email string) error {
	_, err := mail.ParseAddress(email)

	if err != nil {
		return ErrMailAddressInvalidEmail
	}

	return nil
}

func NewMailAddress(id, address string, active bool) (*MailAddress, error) {
	err := ValidateMailAddressID(id)
	if err != nil {
		return nil, err
	}

	err = ValidateEmail(address)
	if err != nil {
		return nil, err
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
	err := ValidateEmail(address)
	if err != nil {
		return err
	}

	m.Address = address

	return nil
}
