-- name: Links_CountAll :one
SELECT COUNT(*) FROM links;

-- name: Links_ExistsByShortCode :one
SELECT EXISTS (
    SELECT 1
    FROM links
    WHERE short_code = $1
) AS exists;

-- name: Links_ExistsByOriginalURL :one
SELECT EXISTS (
    SELECT 1
    FROM links
    WHERE original_url = $1
) AS exists;

-- name: Links_Get :one
SELECT * FROM links WHERE id = $1;

-- name: Links_GetByShortCode :one
SELECT * FROM links WHERE short_code = $1;

-- name: Links_GetByOriginalURL :one
SELECT * FROM links WHERE original_url = $1;

-- name: Links_Create :one
INSERT INTO links (
    short_code,
    original_url,
    http_redirect_code,
    is_active,
    description,
    tags,
    password,
    expires_at,
    created_by_ip,
    created_by_user_agent
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8,
    $9,
    $10
) RETURNING *;

-- name: Links_Update :one
UPDATE links SET
  short_code = COALESCE($2, short_code),
  original_url = COALESCE($3, original_url),
  http_redirect_code = COALESCE($4, http_redirect_code),
  is_active = COALESCE($5, is_active),
  description = COALESCE($6, description),
  tags = COALESCE($7, tags),
  password = COALESCE($8, password),
  expires_at = COALESCE($9, expires_at),
  updated_at = NOW()
WHERE id = $1 RETURNING *;

-- name: Links_Delete :exec
DELETE FROM links WHERE id = $1;

-- name: Links_Paginate :many
SELECT * FROM links
WHERE (
  sqlc.narg('filter_is_active')::BOOLEAN IS NULL
  OR
  is_active = sqlc.narg('filter_is_active')::BOOLEAN
)
AND (
  sqlc.narg('filter_original_url')::TEXT IS NULL
  OR
  original_url ILIKE sqlc.narg('filter_original_url')::TEXT
)
AND (
  sqlc.narg('filter_short_code')::TEXT IS NULL
  OR
  short_code ILIKE sqlc.narg('filter_short_code')::TEXT
)
AND (
  sqlc.narg('filter_description')::TEXT IS NULL
  OR
  description ILIKE sqlc.narg('filter_description')::TEXT
)
AND (
  ARRAY_LENGTH(sqlc.arg('filter_tags')::TEXT[], 1) = 0
  OR
  tags && sqlc.arg('filter_tags')::TEXT[]
)
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;

-- name: Links_PaginateCountTotalMatches :one
SELECT COUNT(*) FROM links
WHERE (
  sqlc.narg('filter_is_active')::BOOLEAN IS NULL
  OR
  is_active = sqlc.narg('filter_is_active')::BOOLEAN
)
AND (
  sqlc.narg('filter_original_url')::TEXT IS NULL
  OR
  original_url ILIKE sqlc.narg('filter_original_url')::TEXT
)
AND (
  sqlc.narg('filter_short_code')::TEXT IS NULL
  OR
  short_code ILIKE sqlc.narg('filter_short_code')::TEXT
)
AND (
  sqlc.narg('filter_description')::TEXT IS NULL
  OR
  description ILIKE sqlc.narg('filter_description')::TEXT
)
AND (
  ARRAY_LENGTH(sqlc.arg('filter_tags')::TEXT[], 1) = 0
  OR
  tags && sqlc.arg('filter_tags')::TEXT[]
);
