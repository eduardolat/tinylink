package hashutil

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestCompareHashAndPassword(t *testing.T) {
	password := "testPassword"
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	require.NoError(t, err)

	// Check if the password matches the hash
	matches := CompareHashAndPassword(string(hash), password)
	require.True(t, matches)
}
