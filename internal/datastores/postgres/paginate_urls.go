package postgres

import (
	"context"
	"fmt"

	"github.com/eduardolat/tinylink/internal/shortener"
)

func (ds *DataStore) PaginateURLS(params shortener.PaginateURLSParams) (shortener.PaginateURLSResponse, error) {
	var (
		urls   []shortener.URLData
		filter string
		args   []interface{}
	)

	baseQuery := `SELECT short_code, original_url, http_redirect_code, is_active, description, tags, password, clicks, first_click_at, last_click_at, redirects, first_redirect_at, last_redirect_at, expires_at, created_by_ip, created_by_user_agent, created_at, updated_at FROM links`
	if len(params.TagsFilter) > 0 {
		filter = ` WHERE tags && $1`
		args = append(args, params.TagsFilter)
	}

	paginationQuery := fmt.Sprintf("%s%s ORDER BY created_at DESC LIMIT $%d OFFSET $%d", baseQuery, filter, len(args)+1, len(args)+2)
	args = append(args, params.Size, (params.Page-1)*params.Size)

	rows, err := ds.conn.Query(context.Background(), paginationQuery, args...)
	if err != nil {
		return shortener.PaginateURLSResponse{}, fmt.Errorf("error retrieving paginated URLs: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var data shortener.URLData
		if err := rows.Scan(&data.ShortCode, &data.OriginalURL, &data.HTTPRedirectCode, &data.IsActive, &data.Description, &data.Tags, &data.Password, &data.Clicks, &data.FirstClickAt, &data.LastClickAt, &data.Redirects, &data.FirstRedirectAt, &data.LastRedirectAt, &data.ExpiresAt, &data.CreatedByIP, &data.CreatedByUserAgent, &data.CreatedAt, &data.UpdatedAt); err != nil {
			return shortener.PaginateURLSResponse{}, fmt.Errorf("error scanning paginated URLs: %w", err)
		}
		urls = append(urls, data)
	}

	// Contar total de filas
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM links%s", filter)
	var totalItems int
	err = ds.conn.QueryRow(context.Background(), countQuery, args[:len(args)-2]...).Scan(&totalItems)
	if err != nil {
		return shortener.PaginateURLSResponse{}, fmt.Errorf("error counting URLs: %w", err)
	}

	return shortener.PaginateURLSResponse{
		Items:      urls,
		TotalItems: totalItems,
		Page:       params.Page,
		Size:       params.Size,
		TotalPages: (totalItems + params.Size - 1) / params.Size,
	}, nil
}
