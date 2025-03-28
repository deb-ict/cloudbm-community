package model

type AttributeTranslation struct {
	Language       string
	Name           string
	NormalizedName string
	Slug           string
	Description    string
}

func (m *AttributeTranslation) Clone() *AttributeTranslation {
	if m == nil {
		return nil
	}
	return &AttributeTranslation{
		Language:       m.Language,
		Name:           m.Name,
		NormalizedName: m.NormalizedName,
		Slug:           m.Slug,
		Description:    m.Description,
	}
}
