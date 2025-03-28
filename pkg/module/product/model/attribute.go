package model

import (
	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/localization"
	"github.com/gosimple/slug"
)

type Attribute struct {
	Id           string
	Translations []*AttributeTranslation
	IsEnabled    bool
}

func (m *Attribute) Normalize(normalizer core.StringNormalizer) {
	for _, translation := range m.Translations {
		translation.Language = localization.NormalizeLanguage(translation.Language)
		translation.NormalizedName = normalizer.NormalizeString(translation.Name)
		if translation.Slug == "" {
			translation.Slug = slug.MakeLang(translation.NormalizedName, translation.Language)
		}
		translation.Slug = normalizer.NormalizeString(translation.Slug)
	}
}

func (m *Attribute) UpdateModel(other *Attribute) {
	m.Translations = make([]*AttributeTranslation, 0)
	m.IsEnabled = other.IsEnabled
	for _, translation := range other.Translations {
		m.Translations = append(m.Translations, translation.Clone())
	}
}

func (m *Attribute) GetTranslation(language string, defaultLanguage string) *AttributeTranslation {
	if len(m.Translations) == 0 {
		return &AttributeTranslation{}
	}

	translation, err := m.TryGetTranslation(language)
	if err == core.ErrTranslationNotFound && language != defaultLanguage {
		translation, err = m.TryGetTranslation(defaultLanguage)
	}
	if err == core.ErrTranslationNotFound {
		translation = m.Translations[0]
	}

	return translation
}

func (m *Attribute) TryGetTranslation(language string) (*AttributeTranslation, error) {
	normalizedLanguage := localization.NormalizeLanguage(language)
	for _, t := range m.Translations {
		if t.Language == normalizedLanguage {
			return t, nil
		}
	}

	return nil, core.ErrTranslationNotFound
}

func (m *Attribute) IsTransient() bool {
	return m.Id == ""
}

func (m *Attribute) Clone() *Attribute {
	if m == nil {
		return nil
	}
	model := &Attribute{
		Id:           m.Id,
		Translations: make([]*AttributeTranslation, 0),
		IsEnabled:    m.IsEnabled,
	}
	for _, translation := range m.Translations {
		model.Translations = append(model.Translations, translation.Clone())
	}
	return model
}
