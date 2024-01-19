package model

type Contact struct {
	Id         string
	Title      *ContactTitle
	FamilyName string
	MiddleName string
	GivenName  string
	Addresses  []*Address
	Emails     []*Email
	Phones     []*Phone
	Uris       []*Uri
	IsEnabled  bool
	IsSystem   bool
}

type ContactFilter struct {
	Name string
}

func (m *Contact) UpdateModel(other *Contact) {
	m.Title = other.Title.Clone()
	m.FamilyName = other.FamilyName
	m.MiddleName = other.MiddleName
	m.GivenName = other.GivenName
	m.IsEnabled = other.IsEnabled
}

func (m *Contact) IsTransient() bool {
	return m.Id == ""
}

func (m *Contact) Clone() *Contact {
	if m == nil {
		return nil
	}
	model := &Contact{
		Id:         m.Id,
		Title:      m.Title.Clone(),
		FamilyName: m.FamilyName,
		MiddleName: m.MiddleName,
		GivenName:  m.GivenName,
		Addresses:  make([]*Address, 0),
		Emails:     make([]*Email, 0),
		Phones:     make([]*Phone, 0),
		Uris:       make([]*Uri, 0),
		IsEnabled:  m.IsEnabled,
		IsSystem:   m.IsSystem,
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
