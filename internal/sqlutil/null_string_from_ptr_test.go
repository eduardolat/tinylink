package sqlutil

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNullStringFromPtr(t *testing.T) {
	t.Run("with nil pointer", func(t *testing.T) {
		result := NullStringFromPtr(nil)
		assert.False(t, result.Valid)
	})

	t.Run("with valid string pointer", func(t *testing.T) {
		testString := "test string"
		result := NullStringFromPtr(&testString)
		assert.True(t, result.Valid)
		assert.Equal(t, testString, result.String)
	})

	t.Run("with multiple strings", func(t *testing.T) {
		for i := 0; i < 10; i++ {
			testString := fmt.Sprintf("test string %d", i)
			result := NullStringFromPtr(&testString)
			assert.True(t, result.Valid)
			assert.Equal(t, testString, result.String)
		}
	})
}
