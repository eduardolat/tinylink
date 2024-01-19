package postgres

import (
	"context"
	"fmt"

	"github.com/eduardolat/tinylink/internal/shortener"
)

func (ds *DataStore) StoreURL(params shortener.StoreURLParams) (shortener.URLData, error) {
	sqlQuery := `
		INSERT INTO links (
			short_code, original_url, http_redirect_code, is_active,
			description, tags, password, expires_at, created_by_ip,
			created_by_user_agent, created_at, updated_at
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, NOW(), NOW());
	`
	_, err := ds.conn.Exec(
		context.Background(), sqlQuery,
		params.ShortCode, params.OriginalURL, params.HTTPRedirectCode, params.IsActive,
		params.Description, params.Tags, params.Password, params.ExpiresAt,
		params.CreatedByIP, params.CreatedByUserAgent,
	)
	if err != nil {
		return shortener.URLData{}, fmt.Errorf("error storing URL: %w", err)
	}

	urlData, err := ds.RetrieveURL(params.ShortCode)
	return urlData, err
}
