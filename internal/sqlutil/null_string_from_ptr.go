package sqlutil

import (
	"database/sql"
)

// NullStringFromPtr returns a sql.NullString from a pointer to a string.
//
// If the pointer is nil, it returns a sql.NullString with Valid set to false.
//
// Otherwise, it returns a sql.NullString with Valid set to true and String set
// to the value pointed to by ptr.
func NullStringFromPtr(ptr *string) sql.NullString {
	if ptr == nil {
		return sql.NullString{
			Valid: false,
		}
	}
	return sql.NullString{
		Valid:  true,
		String: *ptr,
	}
}
