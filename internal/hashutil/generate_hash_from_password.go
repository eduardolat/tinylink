package hashutil

import "golang.org/x/crypto/bcrypt"

// GenerateHashFromPassword generates a hash from a password.
func GenerateHashFromPassword(password string) (string, error) {
	hashByte, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return "", err
	}
	return string(hashByte), nil
}
