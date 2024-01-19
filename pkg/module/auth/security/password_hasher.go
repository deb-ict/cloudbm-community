package security

import (
	"golang.org/x/crypto/bcrypt"
)

type PasswordHasher interface {
	HashPassword(password string) (string, error)
	VerifyPassword(password string, hash string) bool
}

type defaultPasswordHasher struct {
}

func DefaultPasswordHasher() PasswordHasher {
	return &defaultPasswordHasher{}
}

func (h *defaultPasswordHasher) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	return string(bytes), err
}

func (h *defaultPasswordHasher) VerifyPassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
