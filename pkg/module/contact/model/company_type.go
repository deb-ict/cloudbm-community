package model

import (
	"github.com/deb-ict/cloudbm-community/pkg/localization"
)

type CompanyType struct {
	Id           string
	Translations []CompanyTypeTranslation
	IsDefault    bool
	IsSystem     bool
}

type CompanyTypeTranslation struct {
	Language    string
	Name        string
	Description string
}

type CompanyTypeFilter struct {
	Name string
}

func (m *CompanyType) GetTranslation(language string) CompanyTypeTranslation {
	if len(m.Translations) == 0 {
		return CompanyTypeTranslation{}
	}

	normalizedLanguage := localization.NormalizeLanguage(language)
	for _, t := range m.Translations {
		if t.Language == normalizedLanguage {
			return t
		}
	}

	return m.Translations[0]
}

func (m *CompanyType) IsTransient() bool {
	return m.Id == ""
}

func (m *CompanyType) CanDelete() bool {
	return !m.IsDefault && !m.IsSystem
}
