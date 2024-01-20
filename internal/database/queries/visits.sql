-- name: Visits_Get :one
SELECT * FROM visits WHERE id = $1;

-- name: Visits_Create :one
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
) RETURNING *;

-- name: Visits_SetIsRedirected :one
UPDATE visits SET
  is_redirected = $2
WHERE id = $1 RETURNING *;

-- name: Visits_CountAllForLink :one
SELECT COUNT(*) FROM visits
WHERE link_id = $1;

-- name: Visits_CountAllRedirectedForLink :one
SELECT COUNT(*) FROM visits
WHERE link_id = $1
AND is_redirected = TRUE;

-- name: Visits_PaginateForLink :many
SELECT * FROM visits
WHERE link_id = $1
AND (
  sqlc.narg('filter_ip')::TEXT IS NULL
  OR
  ip ILIKE CONCAT('%', sqlc.narg('filter_ip')::TEXT, '%')
)
AND (
  sqlc.narg('filter_is_redirected')::BOOLEAN IS NULL
  OR
  is_redirected = sqlc.narg('filter_is_redirected')::BOOLEAN
)
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;

-- name: Visits_PaginateForLinkCountTotalMatches :one
SELECT COUNT(*) FROM visits
WHERE link_id = $1
AND (
  sqlc.narg('filter_ip')::TEXT IS NULL
  OR
  ip ILIKE CONCAT('%', sqlc.narg('filter_ip')::TEXT, '%')
)
AND (
  sqlc.narg('filter_is_redirected')::BOOLEAN IS NULL
  OR
  is_redirected = sqlc.narg('filter_is_redirected')::BOOLEAN
);
