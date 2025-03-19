package domain

import (
	"errors"
	"fmt"
)

// Error base de la capa Domain
var ErrDomain = errors.New("[DOMAIN]")

// Errores de entidad específica dentro del dominio
var (
	ErrMailAddress = fmt.Errorf("%w [mailAddress]", ErrDomain)
	ErrResult      = fmt.Errorf("%w [result]", ErrDomain)
)

// Errores de validación comunes en el dominio
var (
	ErrInvalidID     = errors.New("invalid ID")
	ErrInvalidEmail  = errors.New("invalid email format")
	ErrInvalidStatus = errors.New("the status is not valid")
)

// Errores compuestos para MailAddress
var (
	ErrMailAddressInvalidID    = fmt.Errorf("%w: %s", ErrInvalidID, ErrMailAddress)
	ErrMailAddressInvalidEmail = fmt.Errorf("%w: %s", ErrInvalidEmail, ErrMailAddress)
)

// Errores compuestos para Result
var (
	ErrResultInvalidID     = fmt.Errorf("%w: %s", ErrInvalidID, ErrResult)
	ErrResultInvalidStatus = fmt.Errorf("%w: %s", ErrInvalidStatus, ErrResult)
)
