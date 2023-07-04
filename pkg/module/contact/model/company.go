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
	IsSystem  bool
}

type CompanyFilter struct {
	Name string
}

func (m *Company) IsTransient() bool {
	return m.Id == ""
}

func (m *Company) CanDelete() bool {
	return !m.IsSystem
}
