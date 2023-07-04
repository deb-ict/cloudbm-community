package model

type Contact struct {
	Id         string
	UserId     string
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

func (m *Contact) IsTransient() bool {
	return m.Id == ""
}

func (m *Contact) CanDelete() bool {
	return !m.IsSystem
}
