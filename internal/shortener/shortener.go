package shortener

import (
	"context"
	"database/sql"
	"errors"

	"github.com/eduardolat/tinylink/internal/config"
	"github.com/eduardolat/tinylink/internal/database/dbgen"
	"github.com/eduardolat/tinylink/internal/logger"
	"github.com/google/uuid"
)

var (
	// ErrShortCodeNotAvailable is the error returned when the short code is not available
	ErrShortCodeNotAvailable = errors.New("short code not available")
	// ErrCannotUseDuplicateIfExists is the error returned when the user tries to
	// use duplicateIfExists when providing a short code
	ErrCannotUseDuplicateIfExists = errors.New("cannot use duplicate if exists when providing a predefined short code")
)

type Shortener struct {
	env      *config.Env
	dbg      *dbgen.Queries
	shortGen ShortGen
}

func NewShortener(
	env *config.Env,
	dbg *dbgen.Queries,
	shortGen ShortGen,
) *Shortener {
	return &Shortener{
		env:      env,
		dbg:      dbg,
		shortGen: shortGen,
	}
}

// ShortenURL is the function that will be used to shorten a URL
func (c *Shortener) ShortenURL(
	params dbgen.Links_CreateParams,
	// duplicateIfExists is a boolean that indicates if the user wants to
	// re-shorten an URL that was already shortened. Otherwise, we return
	// the existing short code.
	duplicateIfExists bool,
) (dbgen.Link, error) {
	// When a short code is provided, we should not allow the user to
	// duplicate the original URL if it already exists in the database.
	if params.ShortCode != "" && duplicateIfExists {
		return dbgen.Link{}, ErrCannotUseDuplicateIfExists
	}

	// If the url is already stored, we return the stored data.
	if !duplicateIfExists {
		isAlreadyStored, err := c.dbg.Links_ExistsByOriginalURL(
			context.Background(),
			params.OriginalUrl,
		)
		if err != nil {
			return dbgen.Link{}, err
		}
		if isAlreadyStored {
			existing, err := c.dbg.Links_GetByOriginalURL(
				context.Background(),
				params.OriginalUrl,
			)

			return existing, err
		}
	}

	// When the user provides a short code, we need to check if it's available.
	// We should not generate a random short code if the user provided one.
	if params.ShortCode != "" {
		isAvailable, err := c.dbg.Links_ExistsByShortCode(
			context.Background(),
			params.ShortCode,
		)
		if err != nil {
			return dbgen.Link{}, err
		}
		if !isAvailable {
			return dbgen.Link{}, ErrShortCodeNotAvailable
		}
	}

	// When the user does not provide a short code, we need to generate a random one
	// and check if it's available.
	// If it's not available, we try again 5 times before giving up.
	if params.ShortCode == "" {
		maxTries := 5

		for i := 0; i < maxTries; i++ {
			sc, err := c.shortGen.Generate()
			if err != nil {
				return dbgen.Link{}, err
			}

			isAvailable, err := c.dbg.Links_ExistsByShortCode(
				context.Background(),
				sc,
			)
			if err != nil {
				return dbgen.Link{}, err
			}

			if isAvailable {
				params.ShortCode = sc
				break
			}

			logger.Warn(
				"short code collision detected, retrying",
				"short_code", sc,
				"try", i+1,
				"max_tries", maxTries,
			)
		}

		if params.ShortCode == "" {
			return dbgen.Link{}, ErrShortCodeNotAvailable
		}
	}

	link, err := c.dbg.Links_Create(
		context.Background(),
		params,
	)
	return link, err
}

type PaginateLinksParams struct {
	Page              int
	Size              int
	FilterIsActive    sql.NullBool
	FilterOriginalUrl sql.NullString
	FilterShortCode   sql.NullString
	FilterDescription sql.NullString
	FilterTags        []string
}

type PaginateLinksResponse struct {
	Page       int
	PrevPage   int
	NextPage   int
	TotalPages int
	TotalItems int
	Items      []dbgen.Link
}

// PaginateURLS is the function that will be used to paginate the URLs
// that were previously shortened
func (c *Shortener) PaginateURLS(params PaginateLinksParams) (PaginateLinksResponse, error) {
	totalCount, err := c.dbg.Links_PaginateCountTotalMatches(
		context.Background(),
		dbgen.Links_PaginateCountTotalMatchesParams{
			FilterIsActive:    params.FilterIsActive,
			FilterOriginalUrl: params.FilterOriginalUrl,
			FilterShortCode:   params.FilterShortCode,
			FilterDescription: params.FilterDescription,
			FilterTags:        params.FilterTags,
		},
	)
	if err != nil {
		return PaginateLinksResponse{}, err
	}

	args := createPaginationDBArgs(params.Page, params.Size)
	items, err := c.dbg.Links_Paginate(
		context.Background(),
		dbgen.Links_PaginateParams{
			FilterIsActive:    params.FilterIsActive,
			FilterOriginalUrl: params.FilterOriginalUrl,
			FilterShortCode:   params.FilterShortCode,
			FilterDescription: params.FilterDescription,
			FilterTags:        params.FilterTags,
			Limit:             args.Limit,
			Offset:            args.Offset,
		},
	)
	pagination := createPagination(params.Page, params.Size, int(totalCount))

	return PaginateLinksResponse{
		Page:       pagination.Page,
		PrevPage:   pagination.PrevPage,
		NextPage:   pagination.NextPage,
		TotalPages: pagination.TotalPages,
		TotalItems: pagination.TotalItems,
		Items:      items,
	}, nil
}

// RetrieveURL is the function that will be used to retrieve a URL
// that was previously shortened
func (c *Shortener) RetrieveURL(id uuid.UUID) (dbgen.Link, error) {
	link, err := c.dbg.Links_Get(
		context.Background(),
		id,
	)
	return link, err
}

// GetByShortCode is the function that will be used to retrieve a URL
// that was previously shortened by its short code
func (c *Shortener) GetByShortCode(shortCode string) (dbgen.Link, error) {
	link, err := c.dbg.Links_GetByShortCode(
		context.Background(),
		shortCode,
	)
	return link, err
}

// GetByOriginalURL is the function that will be used to retrieve a URL
// that was previously shortened by its original URL
func (c *Shortener) GetByOriginalURL(originalURL string) (dbgen.Link, error) {
	link, err := c.dbg.Links_GetByOriginalURL(
		context.Background(),
		originalURL,
	)
	return link, err
}

// UpdateURL is the function that will be used to update a URL
// that was previously shortened
func (c *Shortener) UpdateURL(id uuid.UUID, params dbgen.Links_UpdateParams) (dbgen.Link, error) {
	params.ID = id
	link, err := c.dbg.Links_Update(
		context.Background(),
		params,
	)
	return link, err
}

// DeleteURL is the function that will be used to delete a URL
// that was previously shortened
func (c *Shortener) DeleteURL(id uuid.UUID) error {
	err := c.dbg.Links_Delete(
		context.Background(),
		id,
	)
	return err
}

// CreateLinkVisit is the function that will be used to create a visit
// for a shortened URL
func (c *Shortener) CreateLinkVisit(linkID uuid.UUID, params dbgen.Visits_CreateParams) (dbgen.Visit, error) {
	params.LinkID = linkID
	visit, err := c.dbg.Visits_Create(
		context.Background(),
		params,
	)
	return visit, err
}

// SetLinkVisitAsRedirected is the function that will be used to set a visit
// as redirected
func (c *Shortener) SetLinkVisitAsRedirected(id uuid.UUID) (dbgen.Visit, error) {
	visit, err := c.dbg.Visits_SetIsRedirected(
		context.Background(),
		dbgen.Visits_SetIsRedirectedParams{
			ID:           id,
			IsRedirected: true,
		},
	)
	return visit, err
}

// GetTotalLinkVisits is the function that will be used to get the total
// visits for a shortened URL
func (c *Shortener) GetTotalLinkVisits(linkID uuid.UUID) (int64, error) {
	count, err := c.dbg.Visits_CountAllForLink(
		context.Background(),
		linkID,
	)
	return count, err
}

// GetTotalLinkRedirectedVisits is the function that will be used to get the total
// visits for a shortened URL that were redirected
func (c *Shortener) GetTotalLinkRedirectedVisits(linkID uuid.UUID) (int64, error) {
	count, err := c.dbg.Visits_CountAllRedirectedForLink(
		context.Background(),
		linkID,
	)
	return count, err
}

type PaginateLinkVisitsParams struct {
	Page               int
	Size               int
	FilterIp           sql.NullString
	FilterIsRedirected sql.NullBool
}

type PaginateLinkVisitsResponse struct {
	Page       int
	PrevPage   int
	NextPage   int
	TotalPages int
	TotalItems int
	Items      []dbgen.Visit
}

// PaginateLinkVisits is the function that will be used to paginate the visits
// for a shortened URL
func (c *Shortener) PaginateLinkVisits(
	linkID uuid.UUID,
	params PaginateLinkVisitsParams,
) (PaginateLinkVisitsResponse, error) {
	totalCount, err := c.dbg.Visits_PaginateForLinkCountTotalMatches(
		context.Background(),
		dbgen.Visits_PaginateForLinkCountTotalMatchesParams{
			LinkID:             linkID,
			FilterIp:           params.FilterIp,
			FilterIsRedirected: params.FilterIsRedirected,
		},
	)
	if err != nil {
		return PaginateLinkVisitsResponse{}, err
	}

	args := createPaginationDBArgs(params.Page, params.Size)
	items, err := c.dbg.Visits_PaginateForLink(
		context.Background(),
		dbgen.Visits_PaginateForLinkParams{
			LinkID:             linkID,
			FilterIp:           params.FilterIp,
			FilterIsRedirected: params.FilterIsRedirected,
			Limit:              args.Limit,
			Offset:             args.Offset,
		},
	)
	pagination := createPagination(params.Page, params.Size, int(totalCount))

	return PaginateLinkVisitsResponse{
		Page:       pagination.Page,
		PrevPage:   pagination.PrevPage,
		NextPage:   pagination.NextPage,
		TotalPages: pagination.TotalPages,
		TotalItems: pagination.TotalItems,
		Items:      items,
	}, nil
}

// CreateShortURL is the function that will be used to create a short URL
// from a short code and the base URL
func (c *Shortener) CreateShortURL(shortCode string) string {
	url := ""
	if c.env.TL_URL != nil {
		url = *c.env.TL_URL
	}
	if url[len(url)-1:] != "/" {
		url += "/"
	}

	return url + shortCode
}
