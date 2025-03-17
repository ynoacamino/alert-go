package domain

import (
	"time"
)

type Result struct {
	ID        string
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

var ValidStatus = map[string]bool{
	"AVALIBLE":     true,
	"NOT_AVALIBLE": true,
	"TIMEOUT":      true,
}

func NewResult(id, status string) (*Result, error) {
	if id == "" {
		return nil, ErrResult(ErrInvalidID)
	}

	if !ValidStatus[status] {
		return nil, ErrResult(ErrInvalidStatus)
	}

	return &Result{
		ID:     id,
		Status: status,
	}, nil
}

func (r *Result) UpdateStatus(status string) error {
	if !ValidStatus[status] {
		return ErrResult(ErrInvalidStatus)
	}

	r.Status = status

	return nil
}
