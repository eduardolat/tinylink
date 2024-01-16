package datastores

import "github.com/eduardolat/tinylink/internal/shortener"

// CreatePaginationResponse is a helper function that creates a pagination response
// based on the provided parameters
func CreatePaginationResponse(
	page int,
	size int,
	totalItems int,
	items []shortener.URLData,
) shortener.PaginateURLSResponse {
	totalPages := totalItems / size
	if totalItems%size > 0 {
		totalPages++
	}

	prevPage := page - 1
	if prevPage < 1 {
		prevPage = 0
	}

	nextPage := page + 1
	if nextPage > totalPages {
		nextPage = 0
	}

	return shortener.PaginateURLSResponse{
		PrevPage:   prevPage,
		NextPage:   nextPage,
		Page:       page,
		Size:       size,
		TotalPages: totalPages,
		TotalItems: totalItems,
		Items:      items,
	}
}
