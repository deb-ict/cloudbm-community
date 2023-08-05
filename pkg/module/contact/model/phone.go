package model

type Phone struct {
	Id          string
	Type        PhoneType
	PhoneNumber string
	Extension   string
	IsDefault   bool
}

type PhoneFilter struct {
}

func (m *Phone) IsTransient() bool {
	return m.Id == ""
}

func (f *PhoneFilter) HasFilter() bool {
	return false
}
