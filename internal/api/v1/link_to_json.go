package v1

import (
	"time"

	"github.com/eduardolat/tinylink/internal/database/dbgen"
)

type jsonLink struct {
	ID                 string     `json:"id"`
	ShortCode          string     `json:"short_code"`
	ShortURL           string     `json:"short_url"`
	OriginalURL        string     `json:"original_url"`
	HTTPRedirectCode   int        `json:"http_redirect_code"`
	IsActive           bool       `json:"is_active"`
	Description        *string    `json:"description"`
	Tags               []string   `json:"tags"`
	ExpiresAt          *time.Time `json:"expires_at"`
	CreatedByIP        string     `json:"created_by_ip"`
	CreatedByUserAgent string     `json:"created_by_user_agent"`
	CreatedAt          time.Time  `json:"created_at"`
	UpdatedAt          time.Time  `json:"updated_at"`
}

// linkToJSON converts a dbgen.Link to a jsonLink that has the transformed
// data to return to the client safely as JSON.
func (h *handlers) linkToJSON(link dbgen.Link) jsonLink {
	jl := jsonLink{
		ID:                 link.ID.String(),
		ShortCode:          link.ShortCode,
		ShortURL:           h.shortener.CreateShortLinkFromCode(link.ShortCode),
		OriginalURL:        link.OriginalUrl,
		HTTPRedirectCode:   int(link.HttpRedirectCode),
		IsActive:           link.IsActive,
		Tags:               link.Tags,
		CreatedByIP:        link.CreatedByIp.String,
		CreatedByUserAgent: link.CreatedByUserAgent.String,
		CreatedAt:          link.CreatedAt,
		UpdatedAt:          link.UpdatedAt,
	}

	if link.Description.Valid {
		jl.Description = &link.Description.String
	}

	if link.ExpiresAt.Valid {
		jl.ExpiresAt = &link.ExpiresAt.Time
	}

	return jl
}
