package security

import "golang.org/x/crypto/bcrypt"

// Function to generated a hash from password. We must convert string to byte slice
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// Function to compare my password to a hash password
func ComparePassword(passwordHash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
}
