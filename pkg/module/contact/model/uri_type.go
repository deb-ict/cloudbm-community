package model

import (
	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/localization"
)

type UriType struct {
	Id           string
	Key          string
	Translations []*UriTypeTranslation
	IsDefault    bool
	IsSystem     bool
}

type UriTypeTranslation struct {
	Language       string
	Name           string
	NormalizedName string
	Description    string
}

type UriTypeFilter struct {
	Language string
	Name     string
}

func (m *UriType) Normalize(normalizer core.StringNormalizer) {
	m.Key = normalizer.NormalizeString(m.Key)
	for _, translation := range m.Translations {
		translation.Language = localization.NormalizeLanguage(translation.Language)
		translation.NormalizedName = normalizer.NormalizeString(translation.Name)
	}
}

func (m *UriType) UpdateModel(other *UriType) {
	m.Translations = make([]*UriTypeTranslation, 0)
	m.IsDefault = other.IsDefault
	for _, translation := range other.Translations {
		m.Translations = append(m.Translations, translation.Clone())
	}
}

func (m *UriType) GetTranslation(language string, defaultLanguage string) *UriTypeTranslation {
	if len(m.Translations) == 0 {
		return &UriTypeTranslation{}
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

func (m *UriType) TryGetTranslation(language string) (*UriTypeTranslation, error) {
	normalizedLanguage := localization.NormalizeLanguage(language)
	for _, t := range m.Translations {
		if t.Language == normalizedLanguage {
			return t, nil
		}
	}
	return nil, core.ErrTranslationNotFound
}

func (m *UriType) IsTransient() bool {
	return m.Id == ""
}

func (m *UriType) Clone() *UriType {
	if m == nil {
		return nil
	}
	model := &UriType{
		Id:           m.Id,
		Key:          m.Key,
		Translations: make([]*UriTypeTranslation, 0),
		IsDefault:    m.IsDefault,
		IsSystem:     m.IsSystem,
	}
	for _, translation := range m.Translations {
		model.Translations = append(model.Translations, translation.Clone())
	}
	return model
}

func (m *UriTypeTranslation) Clone() *UriTypeTranslation {
	if m == nil {
		return nil
	}
	return &UriTypeTranslation{
		Language:       m.Language,
		Name:           m.Name,
		NormalizedName: m.NormalizedName,
		Description:    m.Description,
	}
}
