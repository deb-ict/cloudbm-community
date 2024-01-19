package model

type Address struct {
	Id         string
	Type       *AddressType
	Street     string
	StreetNr   string
	Unit       string
	PostalCode string
	City       string
	State      string
	Country    string
	IsDefault  bool
}

type AddressFilter struct {
	TypeId string
}

func (m *Address) UpdateModel(other *Address) {
	m.Type = other.Type.Clone()
	m.Street = other.Street
	m.StreetNr = other.StreetNr
	m.Unit = other.Unit
	m.PostalCode = other.PostalCode
	m.City = other.City
	m.State = other.State
	m.Country = other.Country
	m.IsDefault = other.IsDefault
}

func (m *Address) IsTransient() bool {
	return m.Id == ""
}

func (m *Address) Clone() *Address {
	if m == nil {
		return nil
	}
	return &Address{
		Id:         m.Id,
		Type:       m.Type.Clone(),
		Street:     m.Street,
		StreetNr:   m.StreetNr,
		Unit:       m.Unit,
		PostalCode: m.PostalCode,
		City:       m.City,
		State:      m.State,
		Country:    m.Country,
		IsDefault:  m.IsDefault,
	}
}
