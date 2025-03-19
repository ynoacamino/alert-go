package db

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/mattn/go-sqlite3"
)

var ErrDB = errors.New("[DB]")

var (
	ErrResultRepository      = fmt.Errorf("%w [result repository]", ErrDB)
	ErrMailAddressRepository = fmt.Errorf("%w [mailAddress repository]", ErrDB)
)

var (
	ErrConnection = errors.New("connection error")
	ErrNotFound   = errors.New("not found")
	ErrUnique     = errors.New("unique constraint violation")
)

var (
	ErrResultConnection = fmt.Errorf("%w: %s", ErrConnection, ErrResultRepository)
	ErrResultNotFound   = fmt.Errorf("%w: %s", ErrNotFound, ErrResultRepository)
	ErrResultUnique     = fmt.Errorf("%w: %s", ErrUnique, ErrResultRepository)
)

var (
	ErrMailAddressConnection = fmt.Errorf("%w: %s", ErrConnection, ErrMailAddressRepository)
	ErrMailAddressNotFound   = fmt.Errorf("%w: %s", ErrNotFound, ErrMailAddressRepository)
	ErrMailAddressUnique     = fmt.Errorf("%w: %s", ErrUnique, ErrMailAddressRepository)
)

func mapSQLError(err error) error {
	if errors.Is(err, sql.ErrNoRows) {
		return ErrNotFound
	}

	var sqliteErr sqlite3.Error
	if errors.As(err, &sqliteErr) {
		switch sqliteErr.ExtendedCode {
		case sqlite3.ErrConstraintUnique:
			return ErrUnique
		}
	}

	return err
}

func mapSQLErrorResult(err error) error {
	if errors.Is(err, sql.ErrNoRows) {
		return ErrResultNotFound
	}

	if errors.Is(err, ErrUnique) {
		return ErrResultUnique
	}

	return err
}

func mapSQLErrorMailAddress(err error) error {
	if errors.Is(err, sql.ErrNoRows) {
		return ErrMailAddressNotFound
	}

	if errors.Is(err, ErrUnique) {
		return ErrMailAddressUnique
	}

	return err
}
