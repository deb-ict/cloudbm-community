package model

import (
	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/localization"
)

type PhoneType struct {
	Id           string
	Key          string
	Translations []*PhoneTypeTranslation
	IsDefault    bool
	IsSystem     bool
}

type PhoneTypeTranslation struct {
	Language       string
	Name           string
	NormalizedName string
	Description    string
}

type PhoneTypeFilter struct {
	Language string
	Name     string
}

func (m *PhoneType) Normalize(normalizer core.StringNormalizer) {
	m.Key = normalizer.NormalizeString(m.Key)
	for _, translation := range m.Translations {
		translation.Language = localization.NormalizeLanguage(translation.Language)
		translation.NormalizedName = normalizer.NormalizeString(translation.Name)
	}
}

func (m *PhoneType) UpdateModel(other *PhoneType) {
	m.Translations = make([]*PhoneTypeTranslation, 0)
	m.IsDefault = other.IsDefault
	for _, translation := range other.Translations {
		m.Translations = append(m.Translations, translation.Clone())
	}
}

func (m *PhoneType) GetTranslation(language string, defaultLanguage string) *PhoneTypeTranslation {
	if len(m.Translations) == 0 {
		return &PhoneTypeTranslation{}
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

func (m *PhoneType) TryGetTranslation(language string) (*PhoneTypeTranslation, error) {
	normalizedLanguage := localization.NormalizeLanguage(language)
	for _, t := range m.Translations {
		if t.Language == normalizedLanguage {
			return t, nil
		}
	}
	return nil, core.ErrTranslationNotFound
}

func (m *PhoneType) IsTransient() bool {
	return m.Id == ""
}

func (m *PhoneType) Clone() *PhoneType {
	if m == nil {
		return nil
	}
	model := &PhoneType{
		Id:           m.Id,
		Key:          m.Key,
		Translations: make([]*PhoneTypeTranslation, 0),
		IsDefault:    m.IsDefault,
		IsSystem:     m.IsSystem,
	}
	for _, translation := range m.Translations {
		model.Translations = append(model.Translations, translation.Clone())
	}
	return model
}

func (m *PhoneTypeTranslation) Clone() *PhoneTypeTranslation {
	if m == nil {
		return nil
	}
	return &PhoneTypeTranslation{
		Language:       m.Language,
		Name:           m.Name,
		NormalizedName: m.NormalizedName,
		Description:    m.Description,
	}
}
