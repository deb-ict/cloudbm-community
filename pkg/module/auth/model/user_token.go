package model

import (
	"time"
)

type UserTokenType int

const (
	Undefined UserTokenType = iota
	ActivationToken
	PasswordResetToken
)

type UserToken struct {
	Id         string
	Type       UserTokenType
	Token      string
	Expiration time.Time
}

func (t UserTokenType) String() string {
	switch t {
	case ActivationToken:
		return "ActivationToken"
	case PasswordResetToken:
		return "PasswordResetToken"
	default:
		return "Undefined"
	}
}

func (t *UserToken) HasExpired() bool {
	return time.Now().UTC().After(t.Expiration)
}
