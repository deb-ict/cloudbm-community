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

func (m *Company) IsTransient() bool {
	return m.Id == ""
}

func (f *CompanyFilter) HasFilter() bool {
	return false
}
