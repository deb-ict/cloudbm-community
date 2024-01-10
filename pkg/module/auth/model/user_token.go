package model

import (
	"time"
)

type UserTokenType int

const (
	UserTokenType_Undefined UserTokenType = iota
	UserTokenType_ActivationToken
	UserTokenType_PasswordResetToken
)

type UserToken struct {
	Id         string
	Type       UserTokenType
	Token      string
	Expiration time.Time
}

func (m *UserToken) UpdateModel(other *UserToken) {
	if m == nil || other == nil {
		return
	}
	m.Token = other.Token
	m.Expiration = other.Expiration
}

func (m *UserToken) SetExpiration(duration time.Duration) {
	if m != nil {
		m.Expiration = time.Now().UTC().Add(duration)
	}
}

func (m *UserToken) HasExpired() bool {
	if m == nil {
		return true
	}
	return time.Now().UTC().After(m.Expiration)
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

func (t UserTokenType) String() string {
	switch t {
	case UserTokenType_ActivationToken:
		return "ActivationToken"
	case UserTokenType_PasswordResetToken:
		return "PasswordResetToken"
	default:
		return "Undefined"
	}
}
