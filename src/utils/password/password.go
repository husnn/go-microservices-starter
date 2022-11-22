package password

import (
	"golang.org/x/crypto/bcrypt"
)

func Hash(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

func Verify(hash, provided string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(provided))
}
