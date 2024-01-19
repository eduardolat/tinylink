package shortener

// createPaginationDBArgs is the function that will be used to create the
// pagination arguments for the database based on the page and size
// provided by the user
func createPaginationDBArgs(page int, size int) struct {
	Limit  int32
	Offset int32
} {
	if page <= 0 || size <= 0 {
		return struct {
			Limit  int32
			Offset int32
		}{
			Limit:  0,
			Offset: 0,
		}
	}

	limit := int32(size)
	offset := int32((page - 1) * size)

	return struct {
		Limit  int32
		Offset int32
	}{
		Limit:  limit,
		Offset: offset,
	}
}
