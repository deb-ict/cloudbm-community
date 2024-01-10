package model

import (
	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/localization"
)

type Industry struct {
	Id           string
	Key          string
	Translations []*IndustryTranslation
	IsSystem     bool
}

type IndustryTranslation struct {
	Language       string
	Name           string
	NormalizedName string
	Description    string
}

type IndustryFilter struct {
	Language string
	Name     string
}

func (m *Industry) Normalize(normalizer core.StringNormalizer) {
	m.Key = normalizer.NormalizeString(m.Key)
	for _, translation := range m.Translations {
		translation.Language = localization.NormalizeLanguage(translation.Language)
		translation.NormalizedName = normalizer.NormalizeString(translation.Name)
	}
}

func (m *Industry) UpdateModel(other *Industry) {
	m.Translations = make([]*IndustryTranslation, 0)
	for _, translation := range other.Translations {
		m.Translations = append(m.Translations, translation.Clone())
	}
}

func (m *Industry) GetTranslation(language string, defaultLanguage string) *IndustryTranslation {
	if len(m.Translations) == 0 {
		return &IndustryTranslation{}
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

func (m *Industry) TryGetTranslation(language string) (*IndustryTranslation, error) {
	normalizedLanguage := localization.NormalizeLanguage(language)
	for _, t := range m.Translations {
		if t.Language == normalizedLanguage {
			return t, nil
		}
	}
	return nil, core.ErrTranslationNotFound
}

func (m *Industry) IsTransient() bool {
	return m.Id == ""
}

func (m *Industry) Clone() *Industry {
	if m == nil {
		return nil
	}
	model := &Industry{
		Id:           m.Id,
		Key:          m.Key,
		Translations: make([]*IndustryTranslation, 0),
		IsSystem:     m.IsSystem,
	}
	for _, translation := range m.Translations {
		model.Translations = append(model.Translations, translation.Clone())
	}
	return model
}

func (m *IndustryTranslation) Clone() *IndustryTranslation {
	if m == nil {
		return nil
	}
	return &IndustryTranslation{
		Language:       m.Language,
		Name:           m.Name,
		NormalizedName: m.NormalizedName,
		Description:    m.Description,
	}
}
