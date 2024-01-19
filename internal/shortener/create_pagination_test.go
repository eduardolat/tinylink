package shortener

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatePagination(t *testing.T) {
	// Test case 1: Page 1, size 2, totalItems 2
	response := createPagination(1, 2, 2)
	assert.Equal(t, 1, response.Page)
	assert.Equal(t, 2, response.Size)
	assert.Equal(t, 1, response.TotalPages)
	assert.Equal(t, 2, response.TotalItems)
	assert.Equal(t, 0, response.PrevPage)
	assert.Equal(t, 0, response.NextPage)

	// Test case 2: Page 1, size 1, totalItems 2
	response = createPagination(1, 1, 2)
	assert.Equal(t, 1, response.Page)
	assert.Equal(t, 1, response.Size)
	assert.Equal(t, 2, response.TotalPages)
	assert.Equal(t, 2, response.TotalItems)
	assert.Equal(t, 0, response.PrevPage)
	assert.Equal(t, 2, response.NextPage)

	// Test case 3: Page 2, size 1, totalItems 2
	response = createPagination(2, 1, 2)
	assert.Equal(t, 2, response.Page)
	assert.Equal(t, 1, response.Size)
	assert.Equal(t, 2, response.TotalPages)
	assert.Equal(t, 2, response.TotalItems)
	assert.Equal(t, 1, response.PrevPage)
	assert.Equal(t, 0, response.NextPage)
}
