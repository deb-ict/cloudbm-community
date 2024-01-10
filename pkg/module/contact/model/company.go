package model

type Company struct {
	Id        string
	Name      string
	VatNumber string
	Type      *CompanyType
	Industry  *Industry
	Addresses []*Address
	Emails    []*Email
	Phones    []*Phone
	Uris      []*Uri
	IsEnabled bool
	IsSystem  bool
}

type CompanyFilter struct {
	Name string
}

func (m *Company) UpdateModel(other *Company) {
	m.Name = other.Name
	m.VatNumber = other.VatNumber
	m.Type = other.Type.Clone()
	m.Industry = other.Industry.Clone()
	m.IsEnabled = other.IsEnabled
}

func (m *Company) IsTransient() bool {
	return m.Id == ""
}

func (m *Company) Clone() *Company {
	if m == nil {
		return nil
	}
	model := &Company{
		Id:        m.Id,
		Name:      m.Name,
		Type:      m.Type.Clone(),
		Industry:  m.Industry.Clone(),
		VatNumber: m.VatNumber,
		Addresses: m.Addresses,
		Emails:    m.Emails,
		Phones:    m.Phones,
		Uris:      m.Uris,
		IsEnabled: m.IsEnabled,
		IsSystem:  m.IsSystem,
	}
	for _, address := range m.Addresses {
		model.Addresses = append(model.Addresses, address.Clone())
	}
	for _, email := range m.Emails {
		model.Emails = append(model.Emails, email.Clone())
	}
	for _, phone := range m.Phones {
		model.Phones = append(model.Phones, phone.Clone())
	}
	for _, uri := range m.Uris {
		model.Uris = append(model.Uris, uri.Clone())
	}
	return model
}
