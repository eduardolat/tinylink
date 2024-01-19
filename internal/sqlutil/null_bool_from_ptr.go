package sqlutil

import (
	"database/sql"
)

// NullBoolFromPtr returns a sql.NullBool from a pointer to a bool.
//
// If the pointer is nil, it returns a sql.NullBool with Valid set to false.
//
// Otherwise, it returns a sql.NullBool with Valid set to true and Bool set to
// the value pointed to by ptr.
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
