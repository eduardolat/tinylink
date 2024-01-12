package shortener

import (
	"database/sql"
	"time"
)

// URLData is the data structure that will be stored in the database
type URLData struct {
	// IsActive is a flag that indicates if the URL is active or not
	// If the value is false, the URL will not be accessible
	IsActive bool

	// Description is a short description of the URL
	Description sql.NullString

	// Tags is a list of tags that the user can use to categorize the URL
	Tags []string

	// HTTPRedirectCode is the HTTP code that will be used to redirect the user
	// to the original URL
	HTTPRedirectCode HTTPRedirectCode

	// OriginalURL is the URL that the user wants to shorten
	OriginalURL string

	// ShortCode is the short code that will be used to retrieve the URL.
	// Can be sequential, random, uuid, etc.
	// In some data stores, this field can be the primary key
	ShortCode string

	// Password is the password that the user must enter to access the URL
	// If the value is null, the URL will not be password protected
	Password sql.NullString

	// Clicks is the number of times the URL has been clicked.
	Clicks sql.NullInt64

	// FirstClickAt is the timestamp when the URL was first clicked
	FirstClickAt sql.NullTime

	// LastClickAt is the timestamp when the URL was last clicked
	LastClickAt sql.NullTime

	// Redirects is the number of times the URL has been redirected
	Redirects sql.NullInt64

	// FirstRedirectAt is the timestamp when the URL was first redirected
	FirstRedirectAt sql.NullTime

	// LastRedirectAt is the timestamp when the URL was last redirected
	LastRedirectAt sql.NullTime

	// ExpiresAt is the timestamp when the URL will expire
	// If the value is null, the URL will never expire
	ExpiresAt sql.NullTime

	// CreatedByIP is the IP address of the user/server that created the URL
	CreatedByIP sql.NullString

	// CreatedByUserAgent is the user agent of the user/server that created the URL
	CreatedByUserAgent sql.NullString

	// CreatedAt is the timestamp when the URL was created
	// The default value is the current timestamp
	CreatedAt time.Time

	// UpdatedAt is the timestamp when the URL was last updated
	// The default value is the current timestamp
	UpdatedAt time.Time
}

// StoreURLParams is the data structure that will be used to store a new URL
// in the database
type StoreURLParams struct {
	// ShortCode is the short code that will be used to retrieve the URL
	ShortCode string

	// OriginalURL is the URL that the user wants to shorten
	OriginalURL string

	// HTTPRedirectCode is the HTTP code that will be used to redirect the user
	// to the original URL
	HTTPRedirectCode HTTPRedirectCode

	// Description is a short description of the URL
	Description sql.NullString

	// Tags is a list of tags that the user can use to categorize the URL
	Tags []string

	// Password is the password that the user must enter to access the URL
	// If the value is null, the URL will not be password protected
	Password sql.NullString

	// ExpiresAt is the timestamp when the URL will expire
	// If the value is null, the URL will never expire
	ExpiresAt sql.NullTime

	// IsActive is a flag that indicates if the URL is active or not
	// If the value is false, the URL will not be accessible
	IsActive bool

	// CreatedByIP is the IP address of the user/server that created the URL
	CreatedByIP sql.NullString

	// CreatedByUserAgent is the user agent of the user/server that created the URL
	CreatedByUserAgent sql.NullString

	// DuplicateIfExists is a flag that indicates if the shortener
	// must store a new shortened URL with new short code if the original
	// URL already exists in the database. Otherwise, it will return the
	// last shortened URL that was stored for the original URL.
	DuplicateIfExists bool
}

// UpdateURLParams is the data structure that will be used to update an existing URL
type UpdateURLParams struct {
	// Description is a short description of the URL
	Description sql.NullString

	// Tags is a list of tags that the user can use to categorize the URL
	Tags []string

	// HTTPRedirectCode is the HTTP code that will be used to redirect the user
	HTTPRedirectCode HTTPRedirectCode

	// Password is the password that the user must enter to access the URL
	Password sql.NullString

	// ExpiresAt is the timestamp when the URL will expire
	ExpiresAt sql.NullTime

	// IsActive is a flag that indicates if the URL is active or not
	IsActive bool
}

// DataStore is the interface that will be implemented by all the
// database adapters
type DataStore interface {
	StoreURL(params StoreURLParams) (URLData, error)
	RetrieveURL(shortCode string) (URLData, error)
	UpdateURL(shortCode string, params UpdateURLParams) (URLData, error)
	DeleteURL(shortCode string) error
	IncrementClicks(shortCode string) error
	IncrementRedirects(shortCode string) error
	GetURLsByTag(tag string) ([]URLData, error)
	GetActiveURLs() ([]URLData, error)
	GetExpiredURLs() ([]URLData, error)
}
