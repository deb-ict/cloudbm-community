package model

import (
	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/localization"
)

type EmailType struct {
	Id           string
	Key          string
	Translations []*EmailTypeTranslation
	IsDefault    bool
	IsSystem     bool
}

type EmailTypeTranslation struct {
	Language       string
	Name           string
	NormalizedName string
	Description    string
}

type EmailTypeFilter struct {
	Language string
	Name     string
}

func (m *EmailType) Normalize(normalizer core.StringNormalizer) {
	m.Key = normalizer.NormalizeString(m.Key)
	for _, translation := range m.Translations {
		translation.Language = localization.NormalizeLanguage(translation.Language)
		translation.NormalizedName = normalizer.NormalizeString(translation.Name)
	}
}

func (m *EmailType) UpdateModel(other *EmailType) {
	m.Translations = make([]*EmailTypeTranslation, 0)
	m.IsDefault = other.IsDefault
	for _, translation := range other.Translations {
		m.Translations = append(m.Translations, translation.Clone())
	}
}

func (m *EmailType) GetTranslation(language string, defaultLanguage string) *EmailTypeTranslation {
	if len(m.Translations) == 0 {
		return &EmailTypeTranslation{}
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

func (m *EmailType) TryGetTranslation(language string) (*EmailTypeTranslation, error) {
	normalizedLanguage := localization.NormalizeLanguage(language)
	for _, t := range m.Translations {
		if t.Language == normalizedLanguage {
			return t, nil
		}
	}
	return nil, core.ErrTranslationNotFound
}

func (m *EmailType) IsTransient() bool {
	return m.Id == ""
}

func (m *EmailType) Clone() *EmailType {
	if m == nil {
		return nil
	}
	model := &EmailType{
		Id:           m.Id,
		Key:          m.Key,
		Translations: make([]*EmailTypeTranslation, 0),
		IsDefault:    m.IsDefault,
		IsSystem:     m.IsSystem,
	}
	for _, translation := range m.Translations {
		model.Translations = append(model.Translations, translation.Clone())
	}
	return model
}

func (m *EmailTypeTranslation) Clone() *EmailTypeTranslation {
	if m == nil {
		return nil
	}
	return &EmailTypeTranslation{
		Language:       m.Language,
		Name:           m.Name,
		NormalizedName: m.NormalizedName,
		Description:    m.Description,
	}
}
