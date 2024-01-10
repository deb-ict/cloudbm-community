package model

type Uri struct {
	Id        string
	Type      UriType
	Uri       string
	IsDefault bool
}

type UriFilter struct {
	TypeId string
}

func (m *Uri) IsTransient() bool {
	return m.Id == ""
}
