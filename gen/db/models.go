// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"time"
)

type MailAddress struct {
	ID        string
	Address   string
	Active    bool
	Createdat time.Time
	Updatedat time.Time
}

type Result struct {
	ID        string
	Status    string
	Createdat time.Time
	Updatedat time.Time
}
