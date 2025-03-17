package domain

import "errors"

var (
	ErrInvalidID     = errors.New("invalid ID")
	ErrInvalidEmail  = errors.New("invalid email format")
	ErrNotFound      = errors.New("resource not found")
	ErrConflict      = errors.New("conflict occurred while processing request")
	ErrValidation    = errors.New("validation failed")
	ErrInvalidStatus = errors.New("The status is not valid")
)

func ErrMailAddress(err error) error {
	return errors.New("[ERROR] [DOMAIN] [mailAddress]" + err.Error())
}

func ErrResult(err error) error {
	return errors.New("[ERROR] [DOMAIN] [mailAddress]" + err.Error())
}
