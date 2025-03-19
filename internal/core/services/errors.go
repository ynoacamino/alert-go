package services

import (
	"errors"
	"fmt"
)

var (
	ErrRepository = errors.New("[Repository]")
	ErrDomain     = errors.New("[Domain]")
)

var (
	ErrResultRepository      = fmt.Errorf("%w [ResultRepository]", ErrRepository)
	ErrResultDoamin          = fmt.Errorf("%w [ResultDomain]", ErrDomain)
	ErrMailAddressRepository = fmt.Errorf("%w [MailAddressRepository]", ErrRepository)
	ErrMailAddressDomain     = fmt.Errorf("%w [MailAddressDomain]", ErrDomain)
)
