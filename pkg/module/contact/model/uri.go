package model

type Uri struct {
	Id        string
	Type      UriType
	Uri       string
	IsDefault bool
}

type UriFilter struct {
}

func (m *Uri) IsTransient() bool {
	return m.Id == ""
}

func (m *Uri) CanDelete() bool {
	return !m.IsDefault
}
