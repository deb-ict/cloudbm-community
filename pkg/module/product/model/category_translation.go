package model

type CategoryTranslation struct {
	Language       string
	Name           string
	NormalizedName string
	Slug           string
	Summary        string
	Description    string
}

func (m *CategoryTranslation) Clone() *CategoryTranslation {
	if m == nil {
		return nil
	}
	return &CategoryTranslation{
		Language:       m.Language,
		Name:           m.Name,
		NormalizedName: m.NormalizedName,
		Slug:           m.Slug,
		Summary:        m.Summary,
		Description:    m.Description,
	}
}
