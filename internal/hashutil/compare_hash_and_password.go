package hashutil

import "golang.org/x/crypto/bcrypt"

// CompareHashAndPassword compares a password and a hash and returns true if they match.
func CompareHashAndPassword(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
