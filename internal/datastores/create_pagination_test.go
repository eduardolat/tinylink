package datastores

import (
	"testing"

	"github.com/eduardolat/tinylink/internal/shortener"
	"github.com/stretchr/testify/assert"
)

func TestCreatePaginationResponse(t *testing.T) {
	// Prepare test data
	items := []shortener.URLData{
		{OriginalURL: "http://example.com/1", ShortCode: "a"},
		{OriginalURL: "http://example.com/2", ShortCode: "b"},
	}

	// Test case 1: Page 1, size 2, totalItems 2
	response := CreatePaginationResponse(1, 2, 2, items)
	assert.Equal(t, 1, response.Page)
	assert.Equal(t, 2, response.Size)
	assert.Equal(t, 1, response.TotalPages)
	assert.Equal(t, 2, response.TotalItems)
	assert.Equal(t, 0, response.PrevPage)
	assert.Equal(t, 0, response.NextPage)
	assert.Equal(t, items, response.Items)

	// Test case 2: Page 1, size 1, totalItems 2
	response = CreatePaginationResponse(1, 1, 2, items)
	assert.Equal(t, 1, response.Page)
	assert.Equal(t, 1, response.Size)
	assert.Equal(t, 2, response.TotalPages)
	assert.Equal(t, 2, response.TotalItems)
	assert.Equal(t, 0, response.PrevPage)
	assert.Equal(t, 2, response.NextPage)
	assert.Equal(t, items, response.Items)

	// Test case 3: Page 2, size 1, totalItems 2
	response = CreatePaginationResponse(2, 1, 2, items)
	assert.Equal(t, 2, response.Page)
	assert.Equal(t, 1, response.Size)
	assert.Equal(t, 2, response.TotalPages)
	assert.Equal(t, 2, response.TotalItems)
	assert.Equal(t, 1, response.PrevPage)
	assert.Equal(t, 0, response.NextPage)
	assert.Equal(t, items, response.Items)
}
