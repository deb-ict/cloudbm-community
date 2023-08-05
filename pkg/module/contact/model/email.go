package model

type Email struct {
	Id        string
	Type      EmailType
	Email     string
	IsDefault bool
}

type EmailFilter struct {
}

func (m *Email) IsTransient() bool {
	return m.Id == ""
}

func (f *EmailFilter) HasFilter() bool {
	return false
}
