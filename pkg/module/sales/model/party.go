package model

type Party struct {
	CompanyName  string
	VatNumber    string
	FamilyName   string
	GivenName    string
	AddressLine1 string
	AddressLine2 string
	PostalCode   string
	City         string
	State        string
	Country      string
	Phone        string
	Email        string
}

func (m *Party) Clone() *Party {
	return &Party{
		CompanyName:  m.CompanyName,
		VatNumber:    m.VatNumber,
		FamilyName:   m.FamilyName,
		GivenName:    m.GivenName,
		AddressLine1: m.AddressLine1,
		AddressLine2: m.AddressLine2,
		PostalCode:   m.PostalCode,
		City:         m.City,
		State:        m.State,
		Country:      m.Country,
		Phone:        m.Phone,
		Email:        m.Email,
	}
}
