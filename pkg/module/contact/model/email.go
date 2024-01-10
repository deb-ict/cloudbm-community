package model

type Email struct {
	Id        string
	Type      EmailType
	Email     string
	IsDefault bool
}

type EmailFilter struct {
	TypeId string
}

func (m *Email) IsTransient() bool {
	return m.Id == ""
}
