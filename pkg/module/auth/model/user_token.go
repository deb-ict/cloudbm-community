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
	id         string
	tokenType  UserTokenType
	tokenValue string
	expiration time.Time
}

func (m *UserToken) Id() string {
	if m == nil {
		return ""
	}
	return m.id
}

func (m *UserToken) TokenType() UserTokenType {
	if m == nil {
		return UserTokenType_Undefined
	}
	return m.tokenType
}

func (m *UserToken) TokenValue() string {
	if m == nil {
		return ""
	}
	return m.tokenValue
}

func (m *UserToken) SetTokenValue(value string) {
	if m != nil {
		m.tokenValue = value
	}
}

func (m *UserToken) Expiration() time.Time {
	if m == nil {
		return time.Now().UTC()
	}
	return m.expiration
}

func (m *UserToken) SetAbsoluteExpiration(value time.Time) {
	if m != nil {
		m.expiration = value
	}
}

func (m *UserToken) SetRelativeExpiration(duration time.Duration) {
	if m != nil {
		m.expiration = time.Now().UTC().Add(duration)
	}
}

func (m *UserToken) UpdateModel(other *UserToken) {
	if m == nil || other == nil {
		return
	}
	m.tokenValue = other.tokenValue
	m.expiration = other.expiration
}

func (m *UserToken) HasExpired() bool {
	if m == nil {
		return true
	}
	return time.Now().UTC().After(m.expiration)
}

func (m *UserToken) Clone() *UserToken {
	if m == nil {
		return nil
	}
	return &UserToken{
		id:         m.id,
		tokenType:  m.tokenType,
		tokenValue: m.tokenValue,
		expiration: m.expiration,
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
