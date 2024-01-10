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

func (m *UserToken) UpdateModel(other *UserToken) {
	m.Token = other.Token
	m.Expiration = other.Expiration
}

func (m *UserToken) SetExpiration(duration time.Duration) {
	m.Expiration = time.Now().UTC().Add(duration)
}

func (m *UserToken) HasExpired() bool {
	return time.Now().UTC().After(m.Expiration)
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

func (m *UserToken) Clone() *UserToken {
	if m == nil {
		return nil
	}
	return &UserToken{
		Id:         m.Id,
		Type:       m.Type,
		Token:      m.Token,
		Expiration: m.Expiration,
	}
}
