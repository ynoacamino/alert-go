// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: mailAddress.sql

package db

import (
	"context"
)

const createMailAddress = `-- name: CreateMailAddress :one
INSERT INTO
    mailAddress (id, address)
VALUES
    (?, ?) RETURNING id, address, active, createdat, updatedat
`

type CreateMailAddressParams struct {
	ID      string
	Address string
}

func (q *Queries) CreateMailAddress(ctx context.Context, arg CreateMailAddressParams) (MailAddress, error) {
	row := q.db.QueryRowContext(ctx, createMailAddress, arg.ID, arg.Address)
	var i MailAddress
	err := row.Scan(
		&i.ID,
		&i.Address,
		&i.Active,
		&i.Createdat,
		&i.Updatedat,
	)
	return i, err
}

const deleteMailAddress = `-- name: DeleteMailAddress :exec
DELETE FROM mailAddress
WHERE
    id = ?
`

func (q *Queries) DeleteMailAddress(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteMailAddress, id)
	return err
}

const getMailAddress = `-- name: GetMailAddress :one
SELECT
    id, address, active, createdat, updatedat
FROM
    mailAddress
WHERE
    id = ?
`

func (q *Queries) GetMailAddress(ctx context.Context, id string) (MailAddress, error) {
	row := q.db.QueryRowContext(ctx, getMailAddress, id)
	var i MailAddress
	err := row.Scan(
		&i.ID,
		&i.Address,
		&i.Active,
		&i.Createdat,
		&i.Updatedat,
	)
	return i, err
}

const listAllMailAddresses = `-- name: ListAllMailAddresses :many
SELECT
    id, address, active, createdat, updatedat
FROM
    mailAddress
ORDER BY
    updatedAt DESC
`

func (q *Queries) ListAllMailAddresses(ctx context.Context) ([]MailAddress, error) {
	rows, err := q.db.QueryContext(ctx, listAllMailAddresses)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []MailAddress
	for rows.Next() {
		var i MailAddress
		if err := rows.Scan(
			&i.ID,
			&i.Address,
			&i.Active,
			&i.Createdat,
			&i.Updatedat,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listMailAddresses = `-- name: ListMailAddresses :many
SELECT
    id, address, active, createdat, updatedat
FROM
    mailAddress
ORDER BY
    updatedAt DESC
LIMIT
    ?
OFFSET
    ?
`

type ListMailAddressesParams struct {
	Limit  int64
	Offset int64
}

func (q *Queries) ListMailAddresses(ctx context.Context, arg ListMailAddressesParams) ([]MailAddress, error) {
	rows, err := q.db.QueryContext(ctx, listMailAddresses, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []MailAddress
	for rows.Next() {
		var i MailAddress
		if err := rows.Scan(
			&i.ID,
			&i.Address,
			&i.Active,
			&i.Createdat,
			&i.Updatedat,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateMailAddress = `-- name: UpdateMailAddress :exec
UPDATE mailAddress
SET
    address = ?,
    active = ?,
    updatedAt = CURRENT_TIMESTAMP
WHERE
    id = ?
`

type UpdateMailAddressParams struct {
	Address string
	Active  bool
	ID      string
}

func (q *Queries) UpdateMailAddress(ctx context.Context, arg UpdateMailAddressParams) error {
	_, err := q.db.ExecContext(ctx, updateMailAddress, arg.Address, arg.Active, arg.ID)
	return err
}
