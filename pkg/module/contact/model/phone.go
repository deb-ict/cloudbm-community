package model

type Phone struct {
	Id          string
	Type        PhoneType
	PhoneNumber string
	Extension   string
	IsDefault   bool
}

type PhoneFilter struct {
	TypeId string
}

func (m *Phone) IsTransient() bool {
	return m.Id == ""
}
