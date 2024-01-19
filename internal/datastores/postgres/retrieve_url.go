package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/eduardolat/tinylink/internal/shortener"
)

func (ds *DataStore) RetrieveURL(shortCode string) (shortener.URLData, error) {
	sqlQuery := `
		SELECT
			short_code, original_url, http_redirect_code,
			is_active, description, tags,
			password, clicks, first_click_at,
			last_click_at, redirects, first_redirect_at,
			last_redirect_at, expires_at, created_by_ip,
			created_by_user_agent, created_at, updated_at
		FROM links WHERE short_code = $1;
	`
	var data shortener.URLData
	err := ds.conn.QueryRow(context.Background(), sqlQuery, shortCode).Scan(
		&data.ShortCode, &data.OriginalURL, &data.HTTPRedirectCode,
		&data.IsActive, &data.Description, &data.Tags,
		&data.Password, &data.Clicks, &data.FirstClickAt,
		&data.LastClickAt, &data.Redirects, &data.FirstRedirectAt,
		&data.LastRedirectAt, &data.ExpiresAt, &data.CreatedByIP,
		&data.CreatedByUserAgent, &data.CreatedAt, &data.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return shortener.URLData{}, errors.New("Link not found")
		}
		return shortener.URLData{}, fmt.Errorf("error retrieving URL: %w", err)
	}
	return data, nil
}
