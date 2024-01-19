package shortener

// createPaginationResponse is the response returned by the createPagination
type createPaginationResponse struct {
	PrevPage   int
	NextPage   int
	Page       int
	Size       int
	TotalPages int
	TotalItems int
}

// createPagination is a helper function that creates a pagination response
// calculated fields based on the provided parameters
func createPagination(
	page int,
	size int,
	totalItems int,
) createPaginationResponse {
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

	return createPaginationResponse{
		PrevPage:   prevPage,
		NextPage:   nextPage,
		Page:       page,
		Size:       size,
		TotalPages: totalPages,
		TotalItems: totalItems,
	}
}
