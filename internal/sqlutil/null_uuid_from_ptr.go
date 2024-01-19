package sqlutil

import (
	"github.com/google/uuid"
)

// NullUUIDFromPtr returns a uuid.NullUUID from a pointer to a uuid.UUID.
//
// If the pointer is nil, it returns a uuid.NullUUID with Valid set to false.
//
// Otherwise, it returns a uuid.NullUUID with Valid set to true and UUID set to
// the value pointed to by ptr.
func NullUUIDFromPtr(ptr *uuid.UUID) uuid.NullUUID {
	if ptr == nil {
		return uuid.NullUUID{
			Valid: false,
		}
	}
	return uuid.NullUUID{
		Valid: true,
		UUID:  *ptr,
	}
}
