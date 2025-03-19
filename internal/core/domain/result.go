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

var validStatus = map[string]bool{
	"AVAILABLE":     true,
	"NOT_AVAILABLE": true,
	"TIMEOUT":       true,
}

func ValidateResultID(id string) error {
	if id == "" {
		return ErrResultInvalidID
	}
	return nil
}

func ValidateResultStatus(status string) error {
	if !validStatus[status] {
		return ErrResultInvalidStatus
	}
	return nil
}

func NewResult(id, status string) (*Result, error) {
	err := ValidateResultID(id)
	if err != nil {
		return nil, err
	}

	err = ValidateResultStatus(status)
	if err != nil {
		return nil, err
	}

	return &Result{
		ID:     id,
		Status: status,
	}, nil
}

func (r *Result) UpdateStatus(status string) error {
	err := ValidateResultStatus(status)
	if err != nil {
		return err
	}

	r.Status = status

	return nil
}
