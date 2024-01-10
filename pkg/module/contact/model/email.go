package model

type Email struct {
	Id        string
	Type      *EmailType
	Email     string
	IsDefault bool
}

type EmailFilter struct {
	TypeId string
}

func (m *Email) UpdateModel(other *Email) {
	m.Type = other.Type.Clone()
	m.Email = other.Email
	m.IsDefault = other.IsDefault
}

func (m *Email) IsTransient() bool {
	return m.Id == ""
}

func (m *Email) Clone() *Email {
	if m == nil {
		return nil
	}
	return &Email{
		Id:        m.Id,
		Type:      m.Type.Clone(),
		Email:     m.Email,
		IsDefault: m.IsDefault,
	}
}
