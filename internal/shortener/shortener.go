package shortener

import (
	"errors"

	"github.com/eduardolat/tinylink/internal/logger"
)

var (
	// ErrShortCodeNotAvailable is the error returned when the short code is not available
	ErrShortCodeNotAvailable = errors.New("short code not available")
	// ErrCannotUseDuplicateIfExists is the error returned when the user tries to
	// use DuplicateIfExists when providing a short code
	ErrCannotUseDuplicateIfExists = errors.New("cannot use DuplicateIfExists when providing a short code")
)

type Shortener struct {
	dataStore DataStore
	shortGen  ShortGen
}

func NewShortener(
	dataStore DataStore,
	shortGen ShortGen,
) *Shortener {
	return &Shortener{
		dataStore: dataStore,
		shortGen:  shortGen,
	}
}

// ShortenURL is the function that will be used to shorten a URL
func (c *Shortener) ShortenURL(params StoreURLParams) (string, error) {
	// When a short code is provided, we should not allow the user to
	// duplicate the original URL if it already exists in the database.
	if params.ShortCode != "" && params.DuplicateIfExists {
		return "", ErrCannotUseDuplicateIfExists
	}

	// If the url is already stored, we return the stored data.
	if !params.DuplicateIfExists {
		isAlreadyStored, err := c.dataStore.IsURLAlreadyStored(params.OriginalURL)
		if err != nil {
			return "", err
		}
		if isAlreadyStored {
			existingUrlData, err := c.dataStore.RetrieveURLByOriginalUrl(params.OriginalURL)
			if err != nil {
				return "", err
			}
			return existingUrlData.ShortCode, nil
		}
	}

	// When the user provides a short code, we need to check if it's available.
	// We should not generate a random short code if the user provided one.
	if params.ShortCode != "" {
		isAvailable, err := c.dataStore.IsShortCodeAvailable(params.ShortCode)
		if err != nil {
			return "", err
		}
		if !isAvailable {
			return "", ErrShortCodeNotAvailable
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
				return "", err
			}

			isAvailable, err := c.dataStore.IsShortCodeAvailable(sc)
			if err != nil {
				return "", err
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
			return "", ErrShortCodeNotAvailable
		}
	}

	urlData, err := c.dataStore.StoreURL(params)
	if err != nil {
		return "", err
	}

	return urlData.ShortCode, nil
}

// PaginateURLS is the function that will be used to paginate the URLs
// that were previously shortened
func (c *Shortener) PaginateURLS(params PaginateURLSParams) (PaginateURLSResponse, error) {
	return c.dataStore.PaginateURLS(params)
}

// RetrieveURL is the function that will be used to retrieve a URL
// that was previously shortened
func (c *Shortener) RetrieveURL(shortCode string) (URLData, error) {
	data, err := c.dataStore.RetrieveURL(shortCode)
	if err != nil {
		return URLData{}, err
	}

	return data, nil
}

// UpdateURL is the function that will be used to update a URL
// that was previously shortened
func (c *Shortener) UpdateURL(shortCode string, params UpdateURLParams) (URLData, error) {
	return c.dataStore.UpdateURL(shortCode, params)
}

// DeleteURL is the function that will be used to delete a URL
// that was previously shortened
func (c *Shortener) DeleteURL(shortCode string) error {
	return c.dataStore.DeleteURL(shortCode)
}

// IncrementClicks is the function that will be used to increment the clicks
// of a URL that was previously shortened
func (c *Shortener) IncrementClicks(shortCode string) error {
	return c.dataStore.IncrementClicks(shortCode)
}

// IncrementRedirects is the function that will be used to increment the redirects
// of a URL that was previously shortened
func (c *Shortener) IncrementRedirects(shortCode string) error {
	return c.dataStore.IncrementRedirects(shortCode)
}
