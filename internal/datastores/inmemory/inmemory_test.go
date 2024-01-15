package inmemory

import (
	"database/sql"
	"testing"
	"time"

	"github.com/eduardolat/tinylink/internal/shortener"
	"github.com/stretchr/testify/assert"
)

func TestInMemoryDataStore(t *testing.T) {
	ds := NewDataStore()

	// Test StoreURL
	params := shortener.StoreURLParams{
		IsActive:           true,
		Description:        sql.NullString{String: "Test URL", Valid: true},
		Tags:               []string{"test"},
		HTTPRedirectCode:   302,
		OriginalURL:        "http://example.com",
		ShortCode:          "abc123",
		Password:           sql.NullString{String: "password", Valid: true},
		ExpiresAt:          sql.NullTime{Time: time.Now().Add(24 * time.Hour), Valid: true},
		CreatedByIP:        sql.NullString{String: "127.0.0.1", Valid: true},
		CreatedByUserAgent: sql.NullString{String: "Mozilla/5.0", Valid: true},
	}
	storedData, err := ds.StoreURL(params)
	assert.NoError(t, err)
	assert.Equal(t, params.OriginalURL, storedData.OriginalURL)

	// Test IsShortCodeAvailable
	isAvailable, err := ds.IsShortCodeAvailable("abc123")
	assert.NoError(t, err)
	assert.False(t, isAvailable)

	// Test IsURLAlreadyStored
	isStored, err := ds.IsURLAlreadyStored("http://example.com")
	assert.NoError(t, err)
	assert.True(t, isStored)

	// Test RetrieveURL
	retrievedData, err := ds.RetrieveURL("abc123")
	assert.NoError(t, err)
	assert.Equal(t, storedData, retrievedData)

	// Test UpdateURL
	updateParams := shortener.UpdateURLParams{
		IsActive:         false,
		Description:      sql.NullString{String: "Updated URL", Valid: true},
		Tags:             []string{"updated"},
		HTTPRedirectCode: 301,
		Password:         sql.NullString{String: "newpassword", Valid: true},
		ExpiresAt:        sql.NullTime{Time: time.Now().Add(48 * time.Hour), Valid: true},
	}
	updatedData, err := ds.UpdateURL("abc123", updateParams)
	assert.NoError(t, err)
	assert.Equal(t, updateParams.Description, updatedData.Description)

	// Test IncrementClicks
	err = ds.IncrementClicks("abc123")
	assert.NoError(t, err)
	retrievedData, _ = ds.RetrieveURL("abc123")
	assert.Equal(t, int64(1), retrievedData.Clicks.Int64)

	// Test IncrementRedirects
	err = ds.IncrementRedirects("abc123")
	assert.NoError(t, err)
	retrievedData, _ = ds.RetrieveURL("abc123")
	assert.Equal(t, int64(1), retrievedData.Redirects.Int64)

	// Test GetURLsByTag
	urlsByTag, err := ds.GetURLsByTag("updated")
	assert.NoError(t, err)
	assert.Equal(t, len(urlsByTag), 1)

	// Test GetActiveURLs
	activeURLs, err := ds.GetActiveURLs()
	assert.NoError(t, err)
	assert.Equal(t, len(activeURLs), 0)

	// Test GetExpiredURLs
	expiredURLs, err := ds.GetExpiredURLs()
	assert.NoError(t, err)
	assert.Equal(t, len(expiredURLs), 0)

	// Test DeleteURL
	err = ds.DeleteURL("abc123")
	assert.NoError(t, err)
	_, err = ds.RetrieveURL("abc123")
	assert.Error(t, err)
}
