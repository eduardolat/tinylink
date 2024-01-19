package sqlutil

import (
	"database/sql"
)

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
