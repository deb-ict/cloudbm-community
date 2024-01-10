package model

type Phone struct {
	Id          string
	Type        *PhoneType
	PhoneNumber string
	Extension   string
	IsDefault   bool
}

type PhoneFilter struct {
	TypeId string
}

func (m *Phone) UpdateModel(other *Phone) {
	m.Type = other.Type.Clone()
	m.PhoneNumber = other.PhoneNumber
	m.Extension = other.Extension
	m.IsDefault = other.IsDefault
}

func (m *Phone) IsTransient() bool {
	return m.Id == ""
}

func (m *Phone) Clone() *Phone {
	if m == nil {
		return nil
	}
	return &Phone{
		Id:          m.Id,
		Type:        m.Type.Clone(),
		PhoneNumber: m.PhoneNumber,
		Extension:   m.Extension,
		IsDefault:   m.IsDefault,
	}
}
