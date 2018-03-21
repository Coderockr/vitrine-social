package security

import "golang.org/x/crypto/bcrypt"

// CompareHashAndPassword compares a storaged password hash
// with the a provided password, returns error if fail, or
// nil if they match
func CompareHashAndPassword(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

// GenerateHash generates a hash from the password informed, or error if fail
func GenerateHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
