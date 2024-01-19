package sqlutil

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNullUUIDFromPtr(t *testing.T) {
	t.Run("with nil pointer", func(t *testing.T) {
		result := NullUUIDFromPtr(nil)
		assert.False(t, result.Valid)
	})

	t.Run("with valid UUID pointer", func(t *testing.T) {
		testUUID := uuid.New()
		result := NullUUIDFromPtr(&testUUID)
		assert.True(t, result.Valid)
		assert.Equal(t, testUUID, result.UUID)
	})

	t.Run("with multiple UUIDs", func(t *testing.T) {
		for i := 0; i < 10; i++ {
			testUUID := uuid.New()
			result := NullUUIDFromPtr(&testUUID)
			assert.True(t, result.Valid)
			assert.Equal(t, testUUID, result.UUID)
		}
	})
}
