package sqlutil

import (
	"github.com/google/uuid"
)

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
