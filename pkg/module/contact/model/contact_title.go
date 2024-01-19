package model

import (
	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/localization"
)

type ContactTitle struct {
	Id           string
	Key          string
	Translations []*ContactTitleTranslation
	IsSystem     bool
}

type ContactTitleTranslation struct {
	Language       string
	Name           string
	NormalizedName string
	Description    string
}

type ContactTitleFilter struct {
	Language string
	Name     string
}

func (m *ContactTitle) Normalize(normalizer core.StringNormalizer) {
	m.Key = normalizer.NormalizeString(m.Key)
	for _, translation := range m.Translations {
		translation.Language = localization.NormalizeLanguage(translation.Language)
		translation.NormalizedName = normalizer.NormalizeString(translation.Name)
	}
}

func (m *ContactTitle) UpdateModel(other *ContactTitle) {
	m.Translations = make([]*ContactTitleTranslation, 0)
	for _, translation := range other.Translations {
		m.Translations = append(m.Translations, translation.Clone())
	}
}

func (m *ContactTitle) GetTranslation(language string, defaultLanguage string) *ContactTitleTranslation {
	if len(m.Translations) == 0 {
		return &ContactTitleTranslation{}
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

func (m *ContactTitle) TryGetTranslation(language string) (*ContactTitleTranslation, error) {
	normalizedLanguage := localization.NormalizeLanguage(language)
	for _, t := range m.Translations {
		if t.Language == normalizedLanguage {
			return t, nil
		}
	}
	return nil, core.ErrTranslationNotFound
}

func (m *ContactTitle) IsTransient() bool {
	return m.Id == ""
}

func (m *ContactTitle) Clone() *ContactTitle {
	if m == nil {
		return nil
	}
	model := &ContactTitle{
		Id:           m.Id,
		Key:          m.Key,
		Translations: make([]*ContactTitleTranslation, 0),
		IsSystem:     m.IsSystem,
	}
	for _, translation := range m.Translations {
		model.Translations = append(model.Translations, translation.Clone())
	}
	return model
}

func (m *ContactTitleTranslation) Clone() *ContactTitleTranslation {
	if m == nil {
		return nil
	}
	return &ContactTitleTranslation{
		Language:       m.Language,
		Name:           m.Name,
		NormalizedName: m.NormalizedName,
		Description:    m.Description,
	}
}
