// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: visits.sql

package dbgen

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const visits_CountAllForLink = `-- name: Visits_CountAllForLink :one
SELECT COUNT(*) FROM visits
WHERE link_id = $1
`

func (q *Queries) Visits_CountAllForLink(ctx context.Context, linkID uuid.UUID) (int64, error) {
	row := q.db.QueryRowContext(ctx, visits_CountAllForLink, linkID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const visits_CountAllRedirectedForLink = `-- name: Visits_CountAllRedirectedForLink :one
SELECT COUNT(*) FROM visits
WHERE link_id = $1
AND is_redirected = TRUE
`

func (q *Queries) Visits_CountAllRedirectedForLink(ctx context.Context, linkID uuid.UUID) (int64, error) {
	row := q.db.QueryRowContext(ctx, visits_CountAllRedirectedForLink, linkID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const visits_Create = `-- name: Visits_Create :one
INSERT INTO visits (
    link_id,
    ip,
    user_agent,
    referer,
    is_redirected
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
) RETURNING id, link_id, ip, user_agent, referer, is_redirected, created_at
`

type Visits_CreateParams struct {
	LinkID       uuid.UUID
	Ip           string
	UserAgent    string
	Referer      sql.NullString
	IsRedirected bool
}

func (q *Queries) Visits_Create(ctx context.Context, arg Visits_CreateParams) (Visit, error) {
	row := q.db.QueryRowContext(ctx, visits_Create,
		arg.LinkID,
		arg.Ip,
		arg.UserAgent,
		arg.Referer,
		arg.IsRedirected,
	)
	var i Visit
	err := row.Scan(
		&i.ID,
		&i.LinkID,
		&i.Ip,
		&i.UserAgent,
		&i.Referer,
		&i.IsRedirected,
		&i.CreatedAt,
	)
	return i, err
}

const visits_Get = `-- name: Visits_Get :one
SELECT id, link_id, ip, user_agent, referer, is_redirected, created_at FROM visits WHERE id = $1
`

func (q *Queries) Visits_Get(ctx context.Context, id uuid.UUID) (Visit, error) {
	row := q.db.QueryRowContext(ctx, visits_Get, id)
	var i Visit
	err := row.Scan(
		&i.ID,
		&i.LinkID,
		&i.Ip,
		&i.UserAgent,
		&i.Referer,
		&i.IsRedirected,
		&i.CreatedAt,
	)
	return i, err
}

const visits_PaginateForLink = `-- name: Visits_PaginateForLink :many
SELECT id, link_id, ip, user_agent, referer, is_redirected, created_at FROM visits
WHERE link_id = $1
AND (
  $4::TEXT IS NULL
  OR
  ip ILIKE CONCAT('%', $4::TEXT, '%')
)
AND (
  $5::BOOLEAN IS NULL
  OR
  is_redirected = $5::BOOLEAN
)
ORDER BY created_at DESC
LIMIT $2 OFFSET $3
`

type Visits_PaginateForLinkParams struct {
	LinkID             uuid.UUID
	Limit              int32
	Offset             int32
	FilterIp           sql.NullString
	FilterIsRedirected sql.NullBool
}

func (q *Queries) Visits_PaginateForLink(ctx context.Context, arg Visits_PaginateForLinkParams) ([]Visit, error) {
	rows, err := q.db.QueryContext(ctx, visits_PaginateForLink,
		arg.LinkID,
		arg.Limit,
		arg.Offset,
		arg.FilterIp,
		arg.FilterIsRedirected,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Visit
	for rows.Next() {
		var i Visit
		if err := rows.Scan(
			&i.ID,
			&i.LinkID,
			&i.Ip,
			&i.UserAgent,
			&i.Referer,
			&i.IsRedirected,
			&i.CreatedAt,
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

const visits_PaginateForLinkCountTotalMatches = `-- name: Visits_PaginateForLinkCountTotalMatches :one
SELECT COUNT(*) FROM visits
WHERE link_id = $1
AND (
  $2::TEXT IS NULL
  OR
  ip ILIKE CONCAT('%', $2::TEXT, '%')
)
AND (
  $3::BOOLEAN IS NULL
  OR
  is_redirected = $3::BOOLEAN
)
`

type Visits_PaginateForLinkCountTotalMatchesParams struct {
	LinkID             uuid.UUID
	FilterIp           sql.NullString
	FilterIsRedirected sql.NullBool
}

func (q *Queries) Visits_PaginateForLinkCountTotalMatches(ctx context.Context, arg Visits_PaginateForLinkCountTotalMatchesParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, visits_PaginateForLinkCountTotalMatches, arg.LinkID, arg.FilterIp, arg.FilterIsRedirected)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const visits_SetIsRedirected = `-- name: Visits_SetIsRedirected :one
UPDATE visits SET
  is_redirected = $2
WHERE id = $1 RETURNING id, link_id, ip, user_agent, referer, is_redirected, created_at
`

type Visits_SetIsRedirectedParams struct {
	ID           uuid.UUID
	IsRedirected bool
}

func (q *Queries) Visits_SetIsRedirected(ctx context.Context, arg Visits_SetIsRedirectedParams) (Visit, error) {
	row := q.db.QueryRowContext(ctx, visits_SetIsRedirected, arg.ID, arg.IsRedirected)
	var i Visit
	err := row.Scan(
		&i.ID,
		&i.LinkID,
		&i.Ip,
		&i.UserAgent,
		&i.Referer,
		&i.IsRedirected,
		&i.CreatedAt,
	)
	return i, err
}
