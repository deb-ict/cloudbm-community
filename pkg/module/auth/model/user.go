package model

import "time"

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

func (m *User) Lock(duration time.Duration) {
	m.IsLocked = true
	m.LockEnd = time.Now().UTC().Add(duration)
}
