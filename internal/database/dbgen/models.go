// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0

package dbgen

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Link struct {
	ID                 uuid.UUID
	ShortCode          string
	OriginalUrl        string
	HttpRedirectCode   int16
	IsActive           bool
	Description        sql.NullString
	Tags               []string
	Password           sql.NullString
	ExpiresAt          sql.NullTime
	CreatedByIp        sql.NullString
	CreatedByUserAgent sql.NullString
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

type Visit struct {
	ID           uuid.UUID
	LinkID       uuid.UUID
	Ip           string
	UserAgent    string
	Referer      sql.NullString
	IsRedirected bool
	CreatedAt    time.Time
}
