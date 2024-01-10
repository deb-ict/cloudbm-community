package model

import (
	"time"

	"github.com/deb-ict/cloudbm-community/pkg/module/auth/security"
	"github.com/deb-ict/cloudbm-community/pkg/module/auth/util"
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

func (m *User) Normalize(normalizer util.UserNormalizer) {
	m.NormalizedUsername = normalizer.NormalizeUsername(m.Username)
	m.NormalizedEmail = normalizer.NormalizeEmail(m.Email)
}

func (m *User) UpdateModel(other *User) {
	m.EmailVerified = other.EmailVerified
	m.Phone = other.Phone
	m.PhoneVerified = other.PhoneVerified
	m.LoginFailures = other.LoginFailures
	m.IsEnabled = other.IsEnabled
}

func (m *User) Lock(duration time.Duration) {
	m.IsLocked = true
	m.LockEnd = time.Now().UTC().Add(duration)
}

func (m *User) Unlock() {
	m.IsLocked = false
	m.LockEnd = time.Now().UTC()
}

func (m *User) VerifyPassword(hasher security.PasswordHasher, password string) bool {
	valid := hasher.VerifyPassword(password, m.PasswordHash)
	if !valid {
		m.LoginFailures++
	}
	return valid
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
