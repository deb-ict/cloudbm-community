package model

import (
	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/localization"
	"github.com/gosimple/slug"
)

type AttributeValue struct {
	Id           string
	AttributeId  string
	Translations []*AttributeValueTranslation
	Value        string
	IsEnabled    bool
}

func (m *AttributeValue) Normalize(normalizer core.StringNormalizer) {
	for _, translation := range m.Translations {
		translation.Language = localization.NormalizeLanguage(translation.Language)
		translation.NormalizedName = normalizer.NormalizeString(translation.Name)
		if translation.Slug == "" {
			translation.Slug = slug.MakeLang(translation.NormalizedName, translation.Language)
		}
		translation.Slug = normalizer.NormalizeString(translation.Slug)
	}
}

func (m *AttributeValue) UpdateModel(other *AttributeValue) {
	m.Translations = make([]*AttributeValueTranslation, 0)
	m.IsEnabled = other.IsEnabled
	for _, translation := range other.Translations {
		m.Translations = append(m.Translations, translation.Clone())
	}
	m.Value = other.Value
}

func (m *AttributeValue) GetTranslation(language string, defaultLanguage string) *AttributeValueTranslation {
	if len(m.Translations) == 0 {
		return &AttributeValueTranslation{}
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

func (m *AttributeValue) TryGetTranslation(language string) (*AttributeValueTranslation, error) {
	normalizedLanguage := localization.NormalizeLanguage(language)
	for _, t := range m.Translations {
		if t.Language == normalizedLanguage {
			return t, nil
		}
	}

	return nil, core.ErrTranslationNotFound
}

func (m *AttributeValue) IsTransient() bool {
	return m.Id == ""
}

func (m *AttributeValue) Clone() *AttributeValue {
	if m == nil {
		return nil
	}
	model := &AttributeValue{
		Id:           m.Id,
		AttributeId:  m.AttributeId,
		Translations: make([]*AttributeValueTranslation, 0),
		Value:        m.Value,
		IsEnabled:    m.IsEnabled,
	}
	for _, translation := range m.Translations {
		model.Translations = append(model.Translations, translation.Clone())
	}
	return model
}
