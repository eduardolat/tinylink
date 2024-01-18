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
	HTTPRedirectCode int

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
	// if not provided, the adapter will generate a new random one
	ShortCode string

	// OriginalURL is the URL that the user wants to shorten
	OriginalURL string

	// HTTPRedirectCode is the HTTP code that will be used to redirect the user
	// to the original URL
	HTTPRedirectCode int

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
	HTTPRedirectCode int

	// Password is the password that the user must enter to access the URL
	Password sql.NullString

	// ExpiresAt is the timestamp when the URL will expire
	ExpiresAt sql.NullTime

	// IsActive is a flag that indicates if the URL is active or not
	IsActive bool
}

// PaginateURLSParams is the data structure that will be used to paginate URLs
type PaginateURLSParams struct {
	// Page is the page number, defaults to 1
	Page int
	// Size is the page size, defaults to 10
	Size int
	// TagsFilter is a list of tags that the user can use to filter the URLs
	TagsFilter []string
}

// PaginateURLSResponse is the data structure that will be returned when
// paginating URLs
type PaginateURLSResponse struct {
	// PrevPage is the previous page number
	PrevPage int
	// NextPage is the next page number
	NextPage int
	// Page is the page number
	Page int
	// Size is the page size
	Size int
	// TotalPages is the total number of pages
	TotalPages int
	// TotalItems is the total number of records
	TotalItems int
	// Items is the slice of URLs for the current page
	Items []URLData
}

// DataStore is the interface that will be implemented by all the
// database adapters
type DataStore interface {
	// AutoMigrate will create or update the database schema.
	//
	// It should be called before the application starts.
	//
	// It should keep a record of the schema version in the database
	// and run the necessary migrations to update the schema to the
	// latest version without losing or corrupting data.
	//
	// Is the responsability of the adapter to avoid data loss or
	// corruption when running migrations independently of the
	// database engine and schema version.
	AutoMigrate() error

	// IsShortCodeAvailable checks if a given short code is available for use.
	// Returns true if the short code is available, false otherwise.
	IsShortCodeAvailable(shortCode string) (bool, error)

	// IsURLAlreadyStored checks if a given URL is already stored in the database.
	// Returns true if the URL is stored, false otherwise.
	IsURLAlreadyStored(originalURL string) (bool, error)

	// StoreURL stores a new URL with the given parameters.
	// Returns the stored URL data.
	StoreURL(params StoreURLParams) (URLData, error)

	// PaginateURLS retrieves a slice of URL data for a given page and page size.
	// Returns a slice of URL data.
	PaginateURLS(params PaginateURLSParams) (PaginateURLSResponse, error)

	// RetrieveURL retrieves the URL data for a given short code.
	// Returns the URL data.
	RetrieveURL(shortCode string) (URLData, error)

	// RetrieveURLByOriginalUrl retrieves the URL data for a given original URL.
	// Returns the URL data.
	RetrieveURLByOriginalUrl(originalUrl string) (URLData, error)

	// UpdateURL updates the URL data for a given short code with the given parameters.
	// Returns the updated URL data.
	UpdateURL(shortCode string, params UpdateURLParams) (URLData, error)

	// DeleteURL deletes the URL data for a given short code.
	// Returns an error if the deletion was unsuccessful.
	DeleteURL(shortCode string) error

	// IncrementClicks increments the click count for a given short code.
	// Returns an error if the increment was unsuccessful.
	IncrementClicks(shortCode string) error

	// IncrementRedirects increments the redirect count for a given short code.
	// Returns an error if the increment was unsuccessful.
	IncrementRedirects(shortCode string) error
}
