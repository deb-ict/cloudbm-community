package model

import (
	"time"
)

type User struct {
	Id                 string
	Username           string
	PasswordHash       string
	Email              string
	EmailVerified      bool
	Phone              string
	PhoneVerified      bool
	NormalizedUsername string
	NormalizedEmail    string
	LoginFailures      int
	IsLocked           bool
	IsEnabled          bool
	LockEnd            time.Time
	Tokens             []*UserToken
}

func (m *User) UpdateModel(other *User) {
	m.EmailVerified = other.EmailVerified
	m.Phone = other.Phone
	m.PhoneVerified = other.PhoneVerified
	m.LoginFailures = other.LoginFailures
	m.IsEnabled = other.IsEnabled
}

type UserFilter struct {
	Username string
	Email    string
}

func (m *User) Locked() bool {
	if !m.IsLocked {
		return false
	}
	return time.Now().UTC().After(m.LockEnd)
}

func (m *User) Lock(duration time.Duration) {
	m.IsLocked = true
	m.LockEnd = time.Now().UTC().Add(duration)
}

func (m *User) Clone() *User {
	if m == nil {
		return nil
	}
	model := &User{
		Id:                 m.Id,
		Username:           m.Username,
		PasswordHash:       m.PasswordHash,
		Email:              m.Email,
		EmailVerified:      m.EmailVerified,
		Phone:              m.Phone,
		PhoneVerified:      m.PhoneVerified,
		NormalizedUsername: m.NormalizedUsername,
		NormalizedEmail:    m.NormalizedEmail,
		LoginFailures:      m.LoginFailures,
		IsLocked:           m.IsLocked,
		IsEnabled:          m.IsEnabled,
		LockEnd:            m.LockEnd,
		Tokens:             make([]*UserToken, 0),
	}
	for _, token := range m.Tokens {
		model.Tokens = append(model.Tokens, token.Clone())
	}
	return model
}
