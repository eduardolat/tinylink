package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/eduardolat/tinylink/internal/shortener"
)

func (ds *DataStore) UpdateURL(shortCode string, params shortener.UpdateURLParams) (shortener.URLData, error) {
	sqlQuery := `
			UPDATE links 
			SET http_redirect_code = $2, is_active = $3, description = $4, tags = $5, password = $6, expires_at = $7, updated_at = NOW()
			WHERE short_code = $8
			RETURNING short_code, original_url, http_redirect_code, is_active, description, tags, password, clicks, first_click_at, last_click_at, redirects, first_redirect_at, last_redirect_at, expires_at, created_by_ip, created_by_user_agent, created_at, updated_at;
	`
	var data shortener.URLData
	row := ds.conn.QueryRow(
		context.Background(), sqlQuery,
		params.HTTPRedirectCode, params.IsActive, params.Description,
		params.Tags, params.Password, params.ExpiresAt, shortCode,
	)
	if err := row.Scan(&data.ShortCode, &data.OriginalURL, &data.HTTPRedirectCode, &data.IsActive, &data.Description, &data.Tags, &data.Password, &data.Clicks, &data.FirstClickAt, &data.LastClickAt, &data.Redirects, &data.FirstRedirectAt, &data.LastRedirectAt, &data.ExpiresAt, &data.CreatedByIP, &data.CreatedByUserAgent, &data.CreatedAt, &data.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return shortener.URLData{}, fmt.Errorf("URL not found: %w", err)
		}
		return shortener.URLData{}, fmt.Errorf("error updating URL: %w", err)
	}
	return data, nil
}
