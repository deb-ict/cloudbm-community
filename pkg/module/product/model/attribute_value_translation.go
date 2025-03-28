package model

type AttributeValueTranslation struct {
	Language       string
	Name           string
	NormalizedName string
	Slug           string
	Description    string
}

func (m *AttributeValueTranslation) Clone() *AttributeValueTranslation {
	if m == nil {
		return nil
	}
	return &AttributeValueTranslation{
		Language:       m.Language,
		Name:           m.Name,
		NormalizedName: m.NormalizedName,
		Slug:           m.Slug,
		Description:    m.Description,
	}
}
