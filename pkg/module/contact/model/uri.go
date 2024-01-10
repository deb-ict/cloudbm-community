package model

type Uri struct {
	Id        string
	Type      *UriType
	Uri       string
	IsDefault bool
}

type UriFilter struct {
	TypeId string
}

func (m *Uri) UpdateModel(other *Uri) {
	m.Type = other.Type.Clone()
	m.Uri = other.Uri
	m.IsDefault = other.IsDefault
}

func (m *Uri) IsTransient() bool {
	return m.Id == ""
}

func (m *Uri) Clone() *Uri {
	if m == nil {
		return nil
	}
	return &Uri{
		Id:        m.Id,
		Type:      m.Type.Clone(),
		Uri:       m.Uri,
		IsDefault: m.IsDefault,
	}
}
