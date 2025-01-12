package model

import (
	"time"
)

type Session struct {
	Id                   string
	UserId               string
	CreatedAt            time.Time
	UpdatedAt            time.Time
	ExpiresAt            time.Time
	Lifetime             time.Duration
	UseSlidingExpiration bool
	Data                 map[string]string
}

type SessionFilter struct {
	CreatedBefore time.Time
	CreatedAfter  time.Time
	UpdatedBefore time.Time
	UpdatedAfter  time.Time
	ExpiresBefore time.Time
	ExpiresAfter  time.Time
}

func (m *Session) UpdateModel(other *Session) {
	m.UserId = other.UserId
	m.Lifetime = other.Lifetime
	m.UpdatedAt = time.Now().UTC()
	m.Data = make(map[string]string)
	for k, v := range other.Data {
		m.Data[k] = v
	}
	if m.UseSlidingExpiration {
		m.ExpiresAt = m.UpdatedAt.Add(m.Lifetime)
	}
}

func (m *Session) SetExpiration(duration time.Duration) {
	m.UpdatedAt = time.Now().UTC()
	m.ExpiresAt = m.UpdatedAt.Add(duration)
}

func (m *Session) HasExpired() bool {
	return time.Now().UTC().After(m.ExpiresAt)
}

func (m *Session) IsTransient() bool {
	return m.Id == ""
}

func (m *Session) Clone() *Session {
	if m == nil {
		return nil
	}
	model := &Session{
		Id:                   m.Id,
		UserId:               m.UserId,
		CreatedAt:            m.CreatedAt,
		UpdatedAt:            m.UpdatedAt,
		ExpiresAt:            m.ExpiresAt,
		Lifetime:             m.Lifetime,
		UseSlidingExpiration: m.UseSlidingExpiration,
		Data:                 make(map[string]string),
	}
	for k, v := range m.Data {
		model.Data[k] = v
	}
	return model
}
