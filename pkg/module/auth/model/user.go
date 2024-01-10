package model

import (
	"time"

	"github.com/deb-ict/cloudbm-community/pkg/module/auth/security"
	"github.com/deb-ict/cloudbm-community/pkg/module/auth/util"
)

type User struct {
	id                 string
	username           string
	passwordHash       string
	email              string
	emailVerified      bool
	phone              string
	phoneVerified      bool
	normalizedUsername string
	normalizedEmail    string
	loginFailures      int
	isLocked           bool
	isEnabled          bool
	lockEnd            time.Time
	tokens             []*UserToken
}

type NewUserOptions struct {
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

func NewUser(opts *NewUserOptions) *User {
	if opts == nil {
		opts = &NewUserOptions{}
	}
	return &User{
		id:           opts.Id,
		username:     opts.Username,
		passwordHash: opts.PasswordHash,
		email:        opts.Email,
	}
}

func (m *User) Id() string {
	if m == nil {
		return ""
	}
	return m.id
}

func (m *User) Username() string {
	if m == nil {
		return ""
	}
	return m.username
}

func (m *User) NormalizedUsername() string {
	if m == nil {
		return ""
	}
	return m.normalizedUsername
}

func (m *User) Email() string {
	if m == nil {
		return ""
	}
	return m.email
}

func (m *User) NormalizedEmail() string {
	if m == nil {
		return ""
	}
	return m.normalizedEmail
}

func (m *User) EmailVerified() bool {
	if m == nil {
		return false
	}
	return m.emailVerified
}

func (m *User) SetEmailVerified(value bool) {
	if m != nil {
		m.emailVerified = value
	}
}

func (m *User) Phone() string {
	if m == nil {
		return ""
	}
	return m.phone
}

func (m *User) SetPhone(value string) {
	if m != nil {
		m.phone = value
	}
}

func (m *User) PhoneVerified() bool {
	if m == nil {
		return false
	}
	return m.phoneVerified
}

func (m *User) SetPhoneVerified(value bool) {
	if m != nil {
		m.phoneVerified = value
	}
}

func (m *User) LoginFailures() int {
	if m == nil {
		return 0
	}
	return m.loginFailures
}

func (m *User) IsLocked() bool {
	if !m.isLocked {
		return false
	}
	return time.Now().UTC().After(m.lockEnd)
}

func (m *User) IsEnabled() bool {
	if m == nil {
		return false
	}
	return m.isEnabled
}

func (m *User) SetEnabled(value bool) {
	if m != nil {
		m.isEnabled = value
	}
}

func (m *User) Normalize(normalizer util.UserNormalizer) {
	m.normalizedUsername = normalizer.NormalizeUsername(m.username)
	m.normalizedEmail = normalizer.NormalizeEmail(m.email)
}

func (m *User) UpdateModel(other *User) {
	m.emailVerified = other.emailVerified
	m.phone = other.phone
	m.phoneVerified = other.phoneVerified
	m.loginFailures = other.loginFailures
	m.isEnabled = other.isEnabled
}

func (m *User) Lock(duration time.Duration) {
	m.isLocked = true
	m.lockEnd = time.Now().UTC().Add(duration)
}

func (m *User) Unlock() {
	m.isLocked = false
	m.lockEnd = time.Now().UTC()
}

func (m *User) VerifyPassword(hasher security.PasswordHasher, password string) bool {
	valid := hasher.VerifyPassword(password, m.passwordHash)
	if !valid {
		m.loginFailures++
	}
	return valid
}

func (m *User) Clone() *User {
	if m == nil {
		return nil
	}
	model := &User{
		id:                 m.id,
		username:           m.username,
		passwordHash:       m.passwordHash,
		email:              m.email,
		emailVerified:      m.emailVerified,
		phone:              m.phone,
		phoneVerified:      m.phoneVerified,
		normalizedUsername: m.normalizedUsername,
		normalizedEmail:    m.normalizedEmail,
		loginFailures:      m.loginFailures,
		isLocked:           m.isLocked,
		isEnabled:          m.isEnabled,
		lockEnd:            m.lockEnd,
		tokens:             make([]*UserToken, 0),
	}
	for _, token := range m.tokens {
		model.tokens = append(model.tokens, token.Clone())
	}
	return model
}
