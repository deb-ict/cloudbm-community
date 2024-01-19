package util

import (
	"strings"
)

type UserNormalizer interface {
	NormalizeUsername(username string) string
	NormalizeEmail(email string) string
}

type defaultUserNormalizer struct {
}

func DefaultUserNormalizer() UserNormalizer {
	return &defaultUserNormalizer{}
}

func (n *defaultUserNormalizer) NormalizeUsername(username string) string {
	return strings.ToUpper(username)
}

func (n *defaultUserNormalizer) NormalizeEmail(email string) string {
	return strings.ToUpper(email)
}
