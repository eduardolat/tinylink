package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/eduardolat/tinylink/internal/shortener"
)

func (ds *DataStore) RetrieveURLByOriginalUrl(originalUrl string) (shortener.URLData, error) {
	sqlQuery := `
			SELECT short_code, original_url, http_redirect_code, is_active, description, tags, password, clicks, first_click_at, last_click_at, redirects, first_redirect_at, last_redirect_at, expires_at, created_by_ip, created_by_user_agent, created_at, updated_at 
			FROM links 
			WHERE original_url = $1;
	`
	var data shortener.URLData
	row := ds.conn.QueryRow(context.Background(), sqlQuery, originalUrl)
	if err := row.Scan(&data.ShortCode, &data.OriginalURL, &data.HTTPRedirectCode, &data.IsActive, &data.Description, &data.Tags, &data.Password, &data.Clicks, &data.FirstClickAt, &data.LastClickAt, &data.Redirects, &data.FirstRedirectAt, &data.LastRedirectAt, &data.ExpiresAt, &data.CreatedByIP, &data.CreatedByUserAgent, &data.CreatedAt, &data.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return shortener.URLData{}, fmt.Errorf("URL not found: %w", err)
		}
		return shortener.URLData{}, fmt.Errorf("error retrieving URL by original URL: %w", err)
	}
	return data, nil
}
