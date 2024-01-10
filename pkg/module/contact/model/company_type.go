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
	Language    string
	Name        string
	Description string
}

type CompanyTypeFilter struct {
	Language string
	Name     string
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
