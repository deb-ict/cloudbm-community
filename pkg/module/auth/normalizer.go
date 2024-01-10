package auth

import (
	"strings"

	"github.com/deb-ict/cloudbm-community/pkg/module/auth/model"
)

type UserNormalizer interface {
	NormalizeUser(user *model.User)
	NormalizeUsername(username string) string
	NormalizeEmail(email string) string
}

type defaultUserNormalizer struct {
}

func DefaultUserNormalizer() UserNormalizer {
	return &defaultUserNormalizer{}
}

func (n *defaultUserNormalizer) NormalizeUser(user *model.User) {
	user.NormalizedUsername = n.NormalizeUsername(user.Username)
	user.NormalizedEmail = n.NormalizeEmail(user.Email)
}

func (n *defaultUserNormalizer) NormalizeUsername(username string) string {
	return strings.ToUpper(username)
}

func (n *defaultUserNormalizer) NormalizeEmail(email string) string {
	return strings.ToUpper(email)
}
