package sqlutil

import (
	"database/sql"
)

func NullBoolFromPtr(ptr *bool) sql.NullBool {
	if ptr == nil {
		return sql.NullBool{
			Valid: false,
		}
	}
	return sql.NullBool{
		Valid: true,
		Bool:  *ptr,
	}
}
