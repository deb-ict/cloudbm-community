package model

import (
	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/localization"
)

type CompanyType struct {
	Id           string
	Key          string
	Translations []*CompanyTypeTranslation
	IsSystem     bool
}

type CompanyTypeTranslation struct {
	Language       string
	Name           string
	NormalizedName string
	Description    string
}

type CompanyTypeFilter struct {
	Language string
	Name     string
}

func (m *CompanyType) Normalize(normalizer core.StringNormalizer) {
	m.Key = normalizer.NormalizeString(m.Key)
	for _, translation := range m.Translations {
		translation.Language = localization.NormalizeLanguage(translation.Language)
		translation.NormalizedName = normalizer.NormalizeString(translation.Name)
	}
}

func (m *CompanyType) UpdateModel(other *CompanyType) {
	m.Translations = make([]*CompanyTypeTranslation, 0)
	for _, translation := range other.Translations {
		m.Translations = append(m.Translations, translation.Clone())
	}
}

func (m *CompanyType) GetTranslation(language string, defaultLanguage string) *CompanyTypeTranslation {
	if len(m.Translations) == 0 {
		return &CompanyTypeTranslation{}
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

func (m *CompanyType) TryGetTranslation(language string) (*CompanyTypeTranslation, error) {
	normalizedLanguage := localization.NormalizeLanguage(language)
	for _, t := range m.Translations {
		if t.Language == normalizedLanguage {
			return t, nil
		}
	}
	return nil, core.ErrTranslationNotFound
}

func (m *CompanyType) IsTransient() bool {
	return m.Id == ""
}

func (m *CompanyType) Clone() *CompanyType {
	if m == nil {
		return nil
	}
	model := &CompanyType{
		Id:           m.Id,
		Key:          m.Key,
		Translations: make([]*CompanyTypeTranslation, 0),
		IsSystem:     m.IsSystem,
	}
	for _, translation := range m.Translations {
		model.Translations = append(model.Translations, translation.Clone())
	}
	return model
}

func (m *CompanyTypeTranslation) Clone() *CompanyTypeTranslation {
	if m == nil {
		return nil
	}
	return &CompanyTypeTranslation{
		Language:       m.Language,
		Name:           m.Name,
		NormalizedName: m.NormalizedName,
		Description:    m.Description,
	}
}
