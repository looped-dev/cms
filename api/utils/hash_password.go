package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// hashPassword hashes the password using bcrypt and store the has instead of
// the plain text password.
func HashPassword(password string) (string, error) {
	if password == "" {
		return "", fmt.Errorf("password cannot be empty")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", fmt.Errorf("Error hashing password: %w", err)
	}
	return string(hash), nil
}
