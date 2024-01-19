package hashutil

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestGenerateHashFromPassword(t *testing.T) {
	password := "testPassword"
	hash, err := GenerateHashFromPassword(password)
	require.NoError(t, err)

	// Check if the password matches the hash
	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	require.NoError(t, err)
}
