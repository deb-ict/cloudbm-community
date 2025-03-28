package model

type ProductTranslation struct {
	Language       string
	Name           string
	NormalizedName string
	Slug           string
	Summary        string
	Description    string
}

func (m *ProductTranslation) Clone() *ProductTranslation {
	if m == nil {
		return nil
	}
	return &ProductTranslation{
		Language:       m.Language,
		Name:           m.Name,
		NormalizedName: m.NormalizedName,
		Slug:           m.Slug,
		Summary:        m.Summary,
		Description:    m.Description,
	}
}
