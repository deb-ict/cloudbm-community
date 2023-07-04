package model

type Address struct {
	Id         string
	Type       AddressType
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
}

func (m *Address) IsTransient() bool {
	return m.Id == ""
}

func (m *Address) CanDelete() bool {
	return !m.IsDefault
}
