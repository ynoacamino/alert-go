package db

import "errors"

var ErrDB = errors.New("[ERROR] [DB] ")

var ErrResultRepository = errors.New(ErrDB.Error() + "[result repository] ")

var ErrMailAddressRepository = errors.New(ErrDB.Error() + "[mailAddress repository] ")

func ErrSaveResult(err error) error {
	return errors.New(ErrResultRepository.Error() + "[save result failed]: " + err.Error())
}

func ErrUpdateResult(err error) error {
	return errors.New(ErrResultRepository.Error() + "[update result failed]: " + err.Error())
}

func ErrDeleteResult(err error) error {
	return errors.New(ErrResultRepository.Error() + "[delete result failed]: " + err.Error())
}

func ErrFindByIDResult(err error) error {
	return errors.New(ErrResultRepository.Error() + "[find by id result failed]: " + err.Error())
}

func ErrListResults(err error) error {
	return errors.New(ErrResultRepository.Error() + "[list results failed]: " + err.Error())
}

func ErrSaveMailAddress(err error) error {
	return errors.New(ErrMailAddressRepository.Error() + "[save mail address failed]: " + err.Error())
}

func ErrUpdateMailAddress(err error) error {
	return errors.New(ErrMailAddressRepository.Error() + "[update mail address failed]: " + err.Error())
}

func ErrDeleteMailAddress(err error) error {
	return errors.New(ErrMailAddressRepository.Error() + "[delete mail address failed]: " + err.Error())
}

func ErrFindByIDMailAddress(err error) error {
	return errors.New(ErrMailAddressRepository.Error() + "[find by id mail address failed]: " + err.Error())
}

func ErrListMailAddresses(err error) error {
	return errors.New(ErrMailAddressRepository.Error() + "[list mail addresses failed]: " + err.Error())
}
