package sqlutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNullBoolFromPtr(t *testing.T) {
	t.Run("with nil pointer", func(t *testing.T) {
		result := NullBoolFromPtr(nil)
		assert.False(t, result.Valid)
	})

	t.Run("with valid bool pointer (true)", func(t *testing.T) {
		testBool := true
		result := NullBoolFromPtr(&testBool)
		assert.True(t, result.Valid)
		assert.Equal(t, testBool, result.Bool)
	})

	t.Run("with valid bool pointer (false)", func(t *testing.T) {
		testBool := false
		result := NullBoolFromPtr(&testBool)
		assert.True(t, result.Valid)
		assert.Equal(t, testBool, result.Bool)
	})
}
